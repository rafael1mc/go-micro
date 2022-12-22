package main

import (
	"log"
	"log-service/data"
	"time"
)

type RPCServer struct {
	app *Config
}

type RPCPayload struct {
	Name string
	Data string
}

func (r *RPCServer) LogInfo(payload RPCPayload, resp *string) error {
	logEntry := data.LogEntry{
		Name:      payload.Name,
		Data:      payload.Data,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Reusing our previous Insert
	err := r.app.Models.LogEntry.Insert(logEntry)
	if err != nil {
		log.Println("error writing to mongo", err)
		return err
	}

	/*
		// Without access to Config.Models
		collection := client.Database("logs").Collection("logs")
		_, err := collection.InsertOne(context.TODO(), logEntry)
		if err != nil {
			log.Println("error writing to mongo", err)
			return err
		}
	*/

	*resp = "Processed payload via RPC: " + payload.Name

	return nil
}
