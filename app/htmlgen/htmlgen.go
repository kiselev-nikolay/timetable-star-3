package htmlgen

import (
	"bytes"
	"embed"
	_ "embed"
	"html/template"
	"time"
)

//go:embed templates/*
var templatesFS embed.FS

var (
	tpl *template.Template
)

func init() {
	tpl = template.Must(template.ParseFS(templatesFS, "templates/*"))
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

type Notification struct {
	Text string
}

type State struct {
	Time    string
	Welcome *Notification
}

func (s *State) Update() {
	s.Time = time.Now().Format("15:04:05")
}

func (s *State) Render() string {
	s.Update()
	page := bytes.NewBuffer([]byte{})
	executeErr := tpl.ExecuteTemplate(page, "index.html", s)
	handleError(executeErr)
	return page.String()
}

func New() *State {
	return &State{
		Welcome: &Notification{
			Text: "Hello",
		},
	}
}
