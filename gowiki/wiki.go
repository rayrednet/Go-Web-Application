package main

import (
	"fmt"
	"os"
	"net/http"
	"html/template"
	"regexp"
)

var validPath = regexp.MustCompile(`^/(edit|save|view)/([a-zA-Z0-9]+)$`)

// Page structure
type Page struct {
    Title string
    Body  []byte
}

// Method to save the Page's content to a file
func (p *Page) save() error {
    filename := "data/" + p.Title + ".txt"
    return os.WriteFile(filename, p.Body, 0600)
}

// Method to load the Page's content from a file
func loadPage(title string) (*Page, error) {
    filename := "data/" + title + ".txt"
    body, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

// Handler to display the page content
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
    p, err := loadPage(title)
    if err != nil {
        http.Redirect(w, r, "/edit/"+title, http.StatusFound)
        return
    }
	// Convert the page content to HTML links
	p.Body = convertToLinks(p.Body)
    renderTemplate(w, "view", p)

}

// Handler to edit the page content
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    renderTemplate(w, "edit", p)
}

// Handler to save the page content
func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    p.save()
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

// Function to create a handler for the given function
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        m := validPath.FindStringSubmatch(r.URL.Path)
        if m == nil {
            http.NotFound(w, r)
            return
        }
        fn(w, r, m[2]) // Pass the title as the third argument to fn
    }
}

// Handler to display a message for monkeys
func monkeysHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love monkeys!")
}

// Handler to redirect to the FrontPage
func rootHandler(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/view/FrontPage", http.StatusFound)
}

var templates = template.Must(template.ParseFiles("tmpl/edit.html", "tmpl/view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    err := templates.ExecuteTemplate(w, tmpl+".html", p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// Function to convert the page content to HTML links
func convertToLinks(text []byte) []byte {
    linkRegexp := regexp.MustCompile(`\[(\w+)\]`)
    return linkRegexp.ReplaceAllFunc(text, func(match []byte) []byte {
        pageName := match[1 : len(match)-1] // Remove brackets
        link := fmt.Sprintf(`<a href="/view/%s">%s</a>`, pageName, pageName)
        return []byte(link)
    })
}

func main() {
	http.HandleFunc("/", rootHandler)
    http.HandleFunc("/view/", makeHandler(viewHandler))
    http.HandleFunc("/edit/", makeHandler(editHandler))
    http.HandleFunc("/save/", makeHandler(saveHandler))
    http.HandleFunc("/monkeys", monkeysHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("tmpl"))))

    fmt.Println("Starting server on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
