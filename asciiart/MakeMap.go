package asciiart

func MakeMap(res []string) map[rune][]string {
	ChmoMap := make(map[rune][]string)
	count := 1
	for i := ' '; i <= '~'; i++ {
		ChmoMap[i] = res[count : count+8]
		count += 9
	}
	return ChmoMap
}
