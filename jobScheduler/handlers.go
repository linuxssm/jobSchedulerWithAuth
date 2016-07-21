package jobScheduler

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	entries := Store.AllEntries()
	tmpl := template.New("base")
	template.Must(tmpl.Parse(BaseTmplStr))
	template.Must(tmpl.Parse(ViewTmplStr))
	tmpl.Execute(w, entries)
}

func CssHandler(w http.ResponseWriter, r *http.Request) {
	name := path.Base(r.URL.Path)
	asset, ok := AssetMap[name]

	if ok {
		w.Header().Set("Content-Type", "text/css")
		io.WriteString(w, asset)
	} else {
		w.Header().Set("Content-Type", "text/plain")
		io.WriteString(w, "asset not defined")
	}
}

func JsHandler(w http.ResponseWriter, r *http.Request) {
	name := path.Base(r.URL.Path)
	asset, ok := AssetMap[name]
	if ok {
		w.Header().Set("Content-Type", "application/javascript")
		io.WriteString(w, asset)
	} else {
		w.Header().Set("Content-Type", "text/plain")
		io.WriteString(w, "asset not defined")
	}
}

func NewHandler(w http.ResponseWriter, r *http.Request) {
	entry := NewEntry("", "", "")
	tmpl := template.New("base")
	template.Must(tmpl.Parse(BaseTmplStr))
	template.Must(tmpl.Parse(AddTmplStr))
	tmpl.Execute(w, entry)

}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	entry := Store.GetEntry(id)
	tmpl := template.New("base")
	template.Must(tmpl.Parse(BaseTmplStr))
	template.Must(tmpl.Parse(EditTmplStr))
	tmpl.Execute(w, entry)
}

func SaveHandler(w http.ResponseWriter, r *http.Request) {
	err1 := r.ParseMultipartForm(32 << 20)
	if err1 != nil {
		return
	}
	file, handler, err := r.FormFile("files")
	if err == nil {
		defer file.Close()

		f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		io.Copy(f, file)
	}
	entry := NewEntryFromReq(r)
	entry.Endpoint = strings.TrimSpace(entry.Endpoint)
	replaced := strings.NewReplacer(" ", "%20")
	entry.Endpoint = replaced.Replace(entry.Endpoint)
	Store.SaveEntry(entry)
	Skeddy.ReStart(Store.AllEntries())
	http.Redirect(w, r, "/", 301)
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	err1 := r.ParseMultipartForm(32 << 20)
	if err1 != nil {
		return
	}
	file, handler, err := r.FormFile("files")
	if err == nil {
		defer file.Close()

		f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		io.Copy(f, file)
	}
	//for i:= 0;i<10000;i++{
	//	entry := NewEntryFromReq(r)
	//	entry.Endpoint = strings.TrimSpace(entry.Endpoint)
	//	replaced := strings.NewReplacer(" ", "%20")
	//	entry.Endpoint = replaced.Replace(entry.Endpoint)
	//
	//	entry.ID= entry.ID+string(i)
	//	//newPayload := entry.Payload+string(i)
	//	entry.Payload += strconv.Itoa(i)
	//	Store.SaveEntry(entry)
	//	Skeddy.AddEntry(entry)
	//
	//}
	entry := NewEntryFromReq(r)
	entry.Endpoint = strings.TrimSpace(entry.Endpoint)
	replaced := strings.NewReplacer(" ", "%20")
	entry.Endpoint = replaced.Replace(entry.Endpoint)
	Store.SaveEntry(entry)
	Skeddy.AddEntry(entry)

	//Skeddy.AddEntry(entry)
	http.Redirect(w, r, "/", 301)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	Store.DeleteEntry(id)
	Skeddy.ReStart(Store.AllEntries())
	http.Redirect(w, r, "/", 301)
}

func ValidateExpression(w http.ResponseWriter, r *http.Request) {
	expression := path.Base(r.URL.Path)
	err := ParseExpression(expression)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}
