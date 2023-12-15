package note_v1

import (
	"context"
	"fmt"

	desc "github.com/markgenuine/note-service-api/pkg/note_v1"
)

func (s *Note) GetNote(ctx context.Context, req *desc.GetNoteRequest) (*desc.GetNoteResponse, error) {
	fmt.Println("Get Note")
	fmt.Println("Note id: ", req.GetId())

	nP := &desc.NoteParams{Title: "title1", Text: "text test text", Author: "Konstantin"}
	fmt.Println("Note info: ", nP)

	return &desc.GetNoteResponse{
		Id:   req.GetId(),
		Note: nP,
	}, nil
}
