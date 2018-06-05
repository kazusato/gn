package main

import (
	"bufio"
	"os"
	"fmt"
	"gn/gnclient"
	"log"
)

func main() {
	config, err := gnclient.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	gn := gnclient.NewClientFromConfig(config)
	gn.Connect()

	input := bufio.NewScanner(os.Stdin)
	fmt.Print("GN> ")
	for input.Scan() {
		line := input.Text()
		if line == `\q` {
			os.Exit(0)
		} else {
			resp, err := gn.SendCommand(line)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error occurred during executing command: " + line)
			}
			fmt.Println(resp)

			fmt.Print("GN> ")
		}
	}
}