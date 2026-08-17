package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type SurveyData struct {
	Title          string    `json:"title"`
	Sections       []Section `json:"sections"`
	TotalQuestions int       `json:"total_questions"`
}

type Section struct {
	Code      string     `json:"code"`
	Name      string     `json:"name"`
	Target    string     `json:"target"`
	Questions []Question `json:"questions"`
}

type Question struct {
	ID         int      `json:"id"`
	Serial     string   `json:"serial"`
	Content    string   `json:"content"`
	Type       string   `json:"type"`
	Required   int      `json:"required"`
	Options    []Option `json:"options"`
	OptionsRaw *string  `json:"options_raw"`
}

type Option struct {
	Content string `json:"content"`
	Serial  int    `json:"serial"`
}

// SurveyGroup defines a survey grouping: which sections belong together
type SurveyGroup struct {
	Code    string // A, B, C
	Title   string // Full survey title
	Suffix  string // Section prefix (A, B, C)
	Sections []Section
}

func main() {
	dsn := "user=kcersing password=G7#kL2_mQ9$nR4&w host=pgm-bp14pne9o105t6x57o.pg.rds.aliyuncs.com port=5432 dbname=survey sslmode=disable TimeZone=Asia/Shanghai"
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		fmt.Printf("Failed to connect: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		fmt.Printf("Failed to ping: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Connected to PostgreSQL")

	// Clean old data
	db.Exec("DELETE FROM survey_question WHERE survey_id IN (SELECT id FROM survey)")
	db.Exec("DELETE FROM survey")
	db.Exec("ALTER SEQUENCE survey_id_seq RESTART WITH 1")
	db.Exec("ALTER SEQUENCE survey_question_id_seq RESTART WITH 1")
	fmt.Println("Cleaned old data")

	// Read JSON
	data, err := os.ReadFile("../_survey_data.json")
	if err != nil {
		fmt.Printf("Failed to read JSON: %v\n", err)
		os.Exit(1)
	}
	var allData SurveyData
	if err := json.Unmarshal(data, &allData); err != nil {
		fmt.Printf("Failed to parse JSON: %v\n", err)
		os.Exit(1)
	}

	// Split sections into 3 survey groups
	groups := []SurveyGroup{
		{Code: "A", Suffix: "A", Title: allData.Title + " — A卷 农村老年人"},
		{Code: "B", Suffix: "B", Title: allData.Title + " — B卷 家庭照护者"},
		{Code: "C", Suffix: "C", Title: allData.Title + " — C卷 村庄"},
	}

	for _, sec := range allData.Sections {
		prefix := string(sec.Code[0]) // A, B, or C
		for i := range groups {
			if groups[i].Suffix == prefix {
				groups[i].Sections = append(groups[i].Sections, sec)
				break
			}
		}
	}

	totalQ := 0
	for _, g := range groups {
		qCount := 0
		for _, s := range g.Sections {
			qCount += len(s.Questions)
		}

		// Insert survey
		var sid int
		err = db.QueryRow(
			"INSERT INTO survey (title, created_at) VALUES ($1, $2) RETURNING id",
			g.Title, time.Now(),
		).Scan(&sid)
		if err != nil {
			fmt.Printf("Failed to insert survey %s: %v\n", g.Code, err)
			continue
		}
		fmt.Printf("\nSurvey #%d: %s (%d questions)\n", sid, g.Title, qCount)

		// Insert sections and questions
		sortOrder := 0
		for _, sec := range g.Sections {
			// Section header
			hdr := fmt.Sprintf("[%s] %s", sec.Code, sec.Name)
			db.Exec(
				"INSERT INTO survey_question (survey_id, parent_id, serial, content, type, required, sort) VALUES ($1, 0, $2, $3, 'h2', 1, $4)",
				sid, sec.Code, hdr, sortOrder,
			)
			sortOrder++

			// Questions
			subSort := 0
			for _, q := range sec.Questions {
				optsJSON, _ := json.Marshal(q.Options)
				_, err := db.Exec(
					"INSERT INTO survey_question (survey_id, parent_id, serial, content, type, options, required, sort) VALUES ($1, 0, $2, $3, $4, $5, $6, $7)",
					sid, q.Serial, q.Content, q.Type, string(optsJSON), q.Required, sortOrder*100+subSort,
				)
				if err != nil {
					fmt.Printf("  FAIL %s: %v\n", q.Serial, err)
				}
				subSort++
			}
			fmt.Printf("  [%s] %s — %d qs\n", sec.Code, sec.Name, len(sec.Questions))
		}
		totalQ += qCount
	}

	// Summary
	fmt.Println("\n====================")
	rows, _ := db.Query("SELECT s.id, s.title, count(q.id) FROM survey s LEFT JOIN survey_question q ON q.survey_id=s.id WHERE q.type!='h2' GROUP BY s.id, s.title ORDER BY s.id")
	defer rows.Close()
	for rows.Next() {
		var id int
		var title string
		var cnt int
		rows.Scan(&id, &title, &cnt)
		fmt.Printf("Survey #%d: %s — %d questions\n", id, title, cnt)
	}
	fmt.Printf("\nTotal: %d questions across 3 surveys\n", totalQ)
	fmt.Println("Done!")
}
