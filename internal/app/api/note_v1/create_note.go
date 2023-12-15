package note_v1

import (
	"context"
	"fmt"

	desc "github.com/markgenuine/note-service-api/pkg/note_v1"
)

func (s *Note) CreateNote(ctx context.Context, req *desc.CreateNoteRequest) (*desc.CreateNoteResponse, error) {
	fmt.Println("Create Note")
	fmt.Println("Note params: ", req.GetNote())

	return &desc.CreateNoteResponse{
		Id: 1,
	}, nil
}
