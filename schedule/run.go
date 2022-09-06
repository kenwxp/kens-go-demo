package schedule

import (
	"kens/demo/schedule/region"
	"kens/demo/schedule/user"
	"kens/demo/storage"
)

func Run(db *storage.Database) {
	region.CacheRegionTask(db)
	user.CheckUserLogonTask(db)
}
