package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/plin2k/api-mocker/domain"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	ReturnTypeFileHTML       = "file-html"
	ReturnTypeFile           = "file"
	ReturnTypeFileAttachment = "file-attachment"
	ReturnTypeJSON           = "json"
	ReturnTypeXML            = "xml"
	ReturnTypeTOML           = "toml"
	ReturnTypeYAML           = "yaml"
)

type server struct {
	handler *gin.Engine
}

func NewServer() *server {
	gin.SetMode(gin.TestMode)
	return &server{
		handler: gin.Default(),
	}
}

func (s *server) Run(port int) error {
	return s.handler.Run(fmt.Sprintf(":%d", port))
}

func (s *server) Construct(src *domain.Source) {
	s.addGroups(s.handler.Group("/"), src.Api.Groups)
	s.addRoutes(s.handler.Group("/"), src.Api.Routes)
}

func (s *server) addRoutes(g *gin.RouterGroup, routes []domain.Route) {
	for _, r := range routes {
		switch r.Method {
		case http.MethodGet:
			g.GET(r.Pattern, s.output(r))
		case http.MethodPost:
			g.POST(r.Pattern, s.output(r))
		case http.MethodDelete:
			g.DELETE(r.Pattern, s.output(r))
		case http.MethodHead:
			g.HEAD(r.Pattern, s.output(r))
		case http.MethodOptions:
			g.OPTIONS(r.Pattern, s.output(r))
		case http.MethodPatch:
			g.PATCH(r.Pattern, s.output(r))
		case http.MethodPut:
			g.PUT(r.Pattern, s.output(r))
		default:
			g.Any(r.Pattern, s.output(r))
		}
	}
}

func (s *server) addGroups(rg *gin.RouterGroup, groups []domain.Group) {
	for _, g := range groups {
		group := rg.Group(g.Pattern)
		s.addRoutes(group, g.Routes)
		s.addGroups(group, g.Groups)
	}
}

func (s *server) output(route domain.Route) gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, h := range route.Headers {
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

		time.Sleep(time.Second * route.Delay.Value)

		switch route.Return.Type {
		case ReturnTypeJSON:
			c.Header("Content-Type", "application/json")
			c.IndentedJSON(route.StatusCode, route.Return.Value)
		case ReturnTypeXML:
			c.Header("Content-Type", "application/xml")
			c.XML(route.StatusCode, route.Return.Value)
		case ReturnTypeTOML:
			c.Header("Content-Type", "application/toml")
			c.TOML(route.StatusCode, route.Return.Value)
		case ReturnTypeYAML:
			c.Header("Content-Type", "application/yaml")
			c.YAML(route.StatusCode, route.Return.Value)
		case ReturnTypeFile:
			c.File(route.Return.Value)
		case ReturnTypeFileHTML:
			c.HTML(route.StatusCode, route.Return.Name, route.Return.Value)
		case ReturnTypeFileAttachment:
			parts := strings.Split(route.Return.Value, string(os.PathSeparator))
			c.Header("Content-Disposition", "attachment")
			c.Header("filename", parts[len(parts)-1])
			c.File(route.Return.Value)
		default:
			c.Header("Content-Type", "text/html")
			c.String(http.StatusOK, route.Return.Value)
		}
	}
}
