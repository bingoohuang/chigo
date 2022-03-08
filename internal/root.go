package internal

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/UltiRequiem/chigo"
	"io"
	"os"
)

func StartProcessFromStdin() {
	reader := bufio.NewReader(os.Stdin)

	for i := 1.0; ; i++ {
		input, _, err := reader.ReadRune()
		if err != nil {
			if !errors.Is(err, io.EOF) {
				chigo.PrintWithColors(err.Error())
			}
			break
		}

		rgb := chigo.NewRGB(i)
		fmt.Printf("%s", rgb.Color(string(input)))
	}
}
