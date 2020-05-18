package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	homeEnv := os.Getenv("HOME")

	if _, err := os.Stat(homeEnv + "/.zsh_history"); err != nil && os.IsNotExist(err) {
		fmt.Println("fatal: zsh history file is not exist")
		return
	}

	bs, err := ioutil.ReadFile(homeEnv + "/.zsh_history")
	if err != nil {
		fmt.Println("fatal: read zsh history file failed")
		return
	}

	for line := range strings.Split(string(bs), "\n") {
		fmt.Println(line, 1)
	}

}
