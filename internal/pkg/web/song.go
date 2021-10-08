package web

import (
	"encoding/json"
	"net/http"
	"path"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/lightster/finch/internal/pkg/model"
)

func ServeSong(config *Config) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {

		switch req.Method {
		case "GET":
			ServeViewSong(w, req, NewDependencies(config))
		case "POST":
			ServeCreateSong(w, req, NewDependencies(config))
		case "PUT":
			ServeUpdateSong(w, req, NewDependencies(config))
		case "DELETE":
			//ServeDeleteSong(w, req, NewDependencies(config))
		default:
			WriteMethodNotAllowedError(w)
		}
	}
}

func ServeViewSong(w http.ResponseWriter, req *http.Request, d *Dependencies) {
	cleanedPath := path.Clean(req.URL.Path)
	isMatch, err := path.Match("/song/*", cleanedPath)
	if err != nil {
		WriteServerError(w, err)
		return
	}
	if !isMatch {
		WriteBadRequestError(w)
		return
	}

	songID, err := uuid.Parse(path.Base(cleanedPath))
	if err != nil {
		WriteBadRequestError(w)
		return
	}

	db, err := d.InitDB(req.Context())
	if err != nil {
		WriteServerError(w, err)
	}
	defer func() {
		err := db.Close(req.Context())
		if err != nil {
			LogError(err)
		}
	}()

	song, err := model.FindSongById(req.Context(), db, songID)
	if err == pgx.ErrNoRows {
		WriteNotFoundError(w)
		return
	} else if err != nil {
		WriteServerError(w, err)
		return
	}

	jsonSong, err := json.MarshalIndent(song, "", "  ")
	if err != nil {
		WriteServerError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonSong)
	if err != nil {
		WriteServerError(w, err)
		return
	}
}

func ServeCreateSong(w http.ResponseWriter, req *http.Request, d *Dependencies) {
	song := &model.Song{}
	err := json.NewDecoder(req.Body).Decode(song)
	if err != nil {
		LogError(err)
		WriteBadRequestError(w)
		return
	}

	ok, message := song.Validate()
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(message))
		if err != nil {
			LogError(err)
		}
		return
	}

	db, err := d.InitDB(req.Context())
	if err != nil {
		WriteServerError(w, err)
	}
	defer func() {
		err := db.Close(req.Context())
		if err != nil {
			LogError(err)
		}
	}()

	err = song.Create(req.Context(), db)
	if err != nil {
		WriteServerError(w, err)
		return
	}

	jsonSong, err := json.MarshalIndent(song, "", "  ")
	if err != nil {
		WriteServerError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonSong)
	if err != nil {
		WriteServerError(w, err)
		return
	}
}

func ServeUpdateSong(w http.ResponseWriter, req *http.Request, d *Dependencies) {
	cleanedPath := path.Clean(req.URL.Path)
	isMatch, err := path.Match("/song/*", cleanedPath)
	if err != nil {
		WriteServerError(w, err)
		return
	}
	if !isMatch {
		WriteBadRequestError(w)
		return
	}

	songID, err := uuid.Parse(path.Base(cleanedPath))
	if err != nil {
		WriteBadRequestError(w)
		return
	}

	db, err := d.InitDB(req.Context())
	if err != nil {
		WriteServerError(w, err)
	}
	defer func() {
		err := db.Close(req.Context())
		if err != nil {
			LogError(err)
		}
	}()

	song, err := model.FindSongById(req.Context(), db, songID)
	if err == pgx.ErrNoRows {
		WriteNotFoundError(w)
		return
	} else if err != nil {
		WriteServerError(w, err)
		return
	}

	err = json.NewDecoder(req.Body).Decode(song)
	if err != nil {
		LogError(err)
		WriteBadRequestError(w)
		return
	}
	song.SongID = songID

	ok, message := song.Validate()
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(message))
		if err != nil {
			LogError(err)
		}
		return
	}

	err = song.Update(req.Context(), db)
	if err != nil {
		WriteServerError(w, err)
		return
	}

	jsonSong, err := json.MarshalIndent(song, "", "  ")
	if err != nil {
		WriteServerError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonSong)
	if err != nil {
		WriteServerError(w, err)
		return
	}
}
