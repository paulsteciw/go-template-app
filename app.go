package main

import (
  "html/template"
  "io/ioutil"
  "log"
  "net/http"
  "time"
)

func main() {
  mux := http.NewServeMux()

  server := &http.Server{
    Addr:    ":8080",
    Handler: mux,
  }

  funcs := template.FuncMap{
    "Date": func() string { return time.Now().Format("2006-01-02") },
  }
  layout := initTemplate(template.New("layout"), funcs)
  initTemplate(layout.New("post"), funcs)

  mux.HandleFunc("/", home(layout))

  log.Fatal(server.ListenAndServe())
}

type Post struct {
  Title string
  Body  string
}

func home(layout *template.Template) func(http.ResponseWriter, *http.Request) {
  posts := []Post{
    Post{Title: "Hello World!", Body: "I need to start coming up with something better to say than hello world."},
    Post{Title: "Foo bar", Body: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent egestas."},
  }
  return func(w http.ResponseWriter, r *http.Request) {
    err := layout.Execute(w, posts)
    if err != nil {
      log.Panic(err)
    }
  }
}

func initTemplate(template *template.Template, funcMap template.FuncMap) *template.Template {
  template.Funcs(funcMap)
  contents, err := ioutil.ReadFile(template.Name() + ".html")
  if err != nil {
    log.Panic(err)
  }
  template.Parse(string(contents))
  return template
}