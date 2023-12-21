package utils

func anagrams(s1, s2 string) bool {
	if (len(s1) != len(s2)) {
		return false
	}
	var i int
	var b2 byte
	buf1 := []byte(s1);
	buf2 := []byte(s2);
	for _, b1 := range buf1 {
		for i, b2 = range buf2 {
			if b2 == b1 {
				buf2[i] = 0
				break
			}
		}
		if i == len(buf2) - 1 && buf2[i] != 0 {
			return false
		}
	}
	return true
}
