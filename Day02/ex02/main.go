package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"log"
	"os/exec"
	"strings"
)

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(input), "\n")
	flag.Parse()
	for i := range lines {
		if len(flag.Args()) > 0 {
			cmd := exec.Command(flag.Args()[0], append(flag.Args()[1:], lines[i])...)
			out, err := cmd.CombinedOutput()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s", out)
		}
	}
}