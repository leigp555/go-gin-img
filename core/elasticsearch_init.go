package core

import (
	es "github.com/elastic/go-elasticsearch/v7"
	"img/server/global"
)

func LinkElasticsearch() {
	client, err := es.NewClient(es.Config{
		Addresses: []string{global.Config.Elasticsearch.Addr},
		Username:  global.Config.Elasticsearch.Username,
		Password:  global.Config.Elasticsearch.Password,
	})
	if err != nil {
		global.Slog.Fatalf("elasticsearch连接失败%v\n", err)
	}
	global.Elastic = client
	global.Slog.Info("elasticsearch连接成功")
}
