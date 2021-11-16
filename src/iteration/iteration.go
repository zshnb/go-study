package iteration

func repeat(singleChar string, times int) string {
	result := ""
	for i := 0;i < times;i++ {
		result += singleChar
	}
	return result
}
