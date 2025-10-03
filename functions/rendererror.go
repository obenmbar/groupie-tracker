package groupino

import (
	"html/template"
	"net/http"
)

func RendError(w http.ResponseWriter, msgerror string, code int) {
	Check("templates/pagerror.html")

	temp, err := template.ParseFiles("templates/pageerror.html")
	if err != nil {
		http.Error(w, " 500 - error en parsefile d'error", 500)
		return
	}
	

	Errorpage := Errorpage{
		Mesege: msgerror,
		Status: code,
	}


	err = temp.Execute(w,Errorpage)
	w.WriteHeader(code)
	if err != nil {
			http.Error(w, " 500 - error en execute le file ", 500)
			return
	}
		
}
