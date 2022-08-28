package main

import (
	"fmt"
	"os"

	"github.com/mk-55/go-ruby/ruby"
)

func main() {
	//TODO 現状は未完成で単純に引数を変換するだけ。オプション引数などは要検討
	fmt.Printf("args : %#v\n", os.Args)
	src := os.Args[1]

	conf, err := ruby.NewConfiguration(ruby.RUBY_PATTERN_HTML)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	writer, err := ruby.NewRubyWriter(conf)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	output, err := writer.AddRuby(src)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	fmt.Println(output)
}
