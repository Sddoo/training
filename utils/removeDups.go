package utils

func removeDups(strs []string) []string {
	for i := 0; i < len(strs) - 1; i++ {
		if strs[i] == strs[i + 1] {
			for j := 0; j < len(strs) - i - 2; j++ {
				strs[i + j] = strs[i + j + 2]
			}
			strs = strs[:len(strs) - 2]
			strs = removeDups(strs)
		}
	}
	return strs
}
