package main

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var dsn = "user=kcersing password=G7#kL2_mQ9$nR4&w host=pgm-bp14pne9o105t6x57o.pg.rds.aliyuncs.com port=5432 dbname=survey sslmode=disable TimeZone=Asia/Shanghai"

func main() {
	db, _ := sql.Open("pgx", dsn)
	defer db.Close()
	db.Ping()
	fmt.Println("Connected to RDS")

	// Step 0: Ensure proper schema
	db.Exec("DROP TABLE IF EXISTS survey CASCADE")
	db.Exec(`CREATE TABLE survey (id SERIAL PRIMARY KEY, created_at TIMESTAMP DEFAULT NOW(), updated_at TIMESTAMP DEFAULT NOW(), "delete" BIGINT DEFAULT 0, created_id BIGINT DEFAULT 0, status BIGINT DEFAULT 1, title TEXT DEFAULT '', pic TEXT DEFAULT '', "desc" TEXT DEFAULT '', start_at TIMESTAMP DEFAULT NOW(), end_at TIMESTAMP DEFAULT NOW() + INTERVAL '365 days')`)
	db.Exec("DROP TABLE IF EXISTS survey_question CASCADE")
	db.Exec(`CREATE TABLE survey_question (id SERIAL PRIMARY KEY, created_at TIMESTAMP DEFAULT NOW(), updated_at TIMESTAMP DEFAULT NOW(), "delete" BIGINT DEFAULT 0, created_id BIGINT DEFAULT 0, status BIGINT DEFAULT 1, parent_id BIGINT DEFAULT 0, content TEXT DEFAULT '', type VARCHAR(255) DEFAULT '', sort BIGINT DEFAULT 0, required BIGINT DEFAULT 1, survey_id BIGINT DEFAULT 0, serial VARCHAR(255) DEFAULT '', options JSONB DEFAULT '[]', jump_rules JSONB DEFAULT '[]', show BIGINT DEFAULT 0, remark TEXT DEFAULT '', "level" BIGINT DEFAULT 0, tree TEXT DEFAULT '')`)
	fmt.Println("Schema fixed")

	// Step 1: Re-insert MySQL survey data (the 3 surveys + 500 questions that were deleted)
	fmt.Println("\n[1] Re-inserting MySQL survey data...")
	execSQLFile(db, "../survey_pg.sql", "survey")
	execSQLFile(db, "../survey_pg.sql", "survey_question")

	// Reset sequences to avoid duplicate key errors
	db.Exec("SELECT setval('survey_id_seq', (SELECT COALESCE(max(id),0) FROM survey))")
	db.Exec("SELECT setval('survey_question_id_seq', (SELECT COALESCE(max(id),0) FROM survey_question))")

	// Show current survey IDs
	rows, _ := db.Query("SELECT id, title FROM survey ORDER BY id")
	defer rows.Close()
	fmt.Println("\n  Current surveys:")
	for rows.Next() {
		var id int; var t string
		rows.Scan(&id, &t)
		fmt.Printf("  #%d: %s\n", id, t[:min(50, len(t))])
	}

	// Step 2: Import docx surveys
	fmt.Println("\n[2] Importing docx survey data...")
	importDocxSurveys(db)

	// Step 3: Add extra questions
	fmt.Println("\n[3] Adding A2/B3 extra questions...")
	addExtraQuestions(db)

	// Final summary
	fmt.Println("\n====================")
	rows2, _ := db.Query("SELECT s.id, s.title, count(q.id) FROM survey s LEFT JOIN survey_question q ON q.survey_id=s.id AND q.type!='h2' GROUP BY s.id, s.title ORDER BY s.id")
	defer rows2.Close()
	for rows2.Next() {
		var id int; var t string; var c int
		rows2.Scan(&id, &t, &c)
		fmt.Printf("Survey #%d: %s — %d questions\n", id, t[:min(60, len(t))], c)
	}
	fmt.Println("Done!")
}

func execSQLFile(db *sql.DB, filepath, tableFilter string) {
	f, _ := os.Open(filepath)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, 1024*1024), 10*1024*1024)

	var stmt strings.Builder
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(strings.TrimSpace(line), "--") { continue }
		stmt.WriteString(line + "\n")
		if strings.HasSuffix(strings.TrimSpace(line), ";") {
			sql := strings.TrimSpace(stmt.String())
			stmt.Reset()
			prefix := "insert into " + tableFilter + " "
		if !strings.HasPrefix(strings.ToLower(sql), prefix) { continue }
			if strings.Contains(sql, "ON CONFLICT") {
				_, err := db.Exec(sql)
				if err != nil { fmt.Printf("  FAIL: %v\n", err) } else { count++ }
			}
		}
	}
	fmt.Printf("  %s: %d INSERTs\n", tableFilter, count)
}

