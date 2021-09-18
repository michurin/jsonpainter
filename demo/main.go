package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/michurin/jsonpainter"
)

func main() {
	opts := []jsonpainter.Option{
		jsonpainter.ClrCtl(jsonpainter.Blue),
		jsonpainter.ClrKey(jsonpainter.Green),
		jsonpainter.ClrStr(jsonpainter.Pink),
	}
	r := bufio.NewReader(os.Stdin)
	for {
		s, err := r.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Printf("Original:    %s", s)
		fmt.Printf("Highlighted: %s", jsonpainter.String(s))
		fmt.Printf("Customized:  %s", jsonpainter.String(s, opts...))
	}
}
