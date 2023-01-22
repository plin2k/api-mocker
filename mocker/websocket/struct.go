package websocket

import "time"

const (
	TypeText   = "text"
	TypeBinary = "binary"
)

type Source struct {
	Version string `xml:"version"`

	Name        string `xml:"name"`
	Description string `xml:"description"`

	OnOpen    OnOpen    `xml:"onopen"`
	OnClose   OnClose   `xml:"onclose"`
	OnMessage OnMessage `xml:"onmessage"`
	OnError   OnError   `xml:"onerror"`

	Ping Ping `xml:"ping"`
	Pong Pong `xml:"pong"`

	While struct {
		Messages []*Message `xml:"message"`
	} `xml:"while"`
}

type Message struct {
	Essential
	Delay    int       `xml:"delay,attr"`
	Count    uint      `xml:"count,attr"`
	Response *Response `xml:"response"`
}

type Essential struct {
	Description string `xml:"description,attr"`
}

type OnOpen struct {
	Essential
	Response *Response `xml:"response"`
}

type OnClose struct {
	Essential
	Response *Response `xml:"response"`
}

type OnMessage struct {
	Essential
	Response *Response `xml:"response"`
}

type OnError struct {
	Essential
	Response *Response `xml:"response"`
}

type Ping struct {
	Essential
	Response *Response `xml:"response"`
}

type Pong struct {
	Essential
	Response *Response `xml:"response"`
}

type Response struct {
	Essential
	Value string        `xml:",chardata"`
	Type  string        `xml:"type,attr"`
	Delay time.Duration `xml:"delay,attr"`
}

func (r *Response) GetType() int {
	if r.Type == TypeBinary {
		return 2
	}

	return 1
}
