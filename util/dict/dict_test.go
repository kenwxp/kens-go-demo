package dict

import (
	"testing"
	"time"
)

func TestDict(t *testing.T) {
	//start := time.Now()
	//raw := int64(0)
	//convert := int64(0)
	//for {
	//	raw = raw + 1
	//	rawStr := strconv.FormatInt(raw, 10)
	//	convertStr := rawStr[1:] + rawStr[0:1]
	//	convert, _ = strconv.ParseInt(convertStr, 10, 64)
	//	if convert == 3*raw {
	//		break
	//	}
	//}
	//elapsed := time.Since(start) * time.Nanosecond
	//println(raw)
	//println(convert)
	//println(elapsed.String())
	test(1)
	defer test(2)
	time.Sleep(3 * time.Second)
}

func test(i int) {
	defer println(i, "test defer")
	println(i, "test")
}
