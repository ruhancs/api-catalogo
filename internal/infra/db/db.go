package db

import (
	"os"

	"github.com/elastic/go-elasticsearch/v8"
)

func NewElaticConn() (*elasticsearch.Client, error) {
	//client,err := elasticsearch.NewClient(elasticsearch.Config{
	//	CloudID: "",
	//	APIKey: "",
	//})
	cert, err := os.ReadFile(os.Getenv("ES_CERT"))
	if err != nil {
		panic(err)
	}

	cfg := elasticsearch.Config{
		Addresses: []string{
			"https://localhost:9200",
		},
		Username: "elastic",
		Password: os.Getenv("ES_PASS"),
		CACert:   cert,
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	_, err = client.Info()
	if err != nil {
		return nil, err
	}

	return client, nil
}
