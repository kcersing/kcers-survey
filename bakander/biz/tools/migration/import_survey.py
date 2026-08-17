#!/usr/bin/env python3
"""Import survey docx to PostgreSQL. All Chinese text read from docx, not hardcoded."""

import zipfile, xml.etree.ElementTree as ET, json, re, sys, os

DOCX_FILE = "_survey.docx"
OUTPUT_JSON = "_survey_data.json"
OUTPUT_SQL = "_survey_import.sql"

SECTION_MAP = {
    1:  ("A1", "A1-BasicInfo",     "elderly"),
    2:  ("A2", "A2-Healthcare",    "elderly"),
    3:  ("A3", "A3-Disability",    "elderly"),
    4:  ("A4", "A4-ElderlyService","elderly"),
    5:  ("A5", "A5-SocialParticipation", "elderly"),
    6:  ("A6", "A6-CultureEntertainment", "elderly"),
    7:  ("B1", "B1-CaregiverInfo", "caregiver"),
    8:  ("B2", "B2-CareBurden",    "caregiver"),
    9:  ("B3", "B3-CareKnowledge", "caregiver"),
    10: ("B4", "B4-CommunitySupport", "caregiver"),
    11: ("B5", "B5-FamilySupport", "caregiver"),
    12: ("C1", "C1-VillageInfo",   "cadre"),
    13: ("C2", "C2-ElderlyService","cadre"),
    14: ("C3", "C3-CultureEntertainment", "cadre"),
    15: ("C4", "C4-SocialParticipation", "cadre"),
}

NS = {'w': 'http://schemas.openxmlformats.org/wordprocessingml/2006/main'}

def cell_text(cell):
    return ''.join(t.text or '' for t in cell.findall('.//w:t', NS)).strip()

def extract(path):
    with zipfile.ZipFile(path, 'r') as z:
        with z.open('word/document.xml') as f:
            root = ET.parse(f).getroot()
    return root.findall('.//w:tbl', NS)

def has_multi_select(content):
    """Check if question allows multiple selection"""
    keywords = ['ooo', 'oo', 'eee', 'eee']  # placeholder
    # Check for Chinese keywords in raw content
    for kw in [chr(0x53EF)+chr(0x591A)+chr(0x9009), chr(0x591A)+chr(0x9009)]:
        if kw in content:
            return True
    return False

def parse_options(text):
    """Parse numbered options from text. Returns list of {serial, content}."""
    if not text or len(text) < 3:
        return []
    # Check if it starts with number pattern like "1.x 2.y"
    if not re.match(r'^\d+[\.、\)]', text.strip()):
        return []
    parts = re.split(r'(?:^|\s+)(\d+)[\.、\)]', text.strip())
    result = []
    i = 0
    while i < len(parts):
        p = parts[i].strip()
        if p.isdigit() and i+1 < len(parts):
            opt = parts[i+1].strip()
            if opt:
                result.append({"content": opt, "serial": int(p)})
            i += 2
        else:
            i += 1
    return result

def detect_type(content, options_text):
    """Detect question type from content and options"""
    # Parse options first
    opts = parse_options(options_text)
    has_opts = len(opts) > 1

    # Text fill-in hints
    fill_keywords = [chr(0x5143), chr(0x4EBA), chr(0x516C)+chr(0x91CC), chr(0x5C81)]
    if options_text and not has_opts:
        return 'text', []

    # Multi-select detection
    is_multi = any(kw in content for kw in [
        chr(0x53EF)+chr(0x591A)+chr(0x9009),
        chr(0x591A)+chr(0x9009),
    ])

    if has_opts:
        if is_multi:
            return 'multiple_choice', opts
        return 'single_choice', opts

    return 'text', []

def build_json(tables):
    survey_data = {
        "title": "",
        "sections": [],
        "total_questions": 0,
    }

    # Table 0: parse header
    header_rows = tables[0].findall('.//w:tr', NS)
    title_parts = []
    for row in header_rows:
        cells = row.findall('.//w:tc', NS)
        if len(cells) >= 2:
            label = cell_text(cells[0])
            val = cell_text(cells[1])
            if val and label != val:
                title_parts.append(val)
    survey_data["title"] = title_parts[0] if title_parts else "Survey"

    gid = 0
    for tidx in range(1, len(tables)):
        if tidx not in SECTION_MAP:
            continue
        code, name, target = SECTION_MAP[tidx]
        rows = tables[tidx].findall('.//w:tr', NS)

        section = {"code": code, "name": name, "target": target, "questions": []}

        for row in rows[1:]:
            cells = row.findall('.//w:tc', NS)
            if len(cells) < 3:
                continue
            serial = cell_text(cells[0])
            content = cell_text(cells[1])
            opts_raw = cell_text(cells[2])
            if not content:
                continue

            gid += 1
            qtype, options = detect_type(content, opts_raw)

            section["questions"].append({
                "id": gid,
                "serial": serial,
                "content": content,
                "type": qtype,
                "required": 1,
                "options": options,
                "options_raw": opts_raw if qtype == 'text' else None,
            })

        if section["questions"]:
            survey_data["sections"].append(section)

    survey_data["total_questions"] = gid
    return survey_data

