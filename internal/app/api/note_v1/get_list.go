package note_v1

import (
	"context"
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	desc "github.com/markgenuine/note-service-api/pkg/note_v1"
)

func (s *Note) GetList(ctx context.Context, req *desc.GetListRequest) (*desc.GetListResponse, error) {
	fmt.Println("Get List Note, limit: ", req.GetLimit(), " , offset: ", req.GetOffset())

	db, err := sqlx.Open("pgx", getDbDSN())
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query, args, err := sq.Select("title, text, author").
		From(noteTable).
		OrderBy("id").
		ToSql()

	if err != nil {
		return nil, err
	}

	notes := []*desc.NoteParams{}

	rows, err := db.QueryxContext(ctx, query, args...)
	for rows.Next() {
		note := &desc.NoteParams{}
		err := rows.StructScan(note)
		if err != nil {
			log.Fatalln(err)
		}
		notes = append(notes, note)
	}

	return &desc.GetListResponse{Notes: notes}, nil
}
