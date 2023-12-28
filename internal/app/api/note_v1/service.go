package note_v1

import (
	"fmt"

	desc "github.com/markgenuine/note-service-api/pkg/note_v1"
)

const (
	noteTable  = "note"
	host       = "localhost"
	port       = "54321"
	dbUser     = "service-user"
	dbPassword = "service-password"
	dbName     = "service-db"
	sslMode    = "disable"

	// columns
	idColumn        = "id"
	titleColumn     = "title"
	textColumn      = "text"
	authorColumn    = "author"
	updatedAtColumn = "updated_at"
)

type Note struct {
	desc.UnimplementedNoteServiceServer
}

func NewNote() *Note {
	return &Note{}
}

func getDbDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s  sslmode=%s",
		host, port, dbUser, dbPassword, dbName, sslMode)
}
