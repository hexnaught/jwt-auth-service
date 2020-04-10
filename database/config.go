package database

// MongoConfig contains some configuration data for forming a connection
type MongoConfig struct {
	Host string // Host IP/URI to connect via
	Port string // Port to connect via

	Username   string // Mongo user username
	Password   string // Mongo user password
	AuthSource string // Database name holding above user credentials

	DatabaseName string // Database name to connect to/use
}
