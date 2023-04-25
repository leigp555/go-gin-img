package test

import (
	es "github.com/elastic/go-elasticsearch/v7"
	"log"
	"testing"
)

var (
	client *es.Client
)

func init() {
	var err error
	client, err = es.NewClient(es.Config{
		Addresses: []string{"http://1.117.141.66:9200"},
		Username:  "elastic",
		Password:  "123456abc",
	})
	if err != nil {
		log.Fatal(err)
	}
}

func TestNewESClient(t *testing.T) {
	t.Log(client.Info())
}

//创建索引
