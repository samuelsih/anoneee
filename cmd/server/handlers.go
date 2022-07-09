package server

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (app *App) getAll(w http.ResponseWriter, r *http.Request) {
	app.Data.ToJSON(w)
}

func (app *App) findByID(w http.ResponseWriter, r *http.Request) {

	for _, sliceValue := range app.Data.SliceValue {
		for key, value := range sliceValue {
			if key == "id" {
				id, err := strconv.Atoi( chi.URLParam(r, "id") )
				if err != nil {
					WriteErr(w, http.StatusBadRequest, "Cannot parse parameter")
					return
				}

				if id == value {
					ToJSON(w, sliceValue)
					return
				}
			}
		}
	}

	WriteErr(w, http.StatusNotFound, "Data not found")
}
