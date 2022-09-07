package cfg

import "os"

var SECURE_KEY string

var DB_HOST string
var DB_PORT string
var DB_USERNAME string
var DB_PASSWORD string
var DB_DATABASE string

func InitCfg() {
	SECURE_KEY = getEnv("SECURE_KEY", "gkaGHJwdaAAGHBWDBOUWADOadinfai1231nuio931789132g8731jp4")

	DB_HOST = getEnv("DB_HOST", "localhost")
	DB_PORT = getEnv("DB_PORT", "27017")
	DB_USERNAME = getEnv("DB_USERNAME", "preppy")
	DB_PASSWORD = getEnv("DB_PASSWORD", "preppy-password")
	DB_DATABASE = getEnv("DB_DATABASE", "preppy")
}

func getEnv(v string, def string) string {
	val, ok := os.LookupEnv(v)
	if !ok {
		return def
	}
	return val
}
