#!/usr/bin/env python3
"""
MySQL dump -> PostgreSQL converter v3.
Extracts CREATE TABLE column names, generates explicit-column INSERTs.
"""
import re

INPUT = "survey.sql"
OUTPUT = "survey_pg.sql"

PG_RESERVED = {"desc", "delete", "create", "update", "select", "insert", "from",
               "where", "order", "group", "table", "user", "password", "grant",
               "level", "show", "status"}

def main():
    with open(INPUT, 'r', encoding='utf-8', errors='replace') as f:
        content = f.read()

    # Step 1: Extract all CREATE TABLE -> map table name to column list
    tables = {}
    for m in re.finditer(r"CREATE TABLE\s+`?(\w+)`?\s*\((.*?)\)\s+ENGINE", content, re.DOTALL):
        tname = m.group(1).lower()
        body = m.group(2)
        cols = []
        for line in body.split('\n'):
            line = line.strip().rstrip(',')
            if not line:
                continue
            first = line.split()[0].strip('`"')
            if first.isidentifier() and first.upper() not in (
                'PRIMARY','KEY','CONSTRAINT','UNIQUE','FOREIGN','INDEX','CHECK','FULLTEXT'
            ):
                cols.append(first)
        if cols:
            tables[tname] = cols
            print(f"  {tname}: {len(cols)} columns")

    print(f"\n{len(tables)} tables found")

    with open(OUTPUT, 'w', encoding='utf-8') as fout:
        fout.write("BEGIN;\n\n")

        # Step 2: Drop and Create tables
        for tname, cols in tables.items():
            fout.write(f"DROP TABLE IF EXISTS {tname} CASCADE;\n")
            pg_cols = []
            for c in cols:
                c2 = c.lower()
                if c2 in PG_RESERVED:
                    c2 = f'"{c2}"'
                pg_cols.append(f"{c2} TEXT")
            fout.write(f"CREATE TABLE IF NOT EXISTS {tname} (\n  ")
            fout.write(",\n  ".join(pg_cols))
            fout.write("\n);\n\n")

        # Step 3: Extract and rewrite all INSERT statements
        # Pattern: INSERT INTO `table` VALUES (row1),(row2),... ;
        # The values may span multiple lines

        total = 0
        # Split content into chunks: each INSERT statement
        insert_pattern = re.compile(
            r"INSERT INTO\s+`?(\w+)`?\s+VALUES\s+(.+?);\s*(?:\n|$)",
            re.DOTALL | re.IGNORECASE
        )

        for m in insert_pattern.finditer(content):
            tname = m.group(1).lower()
            values_raw = m.group(2).strip()

            if tname not in tables:
                continue

            cols = tables[tname]

            # Quote reserved column names
            quoted_cols = []
            for c in cols:
                if c.lower() in PG_RESERVED:
                    quoted_cols.append(f'"{c}"')
                else:
                    quoted_cols.append(c)

            col_list = ", ".join(quoted_cols)

            # Parse value rows
            rows = extract_rows(values_raw)

            if not rows:
                continue

            # Write in batches of 200
            for batch_start in range(0, len(rows), 200):
                batch = rows[batch_start:batch_start+200]
                fout.write(f"INSERT INTO {tname} ({col_list}) VALUES\n")
                lines = []
                for r in batch:
                    vals = []
                    for v in r:
                        v = v.strip()
                        if v.upper() == 'NULL':
                            vals.append('NULL')
                        elif v.startswith("'") and v.endswith("'"):
                            inner = v[1:-1].replace("'", "''")
                            # Replace MySQL \r\n with literal newline
                            inner = inner.replace('\\r\\n', '\n').replace('\\n', '\n')
                            vals.append(f"E'{inner}'")
                        elif v.startswith('"') and v.endswith('"'):
                            inner = v[1:-1].replace('"', '""')
                            vals.append(f'"{inner}"')
                        else:
                            vals.append(v)
                    lines.append("  (" + ", ".join(vals) + ")")
                fout.write(",\n".join(lines))
                fout.write("\nON CONFLICT DO NOTHING;\n\n")
                total += len(batch)

        fout.write("COMMIT;\n")

    print(f"\nTotal INSERT rows: {total}")
    print(f"Output: {OUTPUT}")

def extract_rows(values_raw):
    """Parse (a,b,c),(d,e,f) into list of lists"""
    rows = []
    depth = 0
    in_str = False
    str_char = None
    escape_next = False
    current = ""

    i = 0
    while i < len(values_raw):
        c = values_raw[i]

        if escape_next:
            current += '\\' + c
            escape_next = False
            i += 1
            continue

        if c == '\\' and in_str:
            escape_next = True
            i += 1
            continue

        if in_str:
            current += c
            if c == str_char:
                # Check if it's escaped (double quote)
                if i+1 < len(values_raw) and values_raw[i+1] == str_char:
                    current += str_char
                    i += 2
                    continue
                in_str = False
            i += 1
            continue

        if c in ("'", '"'):
            in_str = True
            str_char = c
            current += c
            i += 1
            continue

        if c == '(':
            depth += 1
            if depth == 1:
                current = ""
            else:
                current += c
        elif c == ')':
            depth -= 1
            if depth == 0:
                # Split by comma (but respect nested parens and quotes)
                vals = split_csv(current)
                rows.append(vals)
                current = ""
            else:
                current += c
        elif depth > 0:
            current += c
        i += 1

    return rows

def split_csv(text):
    """Split CSV row while respecting quotes"""
    result = []
    current = ""
    in_str = False
    str_char = None
    depth = 0

    for c in text:
        if in_str:
            current += c
            if c == str_char:
                in_str = False
            continue
        if c in ("'", '"'):
            in_str = True
            str_char = c
            current += c
            continue
        if c == '(':
            depth += 1
            current += c
        elif c == ')':
            depth -= 1
            current += c
        elif c == ',' and depth == 0:
            result.append(current)
            current = ""
        else:
            current += c
    if current:
        result.append(current)
    return result

if __name__ == '__main__':
    main()