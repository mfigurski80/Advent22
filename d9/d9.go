package d9

import (
	"fmt"

	. "github.com/mfigurski80/AOC22/utils"
)

func Main() {
	DoByFileLine("d9/d9.txt", func(line string) {
		fmt.Println(line)
	})
}
