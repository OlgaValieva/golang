package main

import (
	"os"
	"fmt"
	"flag"
	"strings"
	"path/filepath"	
)

var flags Flags

type Flags struct {
	sl, d, f, ex bool
	dir, ext     string
}

func main() {
	checkFlag(&flags)
	input, err := os.Open(flags.dir)
	if flags.dir == "" || err != nil  {
		fmt.Println(err)
		os.Exit(1)
	}
	err = input.Close()
	filepath.Walk(flags.dir, findFunc)
}

func checkFlag(flags *Flags) {
	flag.BoolVar(&flags.sl, "sl", false, "symlink")
	flag.BoolVar(&flags.d, "d", false, "directory")
	flag.BoolVar(&flags.f, "f", false, "file")
	flag.BoolVar(&flags.ex, "ext", false, "extension")
	flag.Parse()
	if flags.ex == true && flags.f == false {
		fmt.Println("works ONLY when -f is specified")
		os.Exit(1)
	} else if flags.f == true && flags.ex == true {
		if len(flag.Args()) == 2 {
			flags.ext = flag.Args()[0]
			flags.dir = flag.Args()[1]
		} else {
			fmt.Println("need flags: -sl, -d or -f [-ext] and ./directory")
			os.Exit(1)
		}
	} else {
		if len(flag.Args()) == 1 {
			flags.dir = flag.Args()[0]
		} else {
			fmt.Println("need flags: -sl, -d or -f [-ext] and ./directory")
			os.Exit(1)
		}
	}
	if flags.sl == false && flags.d == false && flags.f == false {
		flags.sl = true
		flags.d = true
		flags.f = true
	}
}

func checkFile(file string) bool {
	if file[0] == '/' {
		return file[1] == '.'
	}
	return file[0] == '.'
}

func findFunc(s string, info os.FileInfo, err error) error {
	if err == nil {
		newS := strings.TrimPrefix(s, flags.dir)
		if newS != "" && !checkFile(newS) {
			if info.Mode()&(1<<2) != 0 {
				if flags.sl && info.Mode().Type() == os.ModeSymlink {
					realS, err := filepath.EvalSymlinks(s)
					if err != nil {
						fmt.Println(s, "-> [broken]")
					} else {
						fmt.Println(s, "->", realS)
					}
				} else if flags.d && info.IsDir() {
					fmt.Println(s)
				} else if flags.f && info.Mode().IsRegular() {
					if flags.ex == false || flags.ex == true && filepath.Ext(s) == "." + flags.ext {
						fmt.Println(s)
					}
				}
			}
		}
	}
	return nil
}
