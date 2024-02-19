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

func main() {
	curso := Curso{"Go", 40}
	t := template.Must(template.New("CursoTemplate").Parse("Curso: {{.Name}} - Time {{.Time}}"))

	err := t.Execute(os.Stdout, curso)

	if err != nil {
		panic(err)
	}
}
