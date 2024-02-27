package config

import (
	"context"
	"log"
	"myauth/ent"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

func ConnectMySQL() {
	client, err := ent.Open("mysql", "root:autocrypt@tcp(127.0.0.1:3306)/my_auth")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer func() {
		if err := client.Close(); err != nil {
			logrus.WithError(err).Error("failed to close db")
		}
	}()
}

func CreateSchema(client *ent.Client) {
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
