package mocker

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/plin2k/api-mocker/config"
	"github.com/plin2k/api-mocker/faker"
	"github.com/plin2k/api-mocker/mocker/http"
	"github.com/plin2k/api-mocker/mocker/websocket"
)

type Handler interface {
	Run() error
	Construct([]byte) error
}

type protocolXml struct {
	Body     string `xml:",chardata"`
	Protocol string `xml:"protocol,attr"`
	Version  string `xml:"version"`
}

type protocol struct {
	handler Handler
	config  *config.Mocker
}

func New(cfg *config.Mocker) *protocol {
	return &protocol{
		config: cfg,
	}
}

func (p *protocol) Process() error {
	data, err := os.ReadFile(p.config.SrcPath)
	if err != nil {
		return err
	}

	data, err = faker.Process(data)
	if err != nil {
		return err
	}

	var src *protocolXml
	err = xml.Unmarshal([]byte(data), &src)
	if err != nil {
		return err
	}

	switch src.Protocol {
	case http.ProtocolName:
		p.handler = http.New(p.config)
	case websocket.ProtocolName:
		p.handler = websocket.New(p.config)
	default:
		return fmt.Errorf("unknown mocker")
	}

	return p.handler.Construct(data)
}

func (p *protocol) Run() error {
	return p.handler.Run()
}
