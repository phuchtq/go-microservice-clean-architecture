package utils

func GenerateCallBackUrl(sq []string, seperateChar string) string {
	if len(sq) < 1 {
		return ""
	}

	var res string = ""
	for i, v := range sq {
		res += v
		if i < len(sq)-1 && i > 0 {
			res += seperateChar
		}
	}

	return res
}
