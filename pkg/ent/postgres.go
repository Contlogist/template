package ent

import (
	"context"
	"git.legchelife.ru/gitlab-instance-7d441567/catalog_m/config"
	"git.legchelife.ru/gitlab-instance-7d441567/catalog_m/ent"
	_ "github.com/lib/pq"
	"log"
)

func NewPostgresClient(pgConfig config.EntPG, migration bool) (*ent.Client, error) {
	log.Println("Connecting to PostgreSQL...")
	log.Println("pgConfig: ", pgConfig)
	client, err := ent.Open(
		"postgres",
		"host="+pgConfig.Host+" port="+pgConfig.Port+" user="+pgConfig.User+" dbname="+pgConfig.Database+" password="+pgConfig.Password+" sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if migration {
		if err := client.Schema.Create(context.Background()); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
	}

	return client, nil
}
