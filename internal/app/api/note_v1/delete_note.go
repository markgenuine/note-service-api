package note_v1

import (
	"context"
	"fmt"

	desc "github.com/markgenuine/note-service-api/pkg/note_v1"
)

func (s *Note) DeleteNote(ctx context.Context, req *desc.DeleteNoteRequest) (*desc.DeleteNoteResponse, error) {
	fmt.Println("Delete Note id: ", req.GetId())

	return &desc.DeleteNoteResponse{
		Id:     req.GetId(),
		Status: true,
	}, nil
}
