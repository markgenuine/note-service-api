package note_v1

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	desc "github.com/markgenuine/note-service-api/pkg/note_v1"
)

func (s *Note) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	fmt.Println("Get Note")
	fmt.Println("Note id: ", req.GetId())

	db, err := sqlx.Open("pgx", getDbDSN())
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query, args, err := sq.Select("title, text, author").
		From(noteTable).
		Where(sq.Eq{"id": req.GetId()}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	noteInfo := &desc.NoteParams{}
	err = db.GetContext(ctx, noteInfo, query, args...)
	if err != nil {
		return nil, err
	}

	return &desc.GetResponse{
		Id:   req.GetId(),
		Note: noteInfo,
	}, nil
}
