package utils

//字符串处理类
func BuildUnique(pre, sign string) (str string) {
	str = pre + "_" + sign + GetTimeStr("ymdhis", "") + "_" + Krand(4, 3)
	return
}
