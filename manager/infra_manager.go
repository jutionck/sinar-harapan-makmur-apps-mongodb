package manager

import (
	"context"
	"fmt"
	"github.com/jutionck/go-sinar-makmur-mongodb/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/url"
)

// local (ApplyURI): mongodb://localhost:27017
// mongosh: mongodb+srv://jutionck:<password>@itdpsmm.senzmxh.mongodb.net/?retryWrites=true&w=majority

type InfraManager interface {
	Conn() *mongo.Database
}

type infraManager struct {
	db  *mongo.Database
	cfg *config.Config
}

func (i *infraManager) Conn() *mongo.Database {
	return i.db
}

func (i *infraManager) initDb() error {
	dsn := fmt.Sprintf("mongodb+srv://%s:%s@itdpsmm.senzmxh.mongodb.net/?retryWrites=true&w=majority", i.cfg.User, url.QueryEscape(i.cfg.Password))
	clientOpts := options.Client()
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOpts.ApplyURI(dsn).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.Background(), clientOpts)
	if err != nil {
		return err
	}
	if err := client.Database("sinar_makmur_platform").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	// kita membutuhkan database mana yang mau digunakan
	i.db = client.Database(i.cfg.Name)

	return nil
}

func NewInfraManager(cfg *config.Config) (InfraManager, error) {
	conn := &infraManager{cfg: cfg}
	err := conn.initDb()
	if err != nil {
		return nil, err
	}
	return conn, nil
}
