package groupino

import (
	"bytes"
	"html/template"
	"net/http"
)

// RendError renders a custom error page using the given message and HTTP status code.
// It loads the "pageerror.html" template and displays the appropriate error message to the user.
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

	var buffer bytes.Buffer

	err = temp.Execute(&buffer, Errorpage)
	if err != nil {
		RendError(w, "500 internal server error ", http.StatusInternalServerError)
		return
	}
	_, err = buffer.WriteTo(w)
	if err != nil {
		RendError(w, "500 internal server error ", http.StatusInternalServerError)
		return
	}
}
