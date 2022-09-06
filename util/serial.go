package util

import (
	"math/rand"
)

func GenTaskSerial() string {
	return "TK" + RandomNum(6) + TimeNow().Format("20060102150405")
}

func GenEventSerial() string {
	return "EV" + RandomNum(6) + TimeNow().Format("20060102150405")
}

func GenTransSerial() string {
	return "TR" + RandomNum(6) + TimeNow().Format("20060102150405")
}

func GenBillSerial() string {
	return "BL" + RandomNum(6) + TimeNow().Format("20060102150405")
}

const randomRanges = "0123456789"

// RandomString generates a pseudo-random string of length n.
func RandomNum(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = randomRanges[rand.Int63()%int64(len(randomRanges))]
	}
	return string(b)
}
