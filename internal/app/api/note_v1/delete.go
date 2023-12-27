package note_v1

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	desc "github.com/markgenuine/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Note) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	db, err := sqlx.Open("pgx", getDbDSN())
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query, args, err := sq.Delete(noteTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": req.GetId()}).
		ToSql()

	if err != nil {
		return nil, err
	}

	_, err = db.ExecContext(ctx, query, args...)

	if err != nil {
		return nil, err
	}

	fmt.Println("Delete note with id: ", req.GetId())

	return &emptypb.Empty{}, nil
}
