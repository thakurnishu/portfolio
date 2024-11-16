package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var port = ":8989"

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

type Page struct {
    Data Data
}

type Data struct {
    Projects Projects
    Connects Connects
    About About
}

type Projects struct {
    ProjectList []Project
}

type Connects struct {
    ConnectionList []Connection
}

type About struct {
    Description template.HTML
}
type Project struct {
    Title string
    Description string
}

type Connection struct {
    Service string
    Link string
}

func newProject(title, description string) Project {
    return Project{
        Title: title,
        Description: description,
    }
}

func newConnection(service, link string) Connection {
    return Connection{
        Service: service,
        Link: link,
    }
}

func newPage() Page {
    return Page{
        Data: Data{
            Projects: Projects{
                ProjectList: []Project{
                    newProject("Project1", "Descriptions 1"),
                    newProject("Project2", "Descriptions 2"),
                    newProject("Project3", "Descriptions 3 sdfsd sdf fsd dfgd fgdfb fgh fgh fgh fgh fgh fgh fgh 1"),
                    newProject("Project4", "Descriptions 4 sdf ssdf sdf sdf"),
                },
            },
            Connects: Connects{
                ConnectionList: []Connection{
                    newConnection("github", "https://github.com/thakurnishu"),
                    newConnection("linkedin", "https://www.linkedin.com/in/contact-nishant-singh"),
                    newConnection("twitter", "https://x.com/ntsh53"),
                    newConnection("medium", "https://medium.com/@thakurnishant5102003"),
                    newConnection("hashnode", "https://hashnode.com/@Nishu0510"),
                },
            },
            About: About{
                Description: template.HTML(`An undergraduate from Shri Mata Vaishno Devi University.
                            I have a strong passion for 
                            <b>Cloud</b> 
                            and 
                            <b>DevOps</b> technologies and I am continuously learning new technologies. 
                            I have valuable experience in DevOps, implementing Continuous Integration and Deployment (CI/CD) using Github Action.
                            Currently, I am expanding my knowledge on different technologies`),
            },
        },
    }
}

func main() {
    e := echo.New()
    e.Use(middleware.Logger())

    e.Renderer = newTemplate()

    e.Static("/static", "static")

    page := newPage()

    e.GET("/", pageRender("index-page", nil))
    e.GET("/home", pageRender("home", nil))
    e.GET("/about", pageRender("about", page.Data.About))
    e.GET("/project", pageRender("project", page.Data.Projects))
    e.GET("/connect", pageRender("connect", page.Data.Connects))

    e.GET("/healthz", pageRender("healthz", nil))

    e.Logger.Fatal(e.Start(port))
}
