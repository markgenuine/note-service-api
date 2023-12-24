package note_v1

import (
	"context"
	"fmt"

	desc "github.com/markgenuine/note-service-api/pkg/note_v1"
)

func (s *Note) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	fmt.Println("Create Note")
	fmt.Println("Note params: ", req.GetNote())

	return &desc.CreateResponse{
		Id: 1,
	}, nil
}
