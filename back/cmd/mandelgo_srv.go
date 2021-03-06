package main

import (
	"bytes"
	"fmt"
	"image/png"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
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
	fmt.Printf("mandelgo_srv - Mandelbrot Set Image Viewer on port %d...\n", portNum)
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
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		// Retrieves any of the supported path parameters.
		width := 1200
		height := 800
		xmin := -2.0
		ymin := -2.0
		xmax := 2.0
		ymax := 2.0
		var csval int
		var err error
		if width, err = strconv.Atoi(c.Query("width", "1024")); err != nil {
			log.Printf("Invalid width given: %s\n", c.Query("width"))
		}
		if height, err = strconv.Atoi(c.Query("height", "1024")); err != nil {
			log.Printf("Invalid height given: %s\n", c.Query("height"))
		}
		if xmin, err = strconv.ParseFloat(c.Query("xmin", "-2"), 64); err != nil {
			log.Printf("Invalid xmin given: %s\n", c.Query("xmin"))
		}
		if ymin, err = strconv.ParseFloat(c.Query("ymin", "-2"), 64); err != nil {
			log.Printf("Invalid ymin given: %s\n", c.Query("ymin"))
		}
		if xmax, err = strconv.ParseFloat(c.Query("xmax", "2"), 64); err != nil {
			log.Printf("Invalid xmax given: %s\n", c.Query("xmax"))
		}
		if ymax, err = strconv.ParseFloat(c.Query("ymax", "2"), 64); err != nil {
			log.Printf("Invalid xmax given: %s\n", c.Query("ymax"))
		}
		if csval, err = strconv.Atoi(c.Query("colorScheme", "2")); err != nil {
			log.Printf("Invalid ColorScheme given: %s\n", c.Query("colorScheme"))
		}
		cs := mandelgo.ColorSchemeFromValue(csval)

		img := mandelgo.GenerateImage(width, height, xmin, ymin, xmax, ymax, cs)
		var pngimg bytes.Buffer
		png.Encode(&pngimg, img)

		c.Set("Content-Type", "image/png")
		return c.Send(pngimg.Bytes())
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))

}
