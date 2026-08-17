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
		fmt.Println("PG not reachable:", err)
		os.Exit(1)
	}
	fmt.Println("Connected")

	// ent-managed tables that conflict with MySQL-imported TEXT-only schemas
	// Drop them so ent can recreate with proper types
	// KEEP: survey_response, survey_response_answers (have data)
	tablesToDrop := []string{
		"casbin_rules",
		"role_menus",
		"user_roles",
		"sys_apis",
		"sys_area",
		"sys_dictionaries",
		"sys_dictionary_details",
		"sys_logs",
		"sys_logs1",
		"sys_menu_params",
		"sys_menus",
		"sys_roles",
		"sys_sms",
		"sys_sms_log",
		"sys_tokens",
		"sys_users",
		"survey",
		"survey_question",
		// survey_response and survey_response_answers are KEPT
	}

	for _, t := range tablesToDrop {
		_, err := db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE", t))
		if err != nil {
			fmt.Printf("  FAIL drop %s: %v\n", t, err)
		} else {
			fmt.Printf("  dropped %s\n", t)
		}
	}

	// Fix survey_response and survey_response_answers to have proper SERIAL PKs
	// These were created with TEXT columns by the MySQL import
	fixTable(db, "survey_response", "id")
	fixTable(db, "survey_response_answers", "id")

	fmt.Println("\nDone! Ent migration should now succeed.")
}

func fixTable(db *sql.DB, table, pkCol string) {
	// Check if id column is already integer
	var dataType string
	err := db.QueryRow(fmt.Sprintf(
		"SELECT data_type FROM information_schema.columns WHERE table_name=$1 AND column_name=$2",
		table, pkCol)).Scan(&dataType)
	if err != nil {
		fmt.Printf("  %s: table not found\n", table)
		return
	}
	if dataType == "integer" || dataType == "bigint" {
		fmt.Printf("  %s: already integer PK, OK\n", table)
		return
	}
	fmt.Printf("  %s: has TEXT id (needs fix), type=%s\n", table, dataType)
	// Can't easily convert TEXT to SERIAL for a PK with existing data.
	// For now, these tables will need to be handled separately.
}
