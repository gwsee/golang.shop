package utils

import (
	"strings"
	"time"
)

func GetTimeStr(step, symbol string) (str string) {
	mapDate := make(map[string]string)
	mapDate["y"] = time.Now().Format("2006")
	mapDate["m"] = time.Now().Format("01")
	mapDate["d"] = time.Now().Format("02")
	mapDate["h"] = time.Now().Format("15")
	mapDate["i"] = time.Now().Format("04")
	mapDate["s"] = time.Now().Format("05")
	if step == "" {
		step = "ymd"
	}
	ret := strings.Split(step, "")
	var strSlice []string
	for _, v := range ret {
		if mapDate[v] != "" {
			strSlice = append(strSlice, mapDate[v])
		}
	}
	str = strings.Join(strSlice, symbol)
	return
}
