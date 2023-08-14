package data

import "database/sql"

// AESKey is the key used to encrypt and decrypt data
// It must be 16 bytes long
var AESKey = []byte("yxfr2023skyline0")

var OS string

// TempSQLiteConnects is a map of SQLite database connections
var TempSQLiteConnects = make(map[string]*sql.DB)

var DefaultAvatar string
var DefaultBackgroundImage string
var DefaultSignature string
