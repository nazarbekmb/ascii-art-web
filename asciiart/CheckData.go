package asciiart

import (
	"fmt"
)

func CheckData(data []byte) {
	if string(data) == "" {
		fmt.Print("The file is empty ")
		return
	}
}
