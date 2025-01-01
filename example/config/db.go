package config

import (
	"github.com/gocroot/mgdb"
)

var MongoString string = "http://localhost:27017"

var mongoinfo = mgdb.DBInfo{
	DBString: MongoString,
	DBName:   "example",
}

var Mongoconn, ErrorMongoconn = mgdb.MongoConnect(mongoinfo)
