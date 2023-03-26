package arraysandhashing

func IsAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}
	var f [26]int

	for i := 0; i < len([]rune(s)); i++ {
		f[s[i]-'a']++
		f[t[i]-'a']--
	}

	for _, x := range f {
		if x != 0 {
			return false
		}
	}

	return true
}
