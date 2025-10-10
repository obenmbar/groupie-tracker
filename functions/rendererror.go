package groupino

import (
	"html/template"
	"net/http"
)

func RendError(w http.ResponseWriter, msgerror string, Code int) {
	temp, err := template.ParseFiles("templates/pageerror.html")
	if err != nil {
		http.Error(w, " 500 - error en parsefile d'error", 500)
		return
	}

	Errorpage := Errorpage{
		Mesege: msgerror,
		Status: Code,
	}
	w.WriteHeader(Code)
	err = temp.Execute(w, Errorpage)
	if err != nil {
		http.Error(w, " 500 - error en execute le file ", 500)
		return
	}
}
