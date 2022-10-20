package main

import (
	"bufio"
	"fmt"
	"os"
	"syscall"

	"golang.org/x/term"

	"github.com/mk-55/go-ruby/ruby"
)

func main() {
	//TODO 現状は未完成で単純に引数を変換するだけ。オプション引数などは要検討
	fmt.Printf("args : %#v\n", os.Args)

	var src string
	if len(os.Args) > 2 {
		src = os.Args[1]
	} else {
		src = ""
	}

	var sc *bufio.Scanner
	switch src {
	case "":
		// ターミナル入力でsrc未指定なら終了
		if term.IsTerminal(syscall.Stdin) {
			os.Exit(1)
		}
		fallthrough
	case "-":
		sc = bufio.NewScanner(os.Stdin)
	default:

	}

	conf, err := ruby.NewConfiguration(ruby.RUBY_PATTERN_HTML)
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}

	writer, err := ruby.NewRubyWriter(conf)
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}

	if sc != nil {
		for sc.Scan() {
			src = sc.Text()
			output, err := writer.AddRuby(src)
			if err != nil {
				fmt.Printf(err.Error())
				os.Exit(1)
			}
			fmt.Println(output)
		}
	} else {
		output, err := writer.AddRuby(src)
		if err != nil {
			fmt.Printf(err.Error())
			os.Exit(1)
		}
		fmt.Println(output)
	}
	os.Exit(0)
}
