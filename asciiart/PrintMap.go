package asciiart

func PrintMap(chmotext []string, ChmoMap map[rune][]string) string {
	count := 0
	res := ""
	for _, value := range chmotext {
		if len(value) > 0 {
			count++
		}
	}
	if count == 0 {
		for i := 0; i < len(chmotext)-1; i++ {
			// fmt.Println()
			res += "\n"
		}
		return ""
	}
	for _, arg := range chmotext {
		for _, v := range arg {
			if v > 126 || v < 32 {
				// fmt.Println("ERROR: Need ASCII character")
				return "ERROR: Need ASCII character"
			}
		}
		if len(arg) > 0 {
			for i := 0; i < 8; i++ {
				for _, e2 := range arg {
					// fmt.Print(ChmoMap[rune(e2)][i])
					res += string(ChmoMap[rune(e2)][i])
				}
				// fmt.Println()
				res += "\n"
			}
		} else if len(chmotext) == 1 && len(arg) == 0 {
			return ""
		} else {
			// fmt.Println()
			res += "\n"
		}
	}
	return res
}
