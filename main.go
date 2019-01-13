package main
import (
	"fmt"
	"time"
	"net/http"
	"html/template"
)

type WelcomeInfo struct {
	Name string
	Time string
}

func main() {
	welcomeInfo := WelcomeInfo{"Anonymous", time.Now().Format(time.Stamp)}

	htmlTemplate := template.Must(template.ParseFiles("template/welcome-template.html"))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(resWriter http.ResponseWriter, req *http.Request) {
		if name := req.FormValue("name"); name != "" {
			welcomeInfo.Name = name
		}

		if err := htmlTemplate.ExecuteTemplate(resWriter, "welcome-template.html", welcomeInfo); err != nil {
			http.Error(resWriter, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Listening...")
	fmt.Println(http.ListenAndServe(":8080", nil))
}