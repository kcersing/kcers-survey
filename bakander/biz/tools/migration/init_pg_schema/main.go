package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	dsn := "user=kcersing password=G7#kL2_mQ9$nR4&w host=pgm-bp14pne9o105t6x57o.pg.rds.aliyuncs.com port=5432 dbname=survey sslmode=disable TimeZone=Asia/Shanghai"
	db, _ := sql.Open("pgx", dsn)
	defer db.Close()
	if err := db.Ping(); err != nil {
		fmt.Printf("Ping failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Connected")

	// Recreate survey table with proper types
	db.Exec("DROP TABLE IF EXISTS survey CASCADE")
	db.Exec(`CREATE TABLE survey (
		id SERIAL PRIMARY KEY,
		created_at TIMESTAMP DEFAULT NOW(),
		updated_at TIMESTAMP DEFAULT NOW(),
		"delete" BIGINT DEFAULT 0,
		created_id BIGINT DEFAULT 0,
		status BIGINT DEFAULT 1,
		title TEXT DEFAULT '',
		pic TEXT DEFAULT '',
		"desc" TEXT DEFAULT '',
		start_at TIMESTAMP DEFAULT NOW(),
		end_at TIMESTAMP DEFAULT NOW() + INTERVAL '365 days'
	)`)
	fmt.Println("survey table recreated")

	// Recreate survey_question table with proper types
	db.Exec("DROP TABLE IF EXISTS survey_question CASCADE")
	db.Exec(`CREATE TABLE survey_question (
		id SERIAL PRIMARY KEY,
		created_at TIMESTAMP DEFAULT NOW(),
		updated_at TIMESTAMP DEFAULT NOW(),
		"delete" BIGINT DEFAULT 0,
		created_id BIGINT DEFAULT 0,
		status BIGINT DEFAULT 1,
		parent_id BIGINT DEFAULT 0,
		content TEXT DEFAULT '',
		type VARCHAR(255) DEFAULT '',
		sort BIGINT DEFAULT 0,
		required BIGINT DEFAULT 1,
		survey_id BIGINT DEFAULT 0,
		serial VARCHAR(255) DEFAULT '',
		options JSONB DEFAULT '[]',
		jump_rules JSONB DEFAULT '[]',
		show BIGINT DEFAULT 0,
		remark TEXT DEFAULT '',
		"level" BIGINT DEFAULT 0,
		tree TEXT DEFAULT ''
	)`)
	fmt.Println("survey_question table recreated")

	// Keep survey_response and survey_response_answers (936K rows from MySQL import)
	fmt.Println("Done! (survey_response tables preserved)")
}