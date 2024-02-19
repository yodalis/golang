package main

import (
	"os"
	"text/template"
)

// Template.Must te auxilia a fazer tudo que fez no New e Parse em um unico comando
type Curso struct {
	Name string
	Time int
}
type Cursos []Curso

func main() {
	t := template.Must(template.New("template.html").ParseFiles("template.html"))

	err := t.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Java", 10},
		{"Typescript", 20},
	})

	if err != nil {
		panic(err)
	}
}
