package note_v1

import (
	"context"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	desc "github.com/markgenuine/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Note) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	fmt.Println("Update Note id: ", req.GetId())

	db, err := sqlx.Open("pgx", getDbDSN())
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query, args, err := sq.Update(noteTable).
		Set(titleColumn, req.Note.GetTitle()).
		Set(textColumn, req.Note.GetText()).
		Set(authorColumn, req.Note.GetAuthor()).
		Set(updatedAtColumn, time.Now()).
		Where(sq.Eq{idColumn: req.Id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	_, err = db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	fmt.Println("Update note with id: ", req.GetId())

	return &emptypb.Empty{}, nil
}
