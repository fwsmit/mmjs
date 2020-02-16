package database

import (
	"database/sql"
	"mp3bak2/globals"
)

var dbc *sql.DB

// Warmup the mysql connection pool
func Warmup() *sql.DB {
	db, err := sql.Open("mysql",
		globals.DatabaseCredentials.Username+":"+
			globals.DatabaseCredentials.Password+"@/"+
			globals.DatabaseCredentials.Database)

	if err != nil {
		panic("cannot open database")
	}

	err = db.Ping()
	if err != nil {
		panic("connection with database could not be established")
	}

	dbc = db
	return dbc
}

func getConnection() *sql.DB {
	return dbc
}

// StringToSQLNullableString : Convert a string into a nullable string
func StringToSQLNullableString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{String: s, Valid: false}
	}
	return sql.NullString{String: s, Valid: true}
}

// IntToSQLNullableInt : Convert a int into a nullable int
func IntToSQLNullableInt(s int) sql.NullInt64 {
	var i = int64(s)
	if s == 0 {
		return sql.NullInt64{Int64: i, Valid: false}
	}
	return sql.NullInt64{Int64: i, Valid: true}
}
