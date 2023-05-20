package asciiart

import (
	"fmt"
	"os"
	"strings"
)

func AsciiArt(text, style string) (string, error) {
	data, err := os.ReadFile("asciiart/ascii-style/" + style + ".txt")
	if err != nil {
		fmt.Println("There is something wrong with the ascii-style")
		return "", nil
	}
	// if len(os.Args) != 2 {
	// 	fmt.Println("Sorry! The program accepts only one argument")
	// 	return "", nil
	// }
	CheckData(data) // проверка на пустоту файла
	data1 := strings.ReplaceAll(string(data), "\\n", "\n")
	//args := strings.Join(os.Args[1:], " ")
	if len(text) > 49 {
		fmt.Println("Too many items\nMaximum number 49")
		return "", nil
	}
	res := strings.Split(string(data1), "\n")

	ChmoMap := MakeMap(res)
	chmotext := strings.Split(strings.ReplaceAll(text, "\\n", "\n"), "\n")

	output := PrintMap(chmotext, ChmoMap)
	fmt.Println(output)
	return output, err
}

// А ПОТОМ МЫ ЧЕРЕЗ ГЕТ ЕГО ОТСЫЛАЛИ ВО ВТОРУЮ ФОРМУ
