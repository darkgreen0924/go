package _case

import (
	"log"
	"os"
)

func init() {
	log.SetFlags(log.Lshortfile)
	log.SetOutput(os.Stderr)
}

func LogCase() {
	var a, b = -1, -2
	_, err := sum(a, b)
	if err != nil {
		log.Println(err)
	}
	log.Printf("aaaaaaaaaaawqwqqqwwwww")
}
