package models

import (
	"database/sql"
	"fmt"
)

// Album represents a record in the album table
type Album struct {
	ID     int64   `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float32 `json:"price"`
}

// GetAlbums queries for all albums
func GetAlbums(db *sql.DB) ([]Album, error) {
	var albums []Album

	rows, err := db.Query("SELECT * FROM album")
	if err != nil {
		return nil, fmt.Errorf("GetAlbums: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("GetAlbums: %v", err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetAlbums: %v", err)
	}
	return albums, nil
}

// AlbumsByArtist queries for albums that have the specified artist name
func AlbumsByArtist(db *sql.DB, name string) ([]Album, error) {
	var albums []Album

	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("AlbumsByArtist %q: %v", name, err)
	}
	defer rows.Close()

	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("AlbumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("AlbumsByArtist %q: %v", name, err)
	}
	return albums, nil
}

// AlbumByID queries for the album with the specified ID
func AlbumByID(db *sql.DB, id int64) (Album, error) {
	var alb Album

	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("AlbumByID %d: no such album", id)
		}
		return alb, fmt.Errorf("AlbumByID %d: %v", id, err)
	}
	return alb, nil
}

// AddAlbum adds the specified album to the database, returning the album ID of the new entry
func AddAlbum(db *sql.DB, alb Album) (int64, error) {
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("AddAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("AddAlbum: %v", err)
	}
	return id, nil
}
