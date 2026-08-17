package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	dsn := "user=kcersing password=G7#kL2_mQ9$nR4&w host=pgm-bp14pne9o105t6x57o.pg.rds.aliyuncs.com port=5432 dbname=survey sslmode=disable TimeZone=Asia/Shanghai"
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		fmt.Printf("Failed to connect: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()
	db.Ping()
	fmt.Println("Connected")

	// --- Show current state ---
	printSection(db, 1, "A2")
	printSection(db, 2, "B3")

	// --- Shift existing A2-6..A2-12 to A2-9..A2-15 ---
	// Do it from high to low to avoid conflict
	for oldNum := 12; oldNum >= 6; oldNum-- {
		newNum := oldNum + 3
		db.Exec(
			"UPDATE survey_question SET serial=$1 WHERE survey_id=1 AND serial=$2",
			fmt.Sprintf("A2-%d", newNum), fmt.Sprintf("A2-%d", oldNum),
		)
	}

	// --- Shift existing B3-3..B3-6 to B3-5..B3-8 ---
	for oldNum := 6; oldNum >= 3; oldNum-- {
		newNum := oldNum + 2
		db.Exec(
			"UPDATE survey_question SET serial=$1 WHERE survey_id=2 AND serial=$2",
			fmt.Sprintf("B3-%d", newNum), fmt.Sprintf("B3-%d", oldNum),
		)
	}

	// --- Insert new A2-6, A2-7, A2-8 ---
	insertQuestion(db, 1, Question{
		Serial: "A2-6", Content: "您是否签约家庭医生？", Type: "single_choice",
		Options: []Option{
			{Content: "是", Serial: 1},
			{Content: "否（跳转至A2-9）", Serial: 2},
			{Content: "不清楚（跳转至A2-9）", Serial: 3},
		},
	})
	insertQuestion(db, 1, Question{
		Serial: "A2-7", Content: "过去一年内，家庭医生是否主动联系过您并开展过相应服务？", Type: "single_choice",
		Options: []Option{
			{Content: "是，___次", Serial: 1},
			{Content: "否（跳转至A2-9）", Serial: 2},
		},
	})
	insertQuestion(db, 1, Question{
		Serial: "A2-8", Content: "家庭医生为您提供过哪些服务？（可多选）", Type: "multiple_choice",
		Options: []Option{
			{Content: "测量血压血糖等基础检查", Serial: 1},
			{Content: "慢性病随访", Serial: 2},
			{Content: "健康咨询", Serial: 3},
			{Content: "康复护理", Serial: 4},
			{Content: "用药指导", Serial: 5},
			{Content: "转诊服务", Serial: 6},
			{Content: "中医药服务", Serial: 7},
			{Content: "上门诊疗", Serial: 8},
			{Content: "其他", Serial: 9},
		},
	})

	// --- Insert new B3-3, B3-4 ---
	insertQuestion(db, 2, Question{
		Serial: "B3-3", Content: "您接受的培训类型是？", Type: "single_choice",
		Options: []Option{
			{Content: "线上学习", Serial: 1},
			{Content: "村里培训", Serial: 2},
			{Content: "医院/机构培训", Serial: 3},
			{Content: "入户培训", Serial: 4},
			{Content: "其他培训", Serial: 5},
		},
	})
	insertQuestion(db, 2, Question{
		Serial: "B3-4", Content: "您接受的培训内容是？（可多选）", Type: "multiple_choice",
		Options: []Option{
			{Content: "基础生活照护", Serial: 1},
			{Content: "失能护理技能", Serial: 2},
			{Content: "康复训练方法", Serial: 3},
			{Content: "慢病管理", Serial: 4},
			{Content: "应急处理知识", Serial: 5},
			{Content: "营养与膳食管理", Serial: 6},
			{Content: "心理与情绪疏导", Serial: 7},
			{Content: "用药管理与安全指导", Serial: 8},
			{Content: "安宁疗护与临终关怀", Serial: 9},
			{Content: "其他", Serial: 10},
		},
	})

	fmt.Println("\n--- After insertion ---")
	printSection(db, 1, "A2")
	printSection(db, 2, "B3")

	// Count totals
	for _, sid := range []int{1, 2, 3} {
		var cnt int
		db.QueryRow("SELECT count(*) FROM survey_question WHERE survey_id=$1 AND type!='h2'", sid).Scan(&cnt)
		fmt.Printf("Survey #%d: %d questions\n", sid, cnt)
	}

	fmt.Println("Done!")
}

type Question struct {
	Serial   string   `json:"serial"`
	Content  string   `json:"content"`
	Type     string   `json:"type"`
	Options  []Option `json:"options"`
}

type Option struct {
	Content string `json:"content"`
	Serial  int    `json:"serial"`
}

func insertQuestion(db *sql.DB, surveyID int, q Question) {
	opts, _ := json.Marshal(q.Options)
	_, err := db.Exec(
		"INSERT INTO survey_question (survey_id, parent_id, serial, content, type, options, required, sort) VALUES ($1, 0, $2, $3, $4, $5, 1, 0)",
		surveyID, q.Serial, q.Content, q.Type, string(opts),
	)
	if err != nil {
		fmt.Printf("FAILED %s: %v\n", q.Serial, err)
	} else {
		fmt.Printf("OK %s: %s\n", q.Serial, q.Content[:min(50, len(q.Content))])
	}
}

func printSection(db *sql.DB, surveyID int, section string) {
	rows, _ := db.Query(
		"SELECT serial, content, type FROM survey_question WHERE survey_id=$1 AND (serial LIKE $2 OR serial LIKE $3) ORDER BY serial",
		surveyID, section+"-%", section+"-%",
	)
	defer rows.Close()
	fmt.Printf("\n--- Survey #%d, Section %s ---\n", surveyID, section)
	for rows.Next() {
		var s, c, t string
		rows.Scan(&s, &c, &t)
		fmt.Printf("  %s: %s (%s)\n", s, c[:min(55, len(c))], t)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}