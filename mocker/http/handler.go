package http

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/plin2k/api-mocker/v2/config"
)

const (
	ProtocolName = "http"

	ReturnTypeFileHTML       = "file-html"
	ReturnTypeFile           = "file"
	ReturnTypeFileAttachment = "file-attachment"
	ReturnTypeJSON           = "json"
	ReturnTypeXML            = "xml"
	ReturnTypeTOML           = "toml"
	ReturnTypeYAML           = "yaml"
)

type handler struct {
	api    *gin.Engine
	config *config.Mocker
}

func New(cfg *config.Mocker) *handler {
	gin.SetMode(gin.TestMode)
	return &handler{
		api: gin.Default(),

		config: cfg,
	}
}

func (h *handler) Run() error {
	log.Printf("Starting http server on port %d", h.config.Port)
	return h.api.Run(fmt.Sprintf("%s:%d", h.config.Host, h.config.Port))
}

func (h *handler) Construct(body []byte) (err error) {
	var src *Source
	err = xml.Unmarshal(body, &src)
	if err != nil {
		return err
	}

	h.addGroups(h.api.Group("/"), src.Api.Groups)

	return nil
}

func (h *handler) addRoutes(g *gin.RouterGroup, group *Group) {
	for _, r := range group.Routes {
		g.Handle(r.Method, r.Path, h.output(group, r))
	}
}

func (h *handler) addGroups(rg *gin.RouterGroup, groups []*Group) {
	for _, g := range groups {
		group := rg.Group(g.Path)
		h.addRoutes(group, g)
		h.addGroups(group, g.Groups)
	}
}

func (h *handler) output(group *Group, route *Route) gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, h := range append(group.Headers, route.Headers...) {
			c.Header(h.Key, h.Value)
		}

		if route.Cookie.Name != "" {
			c.SetCookie(
				route.Cookie.Name,
				route.Cookie.Value,
				route.Cookie.MaxAge,
				route.Cookie.Path,
				route.Cookie.Domain,
				route.Cookie.Secure,
				route.Cookie.HttpOnly)
		}

		if route.StatusCode == 0 {
			route.StatusCode = http.StatusOK
		}

		time.Sleep(time.Second * route.Response.Delay)

		switch route.Response.Type {
		case ReturnTypeJSON:
			c.Data(route.StatusCode, "application/json", []byte(strings.Trim(route.Response.Value, "\n")))
		case ReturnTypeXML:
			c.Data(route.StatusCode, "application/xml", []byte(route.Response.Value))
		case ReturnTypeTOML:
			c.Data(route.StatusCode, "application/toml", []byte(route.Response.Value))
		case ReturnTypeYAML:
			c.Data(route.StatusCode, "application/yaml", []byte(route.Response.Value))
		case ReturnTypeFile:
			c.File(route.Response.Value)
		case ReturnTypeFileHTML:
			c.HTML(route.StatusCode, route.Response.Name, route.Response.Value)
		case ReturnTypeFileAttachment:
			parts := strings.Split(route.Response.Value, string(os.PathSeparator))
			c.Header("Content-Disposition", "attachment")
			c.Header("filename", parts[len(parts)-1])
			c.File(route.Response.Value)
		default:
			c.Data(route.StatusCode, "text/html\"", []byte(route.Response.Value))
		}
	}
}
