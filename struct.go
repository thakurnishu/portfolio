package main

import "html/template"

type Page struct {
    Data Data `yaml:"data"`
}

type Data struct {
    Home Home
    Projects Projects `yaml:"projects"`
    Connects Connects `yaml:"connects"`
    About About `yaml:"about"`
}

type Home struct {
    Name string `yaml:"name"`
    Description template.HTML `yaml:"description"`
    Email Email `yaml:"email"`
}
type Email struct {
    Addr string `yaml:"addr"`
    Subject string `yaml:"subject"`
}

type Projects struct {
    ProjectList []Project `yaml:"project_list"`
}

type Connects struct {
    ConnectionList []Connection `yaml:"connection_list"`
}

type About struct {
    Description template.HTML `yaml:"description"`
}
type Project struct {
    Title string `yaml:"title"`
    Description string `yaml:"description"`
    Link string `yaml:"link"`
}

type Connection struct {
    Service string `yaml:"service"`
    Link string `yaml:"link"`
}
