package main

import (
	"fmt"
	"kens/demo/util"
	"math/big"
	"testing"
)

func TestTime(t *testing.T) {
	//没生效
	//data := util.HandleXP("2021-09-02", 3, 365, "2021-09-02")
	//fmt.Println("没生效", data)
	//fmt.Println()
	//
	////生效 过去1个月
	//data2 := util.HandleXP("2021-08-02", 3, 365, "2021-09-02")
	//fmt.Println("生效 过去1个月", data2)
	//fmt.Println()
	//
	////生效过去 1年
	//data3 := util.HandleXP("2020-09-02", 3, 365, "2021-09-02")
	//fmt.Println("生效过去 1年", data3)
	//fmt.Println()
	//
	////生效 过去 2年
	//data4 := util.HandleXP("2019-09-02", 3, 365, "2021-09-02")
	//fmt.Println("生效 过去 2年:", data4)
	//fmt.Println()
	//1000000000

	balance := "1000000000"
	extractNum, err1 := util.Digit("0.11", "mul9")
	lock := "0"
	fmt.Println(err1)
	balanceBig, bool1 := new(big.Float).SetString(balance)
	extractNumBig, bool2 := new(big.Float).SetString(extractNum)
	lockBig, bool3 := new(big.Float).SetString(lock)
	fmt.Println(bool1, bool2, bool3)

	if i := balanceBig.Cmp(extractNumBig); i == -1 {
		fmt.Println("提现不足")
	}

	updateBalanceBig := new(big.Float).Sub(balanceBig, extractNumBig)
	updateLockBig := new(big.Float).Add(lockBig, extractNumBig)

	fmt.Println("balanceBig:", balanceBig.String())
	fmt.Println("extractNumBig:", extractNumBig.String())
	fmt.Println("updateBalanceBig:", updateBalanceBig.String())
	fmt.Println("lockBig:", lockBig.String())
	fmt.Println("updateLockBig:", updateLockBig.String())

	fmt.Println("。。。")

	//str, err := util.Digit(extractNum, "mul9")
	//fmt.Println(str,err)
	//new(big.Float).Mul(bf, unit)

}
