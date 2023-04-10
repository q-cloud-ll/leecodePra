package leecode_sword_to_offer

func reverseLeftWords(s string, n int) string {
	str := []byte(s)
	reverse(str[n:])
	reverse(str[:n])
	reverse(str)
	return string(str)
}

func reverse(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		temp := s[len(s)-1-i]
		s[len(s)-1-i] = s[i]
		s[i] = temp
	}
}
