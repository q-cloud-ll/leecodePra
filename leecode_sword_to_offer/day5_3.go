package leecode_sword_to_offer

func firstUniqChar(s string) byte {
	strMap := make(map[byte]int)
	strArr := []byte(s)
	for _, v := range strArr {
		strMap[v]++
	}
	for _, v := range strArr {
		if strMap[v] == 1 {
			return v
		}
	}
	return ' '
}
