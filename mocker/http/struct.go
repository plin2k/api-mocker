package http

import (
	"time"
)

type Source struct {
	Version string `xml:"version"`

	Name        string `xml:"name"`
	Description string `xml:"description"`

	Api struct {
		Groups []*Group `xml:"group"`
	} `xml:"api"`
}

type Group struct {
	Essential
	Path    string    `xml:"path,attr"`
	Headers []*Header `xml:"header"`
	Routes  []*Route  `xml:"route"`
	Groups  []*Group  `xml:"group"`
}

type Route struct {
	Essential
	Path       string    `xml:"path,attr"`
	StatusCode int       `xml:"status-code,attr"`
	Method     string    `xml:"method,attr"`
	Headers    []*Header `xml:"header"`
	Response   Response  `xml:"response"`
	Cookie     Cookie    `xml:"cookie"`
}

type Cookie struct {
	Essential
	Name     string `xml:"name,attr"`
	Value    string `xml:"value,attr"`
	MaxAge   int    `xml:"max-age,attr"`
	Path     string `xml:"path,attr"`
	Domain   string `xml:"domain,attr"`
	Secure   bool   `xml:"secure,attr"`
	HttpOnly bool   `xml:"http-only,attr"`
}

type Response struct {
	Essential
	Value string        `xml:",chardata"`
	Type  string        `xml:"type,attr"`
	Name  string        `xml:"name,attr"`
	Delay time.Duration `xml:"delay,attr"`
}

type Header struct {
	Essential
	Key   string `xml:"key,attr"`
	Value string `xml:"value,attr"`
}

type Essential struct {
	Description string `xml:"description,attr"`
}
