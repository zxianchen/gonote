package main

import "gonote/cmd"

func main() {
	defer cmd.Clear()
	cmd.Start()
}