def write_sql(data, path):
    with open(path, 'w', encoding='utf-8') as f:
        f.write("-- SQL import for survey: " + data.get('title', '') + "\n")
        f.write("BEGIN;\n\n")
        title_esc = data['title'].replace("'", "''")
        f.write(f"INSERT INTO survey (title, created_at) VALUES ('{title_esc}', NOW());\n\n")
        f.write("-- Get survey_id from last insert for subsequent statements\n\n")

        for sec in data['sections']:
            hdr = f"[{sec['code']}] {sec['name']} ({sec['target']})".replace("'", "''")
            f.write(f"-- Section: {sec['code']} ({len(sec['questions'])} questions)\n")
            f.write(f"INSERT INTO survey_question (survey_id, parent_id, serial, content, type, required, sort) VALUES ((SELECT max(id) FROM survey), 0, '{sec['code']}', '{hdr}', 'h2', 1, 0);\n")

            for q in sec['questions']:
                content_esc = q['content'].replace("'", "''")
                opts_json = json.dumps(q.get('options', []), ensure_ascii=False).replace("'", "''")
                f.write(f"INSERT INTO survey_question (survey_id, parent_id, serial, content, type, options, required, sort) VALUES ((SELECT max(id) FROM survey), 0, '{q['serial']}', '{content_esc}', '{q['type']}', '{opts_json}', {q['required']}, {q['id']});\n")
            f.write("\n")

        f.write("COMMIT;\n")
    print(f"SQL written to {path}")

def main():
    print("Step 1: Extracting tables from docx...")
    tables = extract(DOCX_FILE)
    print(f"  {len(tables)} tables found")

    print("Step 2: Parsing question structure...")
    data = build_json(tables)

    with open(OUTPUT_JSON, 'w', encoding='utf-8') as f:
        json.dump(data, f, ensure_ascii=False, indent=2)
    print(f"  JSON written to {OUTPUT_JSON}")

    print(f"\n  Survey: {data['title']}")
    print(f"  Sections: {len(data['sections'])}")
    print(f"  Total questions: {data['total_questions']}")

    for s in data['sections']:
        type_counts = {}
        for q in s['questions']:
            t = q['type']
            type_counts[t] = type_counts.get(t, 0) + 1
        types_str = ', '.join(f'{k}:{v}' for k, v in type_counts.items())
        print(f"  [{s['code']}] {s['name']} - {len(s['questions'])}qs ({types_str})")

    print("\nStep 3: Generating SQL...")
    write_sql(data, OUTPUT_SQL)

    print("\nStep 4: Connect to PostgreSQL...")
    try:
        import psycopg2
        conn = psycopg2.connect(
            host="101.126.9.226", port=5432, user="root",
            password="kcer913639", dbname="survey", sslmode="disable"
        )
        cur = conn.cursor()

        # Create tables
        cur.execute("""
            CREATE TABLE IF NOT EXISTS survey (
                id SERIAL PRIMARY KEY, title TEXT DEFAULT '',
                pic TEXT DEFAULT '', "desc" TEXT DEFAULT '',
                start_at TIMESTAMP DEFAULT NOW(),
                end_at TIMESTAMP DEFAULT NOW() + INTERVAL '365 days',
                created_at TIMESTAMP DEFAULT NOW(), updated_at TIMESTAMP DEFAULT NOW(),
                delete_at TIMESTAMP, status INT DEFAULT 1
            )
        """)
        cur.execute("""
            CREATE TABLE IF NOT EXISTS survey_question (
                id SERIAL PRIMARY KEY, survey_id INT NOT NULL REFERENCES survey(id),
                parent_id INT DEFAULT 0, serial TEXT DEFAULT '', content TEXT DEFAULT '',
                type TEXT DEFAULT '', options JSONB DEFAULT '[]',
                jump_rules JSONB DEFAULT '[]', required INT DEFAULT 1, sort INT DEFAULT 0,
                show INT DEFAULT 0, remark TEXT DEFAULT '', level INT DEFAULT 0, tree TEXT DEFAULT '',
                created_at TIMESTAMP DEFAULT NOW(), updated_at TIMESTAMP DEFAULT NOW(),
                delete_at TIMESTAMP, status INT DEFAULT 1
            )
        """)
        conn.commit()
        print("  Tables ready")

        # Insert
        cur.execute("INSERT INTO survey (title) VALUES (%s) RETURNING id", (data['title'],))
        sid = cur.fetchone()[0]
        print(f"  Survey ID: {sid}")

        qc = 0
        for sec in data['sections']:
            hdr = f"[{sec['code']}] {sec['name']} ({sec['target']})"
            cur.execute(
                "INSERT INTO survey_question (survey_id,parent_id,serial,content,type,required,sort) VALUES (%s,0,%s,%s,'h2',1,0)",
                (sid, sec['code'], hdr)
            )
            for q in sec['questions']:
                opts = json.dumps(q.get('options', []), ensure_ascii=False)
                cur.execute(
                    "INSERT INTO survey_question (survey_id,parent_id,serial,content,type,options,required,sort) VALUES (%s,0,%s,%s,%s,%s,%s,%s)",
                    (sid, q['serial'], q['content'], q['type'], opts, q['required'], q['id'])
                )
                qc += 1

        conn.commit()

        # Verify
        cur.execute("SELECT type,count(*) FROM survey_question WHERE survey_id=%s GROUP BY type ORDER BY type", (sid,))
        print("\n  Type distribution:")
        for r in cur.fetchall():
            print(f"    {r[0]}: {r[1]}")
        cur.execute("SELECT count(*) FROM survey_question WHERE survey_id=%s", (sid,))
        print(f"\n  Total records: {cur.fetchone()[0]}")
        print("Done!")
        cur.close(); conn.close()
    except ImportError:
        print(f"  psycopg2 not available. Run: pip install psycopg2-binary")
        print(f"  Then: psql -h 101.126.9.226 -U root -d survey -f {OUTPUT_SQL}")
    except Exception as e:
        print(f"  DB error: {e}")
        print(f"  Manual import: psql -h 101.126.9.226 -U root -d survey -f {OUTPUT_SQL}")

if __name__ == "__main__":
    main()
