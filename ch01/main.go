package main

import "fmt"
func main() {
	cmd :=parseCmd()
	if cmd.versionFlag{
		fmt.Printf("version 0.0.1\n")
	} else if cmd.helpFlag || cmd.class== ""{
		printUsage()
	} else {
		startJVM(cmd)
	}

}
