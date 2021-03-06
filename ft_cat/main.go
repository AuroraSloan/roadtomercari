package main

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var exit_status = 0

func main() {
	if len(os.Args) < 2 {
		ft_cat(os.Stdin, os.Stdout)
	}
	for _, filename := range os.Args[1:] {
		if filename == "-" {
			ft_cat(os.Stdin, os.Stdout)
		} else {
			src, err := os.Open(filename)
			if err != nil {
				print_err(err.Error())
				continue
			}
			ft_cat(src, os.Stdout)
			src.Close()
		}
	}
	os.Exit(exit_status)
}

func ft_cat(src io.Reader, dst io.Writer) {
	reader := bufio.NewReader(src)
	var err error
	for err != io.EOF {
		bytes, err := reader.ReadBytes('\n')
		if err != nil {
			if err != io.EOF {
				print_err(err.Error())
			}
			break
		}
		dst.Write(bytes)
	}
}

func print_err(err string) {
	exit_status = 1
	error_message := filepath.Base(os.Args[0]) + ":"
	words := strings.Fields(err)
	for i, word := range words[1:] {
		if i == 1 {
			word = strings.Title(word)
		}
		error_message += " " + word
	}
	io.WriteString(os.Stderr, error_message+"\n")
}
