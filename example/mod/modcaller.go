package mod

import (
	"github.com/gocroot/example/mod/idgrup"
	"github.com/whatsauth/itmodel"
	"go.mongodb.org/mongo-driver/mongo"
)

func Caller(Profile itmodel.Profile, Modulename string, Pesan itmodel.IteungMessage, db *mongo.Database) (reply string) {
	if Modulename == "idgrup" {
		reply = idgrup.IDGroup(Pesan)
	}
	return
}