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
	caddy.RegisterEventHook("ondemandcertobtained", onDemandCertObtained)
	caddy.RegisterEventHook("caddydb-cert-failure", onDemandCertFailure)
}

var dbConn *mongo.Client


//OnDemandCertFailure Called when Caddy fails to obtain a certificate for a given host
func onDemandCertFailure(eventType caddy.EventName, eventInfo interface{}) error {
	if eventType != caddy.OnDemandCertFailureEvent {
		// Only listen to the event we are interested in
		return nil
	}

	// Interface containing data about a failed on demand certificate
	type CertFailureData struct {
		Name   string
		Reason error
	}

	data := eventInfo.(CertFailureData)

	fmt.Println("FAILED", data.Name, data.Reason)
	return nil
}

//OnDemandCertObtained Called when Caddy obtains a certificate for a given host
func onDemandCertObtained(eventType caddy.EventName, eventInfo interface{}) error {
	if eventType != caddy.OnDemandCertObtainedEvent {
		// Only listen to the event we are interested in
		return nil
	}
	fmt.Println("SUC", eventInfo.(string))
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
