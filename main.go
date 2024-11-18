package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/yaml.v3"
)

var port = ":8000"

type Templates struct {
    templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
    return &Templates{
        templates: template.Must(template.ParseGlob("template/*.html")),
    }
}

func pageRender(block string, data interface{}) func(c echo.Context) error {
    return func(c echo.Context) error {
        return c.Render(http.StatusOK, block, data)
    }
}

func main() {
	// Open the YAML file
	file, err := os.Open("data.yaml")
	if err != nil {
		log.Fatalf("Error opening YAML file: %v", err)
	}
	defer file.Close()

    e := echo.New()
    e.Use(middleware.Logger())

    e.Renderer = newTemplate()

    page := Page{}
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&page)
	if err != nil {
		log.Fatalf("Error decoding YAML file: %v", err)
	}

    e.GET("/", pageRender("index-page", page.Data.Home))
    e.GET("/home", pageRender("home", page.Data.Home))
    e.GET("/about", pageRender("about", page.Data.About))
    e.GET("/project", pageRender("project", page.Data.Projects))
    e.GET("/connect", pageRender("connect", page.Data.Connects))

    e.GET("/healthz", pageRender("healthz", nil))

    e.Logger.Fatal(e.Start(port))
}
