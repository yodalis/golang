package main

import (
	"os"
	"text/template"
)

// Templates: Os templates podem ser utilizados como facilitadores em processos
// Exemplo: enviar um email periódicamente

type Curso struct {
	Name string
	Time int
}

func main() {
	curso := Curso{"Go", 40}

	tmp := template.New("CursoTemplate")
	tmp, _ = tmp.Parse("Curso: {{.Name}} - Time {{.Time}}")
	err := tmp.Execute(os.Stdout, curso)

	if err != nil {
		panic(err)
	}
}
