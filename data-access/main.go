package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Album represents a record album
type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func main() {
	// Print environment variables for debugging
	fmt.Printf("DBUSER: %s, DBPASS: %s\n", os.Getenv("DBUSER"), os.Getenv("DBPASS"))

	// Capture connection properties
	cfg := mysql.Config{
		User:                 strings.TrimSpace(os.Getenv("DBUSER")),
		Passwd:               strings.TrimSpace(os.Getenv("DBPASS")),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "recordings",
		AllowNativePasswords: true,
	}

	// Print DSN for debugging
	fmt.Println("DSN:", cfg.FormatDSN())

	// Get a database handle
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	// Test the database connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	fmt.Println("Connected to the database!")

	// Fetch albums by a specific artist
	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatalf("Error fetching albums: %v", err)
	}
	fmt.Printf("Albums found: %v\n", albums)

	// Fetch an album by ID
	alb, err := albumByID(2)
	if err != nil {
		log.Fatalf("Error fetching album: %v", err)
	}
	fmt.Printf("Album found: %v\n", alb)

	// Add a new album
	newAlbum := Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	}
	albID, err := addAlbum(newAlbum)
	if err != nil {
		log.Fatalf("Error adding album: %v", err)
	}
	fmt.Printf("ID of added album: %v\n", albID)
}

// albumsByArtist queries for albums that have the specified artist name
func albumsByArtist(name string) ([]Album, error) {
	// An albums slice to hold data from returned rows
	var albums []Album

	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}

// albumByID queries for the album with the specified ID
func albumByID(id int64) (Album, error) {
	// An album to hold data from the returned row
	var alb Album

	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumByID %d: no such album", id)
		}
		return alb, fmt.Errorf("albumByID %d: %v", id, err)
	}
	return alb, nil
}

// addAlbum adds the specified album to the database, returning the album ID of the new entry
func addAlbum(alb Album) (int64, error) {
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}
