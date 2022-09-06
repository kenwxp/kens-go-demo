package user

import (
	"context"
	"kens/demo/storage"
	"kens/demo/storage/types"
	"kens/demo/util"
	"kens/demo/util/enty_logger"
	"strconv"
	"strings"
	"time"
)

const checkInterval = time.Minute * 5
const logonTimeOut = time.Minute * 10

func CheckUserLogonTask(db *storage.Database) {
	go autoCheckUserLogon(db)
}

func autoCheckUserLogon(db *storage.Database) {
	for {
		checkUserLogon(db)
		time.Sleep(checkInterval)
	}
}

func checkUserLogon(db *storage.Database) {
	ctx := context.TODO()
	enty_logger.Info("execute task: checkUserLogon start...")
	start := time.Now()
	nowTs := util.TimeNow().Unix()
	userList, err := db.SelectUserListWithCondition(ctx, nil, types.User{})
	for _, user := range userList {
		if user.Token != "" {
			mid := strings.Split(user.Token, ":")
			timeStamp := mid[1]
			preTs, _ := strconv.ParseInt(timeStamp, 10, 64)
			diff := nowTs - preTs
			if time.Duration(diff)*time.Second > logonTimeOut {
				err = db.UpdateUserLogon(ctx, nil, user.AccNo, "0")
			} else {
				err = db.UpdateUserLogon(ctx, nil, user.AccNo, "1")
			}
		}
	}
	if err != nil {
		enty_logger.Info(err.Error())
		panic(err)
	}
	enty_logger.Info("finish task: checkUserLogon , cost:", time.Since(start).String())
}
