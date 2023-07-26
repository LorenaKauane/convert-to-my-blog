/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"convert-to-my-blog/cmd"
	config "convert-to-my-blog/configs"
)
func main() {
	config.StartConfig()
	cmd.Execute()
}

func main2() {
	config.StartConfig()
	cmd.Execute()
}
