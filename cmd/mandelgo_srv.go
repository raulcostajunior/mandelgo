package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/raulcostajunior/mandelgo"
)

func main() {

	if len(os.Args) > 2 {
		printUsage()
		os.Exit(1)
	}

	var portNum = 8080
	if len(os.Args) == 2 {
		var err error
		if portNum, err = strconv.Atoi(os.Args[1]); err != nil {
			fmt.Println()
			fmt.Printf("mandelgo_srv - error in command line: '%s' is not a valid port number.\n", os.Args[1])
			printUsage()
			os.Exit(1)
		}
	}
	fmt.Println()
	fmt.Printf("mandelgo_srv - running web server on port %d...\n", portNum)
	fmt.Println()
	fmt.Println("Press <Ctrl> + <C> to stop the mandelgo server.")
	runServer(portNum)

}

func printUsage() {
	fmt.Println()
	fmt.Println("mandelgo_srv - Mandelbot Set Image Server")
	fmt.Println()
	fmt.Println("Usage: mandelgo_srv [port]")
	fmt.Println()
	fmt.Println("    port(8080) - port number the web server will be listening at (optional).")
	fmt.Println()
	fmt.Println("Launches a web server that replies to requests for Mandelbrot set images in PNG format.")
	fmt.Println()
	fmt.Println("The parameters for generating the image must be sent as query parameters of the HTTP request.")
	fmt.Println("For details about the parameters, take a look at the \"mandelgo.GenerateImage\" function.")
	fmt.Println()
}

func runServer(port int) {
	mandelgo.GenerateImage(1024, 1024, -2, -2, 2, 2)
}
