package inner

import (
	"fmt"
	"os"
)

func RequestOuterFile() {

	file, err := os.Open("a.txt")

	fmt.Println("a: ", err)

	file.Close()

}
