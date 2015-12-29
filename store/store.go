package store

func InitStore() {
	initDatabase()
	initRedis()
}

func Close() {
	Database.Close()
	Redis.Close()
}
