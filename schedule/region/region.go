package region

import (
	"context"
	"kens/demo/cache"
	"kens/demo/storage"
	"kens/demo/util/enty_logger"
	"time"
)

type SearchType int
type SearchLevel int

const (
	CacheRegionNodeKey  = "region_tree"
	SearchById          = SearchType(0)
	SearchByCode        = SearchType(1)
	SearchByName        = SearchType(2)
	SearchSelf          = SearchLevel(0)
	SearchChild         = SearchLevel(1)
	SearchParent        = SearchLevel(2)
	CacheRegionInterval = time.Minute * 10
)

func CacheRegionTask(db *storage.Database) {
	go autoUpdateRegionCache(db)
}

func autoUpdateRegionCache(db *storage.Database) {
	for {
		RefreshRegionCache(db)
		time.Sleep(CacheRegionInterval)
	}
}

func RefreshRegionCache(db *storage.Database) {
	enty_logger.Info("execute task: refreshRegionCache start...")
	start := time.Now()
	paramList, err := db.SelectRegionList(context.TODO(), nil)
	if err != nil {
		enty_logger.Info("SelectRegionList err:", err)
		panic(err)
	}
	regionTree := Node{
		Id:    "0000000",
		Name:  "中国",
		Level: "0",
		Code:  "0",
	}
	child, err := genNode(paramList, regionTree, 3)
	regionTree.Child = child
	if err != nil {
		enty_logger.Info("genNode err:", err)
		panic(err)
	}
	cache.CacheAll.Set(CacheRegionNodeKey, regionTree, CacheRegionInterval)
	enty_logger.Info("finish task: refreshRegionCache , cost:", time.Since(start).String())
}
