package note_v1

import (
	"context"
	"fmt"

	desc "github.com/markgenuine/note-service-api/pkg/note_v1"
)

func (s *Note) UpdateNote(ctx context.Context, req *desc.UpdateNoteRequest) (*desc.UpdateNoteResponse, error) {
	fmt.Println("Update Note id: ", req.GetId())

	return &desc.UpdateNoteResponse{
		Id:     req.GetId(),
		Status: true,
	}, nil
}
