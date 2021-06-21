package internal

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
)

type HttpServer struct {
	port int
	tmplFile string
	members *Members
}

func NewHttpServer(port int) *HttpServer  {
	members := NewMembers()
	srv := HttpServer{
		port: port,
		tmplFile: "./web/index.gohtml",
		members: &members,
	}
	fs := http.FileServer(StaticFileSystem{"./web/static/"})
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", srv.handle)
	return &srv
}

func (srv *HttpServer) Start()  {
	log.Printf("Start server on port :%d...\n", srv.port)
	port := fmt.Sprintf(":%d", srv.port)
	if err := http.ListenAndServe(port, nil); err != http.ErrServerClosed {
		log.Fatalf("Error start http server: %s", err.Error())
	}
	log.Printf("Server shutdown")
}

func (srv *HttpServer) handle(w http.ResponseWriter, r *http.Request)  {
	tmplData := NewTemplateData()
	tmplData.Members = srv.members
	if r.Method == http.MethodPost {
		m := Member{
			Name:  strings.TrimSpace(r.PostFormValue("name")),
			Email: strings.TrimSpace(r.PostFormValue("email")),
		}
		err := srv.members.AddMember(&m)
		if err != nil {
			tmplData.Fields.Name = m.Name
			tmplData.Fields.Email = m.Email
			tmplData.Errors = err
		} else {
			log.Printf("Add new member: %s\n", m)
		}
	}
	srv.render(w, tmplData)
}

func (srv *HttpServer) render(w http.ResponseWriter, data interface{}) {
	tmpl, err := template.ParseFiles(srv.tmplFile)
	if err != nil {
		log.Printf("Error parse template: %s\n", err.Error())
		srv.sendServerError(w)
	}
	if err := tmpl.Execute(w, data); err != nil {
		log.Printf("Error render template: %s\n", err.Error())
		srv.sendServerError(w)
	}
}

func (srv *HttpServer) sendServerError(w http.ResponseWriter) {
	http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
}