func importDocxSurveys(db *sql.DB) {
	data, _ := os.ReadFile("../_survey_data.json")
	var d struct {
		Title    string `json:"title"`
		Sections []struct {
			Code   string `json:"code"`
			Name   string `json:"name"`
			Target string `json:"target"`
			Questions []struct {
				ID      int    `json:"id"`
				Serial  string `json:"serial"`
				Content string `json:"content"`
				Type    string `json:"type"`
				Required int   `json:"required"`
				Options []struct {
					Content string `json:"content"`
					Serial  int    `json:"serial"`
				} `json:"options"`
			} `json:"questions"`
		} `json:"sections"`
	}
	json.Unmarshal(data, &d)

	groups := []struct{ Code, Suffix, Title string }{
		{"A", "A", d.Title + " — A卷 农村老年人"},
		{"B", "B", d.Title + " — B卷 家庭照护者"},
		{"C", "C", d.Title + " — C卷 村庄"},
	}
	for _, g := range groups {
		qc := 0
		for _, s := range d.Sections {
			if strings.HasPrefix(s.Code, g.Suffix) {
				qc += len(s.Questions)
			}
		}

		var sid int
		err := db.QueryRow("INSERT INTO survey (title, created_at) VALUES ($1, $2) RETURNING id", g.Title, time.Now()).Scan(&sid)
		if err != nil { fmt.Printf("  FAIL survey %s: %v\n", g.Code, err); continue }
		fmt.Printf("  Survey #%d: %s (%d qs)\n", sid, g.Title, qc)

		sortOrder := 0
		for _, sec := range d.Sections {
			if !strings.HasPrefix(sec.Code, g.Suffix) { continue }
			hdr := fmt.Sprintf("[%s] %s", sec.Code, sec.Name)
			db.Exec("INSERT INTO survey_question (survey_id,parent_id,serial,content,type,required,sort) VALUES ($1,0,$2,$3,'h2',1,$4)", sid, sec.Code, hdr, sortOrder)
			sortOrder++
			for i, q := range sec.Questions {
				opts, _ := json.Marshal(q.Options)
				db.Exec("INSERT INTO survey_question (survey_id,parent_id,serial,content,type,options,required,sort) VALUES ($1,0,$2,$3,$4,$5,$6,$7)", sid, q.Serial, q.Content, q.Type, string(opts), q.Required, sortOrder*100+i)
			}
			fmt.Printf("    [%s] %d qs\n", sec.Code, len(sec.Questions))
		}
	}
}

func addExtraQuestions(db *sql.DB) {
	// A2: shift A2-6..A2-12 to A2-9..A2-15, insert new A2-6/7/8
	// Find survey #1 A卷 (first survey with "A卷" in title)
	var aSid int
	db.QueryRow("SELECT id FROM survey WHERE title LIKE '%A卷%' LIMIT 1").Scan(&aSid)
	if aSid == 0 { fmt.Println("  A卷 survey not found"); return }

	// Shift existing
	for oldNum := 12; oldNum >= 6; oldNum-- {
		db.Exec("UPDATE survey_question SET serial=$1 WHERE survey_id=$2 AND serial=$3",
			fmt.Sprintf("A2-%d", oldNum+3), aSid, fmt.Sprintf("A2-%d", oldNum))
	}

	// Insert new A2-6/7/8
	insertQ(db, aSid, "A2-6", "您是否签约家庭医生？", "single_choice", `[{"content":"是","serial":1},{"content":"否（跳转至A2-9）","serial":2},{"content":"不清楚（跳转至A2-9）","serial":3}]`)
	insertQ(db, aSid, "A2-7", "过去一年内，家庭医生是否主动联系过您并开展过相应服务？", "single_choice", `[{"content":"是，___次","serial":1},{"content":"否（跳转至A2-9）","serial":2}]`)
	insertQ(db, aSid, "A2-8", "家庭医生为您提供过哪些服务？（可多选）", "multiple_choice", `[{"content":"测量血压血糖等基础检查","serial":1},{"content":"慢性病随访","serial":2},{"content":"健康咨询","serial":3},{"content":"康复护理","serial":4},{"content":"用药指导","serial":5},{"content":"转诊服务","serial":6},{"content":"中医药服务","serial":7},{"content":"上门诊疗","serial":8},{"content":"其他","serial":9}]`)

	// B3: find B卷 survey
	var bSid int
	db.QueryRow("SELECT id FROM survey WHERE title LIKE '%B卷%' LIMIT 1").Scan(&bSid)
	if bSid == 0 { fmt.Println("  B卷 survey not found"); return }

	// Update B3-3/4 content and options
	db.Exec("UPDATE survey_question SET content='您接受的培训类型是？',type='single_choice',options=$1 WHERE survey_id=$2 AND serial='B3-3'",
		`[{"content":"线上学习","serial":1},{"content":"村里培训","serial":2},{"content":"医院/机构培训","serial":3},{"content":"入户培训","serial":4},{"content":"其他培训","serial":5}]`, bSid)
	db.Exec("UPDATE survey_question SET content='您接受的培训内容是？（可多选）',type='multiple_choice',options=$1 WHERE survey_id=$2 AND serial='B3-4'",
		`[{"content":"基础生活照护","serial":1},{"content":"失能护理技能","serial":2},{"content":"康复训练方法","serial":3},{"content":"慢病管理","serial":4},{"content":"应急处理知识","serial":5},{"content":"营养与膳食管理","serial":6},{"content":"心理与情绪疏导","serial":7},{"content":"用药管理与安全指导","serial":8},{"content":"安宁疗护与临终关怀","serial":9},{"content":"其他","serial":10}]`, bSid)

	fmt.Println("  A2-6/7/8 inserted, B3-3/4 updated")
}

func insertQ(db *sql.DB, sid int, serial, content, qtype, opts string) {
	_, err := db.Exec("INSERT INTO survey_question (survey_id,parent_id,serial,content,type,options,required,sort) VALUES ($1,0,$2,$3,$4,$5,1,0)",
		sid, serial, content, qtype, opts)
	if err != nil { fmt.Printf("  FAIL %s: %v\n", serial, err) } else { fmt.Printf("  OK %s\n", serial) }
}

func min(a, b int) int { if a < b { return a }; return b }
