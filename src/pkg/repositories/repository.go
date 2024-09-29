package repositories

import (
    "SongsLibrary/src/internal/databases"
    // "SongsLibrary/src/pkg/controllers"
    "github.com/jackc/pgx/v5"
)

type SongsRepository struct {
    conn *pgx.Conn
}

func (p *SongsRepository) Create() {

}
func (p *SongsRepository) Read() {
}
func (p *SongsRepository) Update() {

}
func (p *SongsRepository) Delete() {

}

func NewSongsRepository(conn *pgx.Conn) *Repository {
    db := databases.NewPostgresDatabase()
    songsRepo := &SongsRepository{conn: db.Connect()}
    return NewRepository(songsRepo)
}
