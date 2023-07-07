package blind

func isAnagram(s string, t string) bool {
	sM := make(map[rune]int)

	for _, sc := range s {
		sM[rune(sc)] += 1
	}

	for _, tc := range t {
		if _, exists := sM[rune(tc)]; exists {
			sM[rune(tc)] -= 1
		} else {
			return false
		}
	}

	for _, v := range sM {
		if v != 0 {
			return false
		}
	}

	return true
}

// 输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
// 输出: [["bat"],["nat","tan"],["ate","eat","tea"]]

func groupAnagram(strs []string) [][]string {
	res := [][]string{}

	for len(strs) > 0 {
		group := []string{strs[0]}
		for j := 1; j < len(strs); j++ {
			if isAnagram(strs[0], strs[j]) {
				group = append(group, strs[j])
				strs = append(strs[:j], strs[j+1:]...)
			}
		}

		res = append(res, group)
	}

	return res
}
