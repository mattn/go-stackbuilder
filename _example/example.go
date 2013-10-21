package main

import (
	"github.com/justinas/nosurf"
	"github.com/keep94/weblogs"
	"github.com/mattn/go-stackbuilder"
	"html/template"
	"net/http"
)

var templateString string = `
<!doctype html>
<html>
<body>
{{ if .name }}
<p>Your name: {{ .name }}</p>
{{ end }}
<form action="/" method="POST">
<input type="text" name="name">

<!-- Try removing this or changing its value
     and see what happens -->
<input type="hidden" name="csrf_token" value="{{ .token }}">
<input type="submit" value="Send">
</form>
</body>
</html>
`
var templ = template.Must(template.New("t1").Parse(templateString))

func main() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		context := make(map[string]string)
		context["token"] = nosurf.Token(r)
		if r.Method == "POST" {
			context["name"] = r.FormValue("name")
		}
		templ.Execute(w, context)
	}))

	http.ListenAndServe(":8888", stackbuilder.NewStackBuilder(
		nosurf.New,
		weblogs.Handler,
	))
}
