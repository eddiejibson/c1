package certdb

import (
	"context"
	"fmt"
	"log"

	"github.com/caddyserver/caddy"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Cert - Successful cert struct for putting into collection
type Cert struct {
	Name string
	Time string
}

func init() {
	caddy.RegisterPlugin("certdb", caddy.Plugin{
		ServerType: "http",
		Action:     Setup,
	})
	caddy.RegisterEventHook("ondemandcertobtained", CertObtained)
	caddy.RegisterEventHook("ondemandcertfailure", CertFailed)
}

var dbConn *mongo.Client

//CertObtained - Called when a certificate for a domain has been obtained su
func CertObtained(event caddy.EventName, info interface{}) error {
	fmt.Printf("%v", info)
	return nil
}

//CertFailed - guess what this does
func CertFailed(event caddy.EventName, info interface{}) error {
	fmt.Printf("%v", info)
	return nil
}

//Setup function
func Setup(c *caddy.Controller) error {
	if !c.NextArg() {
		return c.ArgErr()
	}
	value := c.Val()
	if len(value) <= 0 {
		value = "mongodb://mongodb:27017"
	}
	clientOptions := options.Client().ApplyURI(value)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the cert Database")
	return nil
}
