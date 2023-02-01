package main

import (
	"fmt"
	"os"
	"github.com/artukshini/WebStressTest/app"
	)

func help() {
	fmt.Println("WebStressTest <URL> <COUNT>\nExample:\nWebStressTest https://www.testime.siguria.ne.internetPROJEKTI-III.com 1000")
}

func main() {

	if len(os.Args) < 2 {
		help()
		os.Exit(0)
	}
	
		app.MainStress(os.Args[1], os.Args[2])
}
