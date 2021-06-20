package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
	"sync"

	halftone "github.com/alkazarix/go-halftone"
)

var (
	outputDir string
	methods   string
)

func usage() {

	const helper = `
		Usage:
			go run <image>
		Options:
  	-method string
    	name of the filters to apply, options: 'bayer', 'bayer2', 'bayer4' (default bayer2)
  	-output string
			Directory name, where to save the generated images (default "output")`

	fmt.Println(helper)
	os.Exit(1)
}

func main() {

	if err := process(); err != nil {
		fmt.Printf("%s", err)
		usage()
	}

	fmt.Printf("DONE")
}

func process() error {

	command := flag.NewFlagSet("command", flag.ExitOnError)
	command.StringVar(&outputDir, "output", "output", "directory where to save generated image")
	command.StringVar(&methods, "method", "bayer2", "filter to apply (default: bayer2)")
	command.Usage = usage

	if len(os.Args) < 2 {
		usage()
	}

	command.Parse(os.Args[2:])
	img, err := openImage(os.Args[1])
	if err != nil {
		return err
	}

	filters, err := parseMethod(methods)
	if err != nil {
		return err
	}

	_ = os.Mkdir(outputDir, os.ModePerm)

	var wg sync.WaitGroup
	for _, filter := range filters {
		wg.Add(1)
		go func(filter *halftone.Filter) {
			defer wg.Done()
			generatedImage := halftone.Halftone(img, *filter)
			generatedFilename := filepath.Join(outputDir, filter.Name+".png")
			writePng(generatedFilename, generatedImage)
		}(filter)
	}

	wg.Wait()

	return nil
}

func openImage(filename string) (image.Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(file)
	return img, err
}

func parseMethod(methods string) ([]*halftone.Filter, error) {
	var filters []*halftone.Filter
	for _, name := range strings.Split(methods, ",") {
		switch strings.TrimSpace(name) {
		case "bayer":
			filters = append(filters, halftone.Bayer)
		case "bayer2":
			filters = append(filters, halftone.Bayer2)
		case "bayer4":
			filters = append(filters, halftone.Bayer4)
		}
	}

	if filters == nil {
		return nil, fmt.Errorf("invalid method name %s", methods)
	}

	return filters, nil
}

func writePng(filename string, img image.Image) {
	f, _ := os.Create(filename)
	png.Encode(f, img)
}
