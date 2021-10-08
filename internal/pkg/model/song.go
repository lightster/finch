package model

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

type Song struct {
	SongID   uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Genre    string    `json:"genre"`
	Artist   string    `json:"artist"`
	Length   int       `json:"length"`
	FilePath string    `json:"file_path"`
	Ranking  int       `json:"ranking"`
}

func FindSongById(ctx context.Context, db *pgx.Conn, songID uuid.UUID) (*Song, error) {
	song := &Song{SongID: songID}

	err := db.QueryRow(
		ctx,
		"SELECT name, genre, artist, length, file_path, ranking "+
			"FROM songs "+
			"WHERE song_id = $1",
		songID,
	).Scan(&song.Name, &song.Genre, &song.Artist, &song.Length, &song.FilePath, &song.Ranking)
	if err != nil {
		return nil, err
	}

	return song, nil
}

func (s *Song) Validate() (bool, string) {
	return true, ""
}

func (s *Song) Create(ctx context.Context, db *pgx.Conn) error {
	var songID string

	err := db.QueryRow(
		ctx,
		"INSERT INTO songs (name, genre, artist, length, file_path, ranking) VALUES "+
			"($1, $2, $3, $4, $5, $6)"+
			"RETURNING song_id",
		s.Name,
		s.Genre,
		s.Artist,
		s.Length,
		s.FilePath,
		s.Ranking,
	).Scan(&songID)
	if err != nil {
		return err
	}

	songUUID, err := uuid.Parse(songID)
	if err != nil {
		return err
	}

	s.SongID = songUUID

	return nil
}

func (s *Song) Update(ctx context.Context, db *pgx.Conn) error {
	_, err := db.Exec(
		ctx,
		"UPDATE songs SET "+
			"name = $1, "+
			"genre = $2, "+
			"artist = $3, "+
			"length = $4, "+
			"file_path = $5, "+
			"ranking = $6 "+
			"WHERE song_id = $7",
		s.Name,
		s.Genre,
		s.Artist,
		s.Length,
		s.FilePath,
		s.Ranking,
		s.SongID,
	)

	return err
}
