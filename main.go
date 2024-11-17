package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
    Link string
}

type Connection struct {
    Service string
    Link string
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
                    {
                        Title: "Integrated K8S Controller for Ingress and Services",
                        Description: `Developed a Kubernetes Controller using KubeBuilder and Client-Go 
                                      to automate and streamline the lifecycle management of services and ingress resources, 
                                      ensuring efficient and seamless deployments`,
                        Link: "https://github.com/thakurnishu/sik-controller-kubebuilder",
                    },
                    {
                        Title: "Orchestrating Kubernetes Deployments using CICD, GitOps and IAC",
                        Description: `Orchestrated efficient Kubernetes deployments using CI/CD, GitOps, and IaC.
                                      Designed a robust Jenkins pipeline with a shared library approach integrated 
                                      ArgoCD for seamless Kubernetes deployments, and automated infrastructure 
                                      management with Terraform, ensuring consistency and scalability.`,
                        Link: "https://github.com/thakurnishu/GoLang-WebServer",
                    },
                    {
                        Title: "Deployed Static Website with Edge Caching and CI/CD on Azure DevOps",
                        Description: `Deployed a static website with edge caching and CI/CD on Microsoft Azure, leveraging
                                      Azure Storage and Azure CDN for cost-effective hosting and optimized global performance.
                                      Automated deployments using Azure DevOps CI/CD pipelines, ensuring rapid updates and scalability.
                                      Configured Azure CDN to enhance user experience with faster website loading times worldwide.`,
                        Link: "https://github.com/thakurnishu/StaticWebsite-AzureCloud",

                    },
                    {
                        Title: "Deployment of a Scalable and Highly Available 3-Tier Web Application on Azure",
                        Description: `Implemented a scalable 3-tier web application on Azure using Golang, Kubernetes, Docker, and Terraform.
                                      Leveraged VMSS, Load Balancer, and Azure Application Gateway for high availability and dynamic scaling.
                                      Automated deployments with Azure CLI and secured infrastructure with Trivy and SonarQube,
                                      integrating Azure Database for MySQL for robust data management.`,
                        Link: "https://github.com/thakurnishu/Tier3WebApplication-AzureCloud",

                    },
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

    page := newPage()

    e.GET("/", pageRender("index-page", nil))
    e.GET("/home", pageRender("home", nil))
    e.GET("/about", pageRender("about", page.Data.About))
    e.GET("/project", pageRender("project", page.Data.Projects))
    e.GET("/connect", pageRender("connect", page.Data.Connects))

    e.GET("/healthz", pageRender("healthz", nil))

    e.Logger.Fatal(e.Start(port))
}
