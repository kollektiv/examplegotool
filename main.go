package main

import (
	"log"
	"runtime/debug"
)

func main() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		log.Fatalln("could not reading build info")
	}

	log.Println("main module version:", info.Main.Version)
}
