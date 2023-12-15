package note_v1

import (
	"context"
	"fmt"

	desc "github.com/markgenuine/note-service-api/pkg/note_v1"
)

func (s *Note) GetListNote(ctx context.Context, req *desc.GetListNoteRequest) (*desc.GetListNoteResponse, error) {
	fmt.Println("Get List Note, limit: ", req.GetLimit())
	nP := []*desc.NoteParams{
		&desc.NoteParams{Title: "Title1", Text: "Text1", Author: "konstantin"},
		&desc.NoteParams{Title: "Title2", Text: "Text2", Author: "konstantin"},
	}
	fmt.Println("Note params: ", nP)
	return &desc.GetListNoteResponse{Notes: nP}, nil
}
