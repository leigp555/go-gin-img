package core

import (
	"img/server/models"
	"sync"
)

var wg sync.WaitGroup

func InitDeps() {
	InitConf()
	InitLogger()
	wg.Add(3)
	go func() {
		LinkMysqlDB()
		models.CreateTables()
		defer wg.Done()
	}()
	go func() {
		LinkRedisDB()
		defer wg.Done()
	}()
	go func() {
		LinkElasticsearch()
		defer wg.Done()
	}()
	wg.Wait()
}
