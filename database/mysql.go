package database

var connection string

// init otomatis dieksekusi pertama kali saat package dipanggil
func init() {
	connection = "MySql"
}

func GetDatabase() string {
	return connection
}
