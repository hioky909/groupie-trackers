package handler

import (
	"fmt"
	structure "groupie-tracker/models"
	"groupie-tracker/utils"
	"html/template"
	"net/http"
	"strconv"
)

var artists []structure.Artists

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	Artists, err := utils.FetchArtists()
	if err != nil {
		fmt.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError, "Internal server error")
		return
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Println(err)
		ErrorHandler(w, r, http.StatusNotFound, "index.html not found")
		return
	}
	artists = Artists
	err = tmpl.Execute(w, artists)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Internal server error")
		return
	}

}

func DetailsHandler(w http.ResponseWriter, r *http.Request) {
	Id := r.URL.Path[len("/artist/"):]
	id, err := strconv.Atoi(Id)
	if err != nil || id < 0 {
		ErrorHandler(w, r, http.StatusNotFound, "Page not found")
		return
	}

	artist := artists[id-1]

	relation, err := utils.FetchRelation(Id)
	if err != nil {
		fmt.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError, "Internal sevrer error")
		return
	}

	tmpl, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		fmt.Println(err)
		ErrorHandler(w, r, http.StatusNotFound, "Artist.html not found")
		return
	}

	data := struct {
		Artist   structure.Artists
		Relation structure.Relation
	}{
		Artist:   artist,
		Relation: relation,
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Internal server error")
		return
	}
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, code int, message string) {
	t, err := template.ParseFiles("templates/error.html")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	data := struct {
		ErrorCode    int
		ErrorMessage string
	}{
		ErrorCode:    code,
		ErrorMessage: message,
	}
	err = t.Execute(w, data)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Internal server error")
		return
	}
}

func LocationHandler(w http.ResponseWriter, r *http.Request) {
	Id := r.URL.Path[len("/artist/"):]
	Id = Id[:len(Id)-len("/locations")]
	id, err := strconv.Atoi(Id)
	if err != nil || id < 0 {
		ErrorHandler(w, r, http.StatusNotFound, "Page not found")
		return
	}

	artist := artists[id-1]

	relation, err := utils.FetchRelation(Id)
	if err != nil {
		fmt.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError, "Internal server error")
		return
	}

	tmpl, err := template.ParseFiles("templates/locations.html")
	if err != nil {
		fmt.Println(err)
		ErrorHandler(w, r, http.StatusNotFound, "locations.html not found")
		return
	}

	data := struct {
		Artist   structure.Artists
		Relation structure.Relation
	}{
		Artist:   artist,
		Relation: relation,
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "InternalServerError")
		return
	}
}

