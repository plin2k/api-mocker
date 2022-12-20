package domain

import (
	"encoding/xml"
	"time"
)

type Source struct {
	XMLName xml.Name `xml:"xml"`

	Protocol string `xml:"protocol,attr"`
	Version  string `xml:"version"`

	Name        string `xml:"name"`
	Description string `xml:"description"`

	Api struct {
		Groups []Group `xml:"group"`
		Routes []Route `xml:"route"`
	} `xml:"api"`
}

type Group struct {
	Pattern     string  `xml:"pattern,attr"`
	Description string  `xml:"description,attr"`
	Headers     Headers `xml:"header"`
	Routes      []Route `xml:"route"`
	Groups      []Group `xml:"group"`
}

type Route struct {
	Pattern     string  `xml:"pattern,attr"`
	Description string  `xml:"description,attr"`
	StatusCode  int     `xml:"status-code,attr"`
	Method      string  `xml:"method,attr"`
	Headers     Headers `xml:"header"`
	Delay       Delay   `xml:"delay"`
	Return      Return  `xml:"return"`
	Cookie      Cookie  `xml:"cookie"`
}

type Cookie struct {
	Name        string `xml:"name,attr"`
	Value       string `xml:"value,attr"`
	MaxAge      int    `xml:"max-age,attr"`
	Path        string `xml:"path,attr"`
	Domain      string `xml:"domain,attr"`
	Secure      bool   `xml:"secure,attr"`
	HttpOnly    bool   `xml:"http-only,attr"`
	Description string `xml:"description,attr"`
}

type Delay struct {
	Value       time.Duration `xml:"value,attr"`
	Description string        `xml:"description,attr"`
}

type Return struct {
	Value       string `xml:",chardata"`
	Type        string `xml:"type,attr"`
	Description string `xml:"description,attr"`
	Name        string `xml:"name,attr"`
}

type Headers []Header

type Header struct {
	Key         string `xml:"key,attr"`
	Value       string `xml:"value,attr"`
	Description string `xml:"description,attr"`
}
