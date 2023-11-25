package db

// TODO: define mongo schema.
const (
	MongoURI        = "mongodb://admin:password@localhost:27017"
	DB              = "SingboxConvertor"
	UserCollection  = "users"
	DNSCollection   = "DNS"
	RouteCollection = "routes"
	DNSConfig       = "./db/config/DNS.json"
	RouteConfig     = "./db/config/route.json"
)
