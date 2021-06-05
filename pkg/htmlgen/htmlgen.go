package htmlgen

import (
	"bytes"
	"embed"
	_ "embed"
	"html/template"
	"strings"
)

//go:embed templates/*
var templatesFS embed.FS

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

type State struct {
	tpl       *template.Template
	callbacks map[string]func(string)
	Data      map[string]interface{}
}

func (s *State) AddCallback(key string, callback func(action string)) {
	s.callbacks[key] = callback
}

func (s *State) Update(key string) {
	command := strings.Split(key, " ")
	s.callbacks[command[1]](command[0])
}

func (s *State) Render() string {
	page := bytes.NewBuffer([]byte{})
	executeErr := s.tpl.ExecuteTemplate(page, "<root>", s.Data)
	handleError(executeErr)
	return page.String()
}

func New(rootTemplate string) *State {
	tpl := template.Must(template.ParseFS(templatesFS, "templates/*"))
	tpl.New("<root>").Parse(rootTemplate)
	return &State{
		tpl:       tpl,
		callbacks: make(map[string]func(string)),
		Data:      make(map[string]interface{}),
	}
}
