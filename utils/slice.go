package utils

//切片判断是否包含
func NumInSlice(num uint64, arr []uint64) (flag bool) {
	flag = false
	for _, v := range arr {
		if v == num {
			flag = true
			break
		}
	}
	return
}

//切片去重

func UniqueSlice(arr []uint64) (res []uint64) {
	all := make(map[uint64]int)
	for _, v := range arr {
		all[v] = 1
	}
	for k, _ := range all {
		res = append(res, k)
	}
	return
}
