package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	host := "pgm-bp14pne9o105t6x57o.pg.rds.aliyuncs.com"
	user := "kcersing"
	pass := "G7#kL2_mQ9$nR4&w"
	if len(os.Args) > 1 {
		host = os.Args[1]
	}
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=5432 dbname=survey sslmode=disable TimeZone=Asia/Shanghai", user, pass, host)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		fmt.Printf("Connect failed: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		fmt.Printf("Ping failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Connected to PostgreSQL")

	file, err := os.Open("../survey_pg.sql")
	if err != nil {
		fmt.Printf("Open failed: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 1024*1024), 10*1024*1024) // 10MB buffer for large INSERTs

	var stmt strings.Builder
	total := 0
	ok := 0
	failed := 0
	skipped := 0

	for scanner.Scan() {
		line := scanner.Text()

		// Skip comments and empty lines within statements
		if strings.HasPrefix(strings.TrimSpace(line), "--") {
			continue
		}

		stmt.WriteString(line)
		stmt.WriteString("\n")

		// Check if statement is complete (ends with semicolon)
		trimmed := strings.TrimSpace(line)
		if strings.HasSuffix(trimmed, ";") {
			sql := strings.TrimSpace(stmt.String())
			stmt.Reset()
			total++

			if sql == "" || sql == ";" {
				skipped++
				continue
			}

			// Skip BEGIN/COMMIT (we handle transactions differently)
			if sql == "BEGIN;" {
				skipped++
				continue
			}

			_, err := db.Exec(sql)
			if err != nil {
				errStr := err.Error()
				// Acceptable errors
				if strings.Contains(errStr, "already exists") ||
					strings.Contains(errStr, "does not exist") ||
					strings.Contains(errStr, "duplicate key") ||
					strings.Contains(errStr, "violates foreign key") {
					skipped++
					if total%50 == 0 {
						fmt.Printf("  [skip] %s...\n", sql[:min(60, len(sql))])
					}
				} else {
					failed++
					fmt.Printf("  [FAIL #%d] %s\n", total, errStr)
					fmt.Printf("  SQL: %s\n", sql[:min(200, len(sql))])
				}
			} else {
				ok++
				if total%100 == 0 {
					fmt.Printf("  [ok #%d] processed\n", total)
				}
			}
		}
	}

	fmt.Printf("\nDone: %d total, %d ok, %d skipped, %d failed\n", total, ok, skipped, failed)

	// Show table counts
	rows, _ := db.Query(`
		SELECT table_name FROM information_schema.tables
		WHERE table_schema='public' AND table_type='BASE TABLE'
		ORDER BY table_name
	`)
	defer rows.Close()
	fmt.Println("\nTables with row counts:")
	for rows.Next() {
		var t string
		rows.Scan(&t)
		var cnt int
		db.QueryRow(fmt.Sprintf("SELECT count(*) FROM %s", t)).Scan(&cnt)
		fmt.Printf("  %-30s %d rows\n", t, cnt)
	}
}

func min(a, b int) int {
	if a < b { return a }
	return b
}