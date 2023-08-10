package data

import "database/sql"

// AESKey is the key used to encrypt and decrypt data
var AESKey = []byte("0123456789abcdef")

var OS string

// TempSQLiteConnects is a map of SQLite database connections
var TempSQLiteConnects = make(map[string]*sql.DB)
