package main

import (
	"encoding/xml"
	"flag"
	"github.com/plin2k/api-mocker/domain"
	"github.com/plin2k/api-mocker/http"
	"io/ioutil"
	"log"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "Port (default: 8080)")
	var srcFile string
	flag.StringVar(&srcFile, "src", "source.xml", "Source file (default: source.xml)")
	flag.Parse()

	src, err := getSource(srcFile)
	if err != nil {
		log.Fatal(err)
	}

	api := http.NewServer()
	api.Construct(src)

	log.Printf("Starting server on port %d", port)
	if err = api.Run(port); err != nil {
		log.Fatal(err)
	}

}

func getSource(path string) (*domain.Source, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var src *domain.Source
	err = xml.Unmarshal([]byte(data), &src)
	if err != nil {
		return nil, err
	}

	return src, nil
}
