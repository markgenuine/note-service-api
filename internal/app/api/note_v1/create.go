package note_v1

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	desc "github.com/markgenuine/note-service-api/pkg/note_v1"
)

func (s *Note) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	fmt.Println("Create Note")
	fmt.Println("Note params: ", req.GetNote())

	db, err := sqlx.Open("pgx", getDbDSN())
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query, args, err := sq.Insert(noteTable).
		PlaceholderFormat(sq.Dollar).
		Columns("title, text, author").
		Values(req.GetNote().Title, req.GetNote().Text, req.GetNote().Author).
		Suffix("returning id").
		ToSql()

	if err != nil {
		return nil, err
	}

	row, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	row.Next()
	var id int64
	err = row.Scan(&id)
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
