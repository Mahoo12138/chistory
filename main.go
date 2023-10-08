package main

import "github.com/mahoo12138/chistory/cmd"

func main() {
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
