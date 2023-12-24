package note_v1

import (
	"context"
	"fmt"

	desc "github.com/markgenuine/note-service-api/pkg/note_v1"
)

func (s *Note) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	fmt.Println("Get Note")
	fmt.Println("Note id: ", req.GetId())

	params := &desc.NoteParams{Title: "title1", Text: "text test text", Author: "Konstantin"}
	fmt.Println("Note info: ", params)

	return &desc.GetResponse{
		Id:   req.GetId(),
		Note: params,
	}, nil
}
