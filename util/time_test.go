package util

import (
	"fmt"
	"testing"
)

const TIME_NOW = "2021-09-06 14:24:30"
const TIME_NOW_LONDON = "2021-09-06 5:58:30"

func TestTimeZone(t *testing.T) {
	//now := util.TimeNow().Unix()
	//fmt.Println(now)
	//ts := TimeStrToUnix(TIME_NOW)
	//fmt.Println(now, ts, (now-ts)/60/60)
	//zoneTs := TimeStrToUnix(TIME_NOW)

	//now := util.TimeNow()
	//nowTs := now.Unix()
	//fmt.Println(nowTs)
	//t2, err := time.Parse("2006/01/02 15:04:05.000000", now.Format("2006/01/02 15:04:05.000000"))
	//if err != nil {
	//	panic("")
	//}
	//fmt.Println(t2.Unix())
	//GetWeekMonday()

	//a := int64(14415616)
	//b := int64(1000)
	//
	//diffRate := float64(a) / float64(b)
	//c := strconv.FormatFloat(diffRate, 'f', 2, 64)
	//
	//fmt.Println(c)
	fmt.Println(TimeStrToUnix("2022-08-01 17:00:01"))
	fmt.Println(TimeStrToUnixStr("2022-08-01 17:00:01"))
}
