package main

import (
	"context"
	"log"

	desc "github.com/markgenuine/note-service-api/pkg/note_v1"
	"google.golang.org/grpc"
)

const address = "localhost:50051"

func main() {
	connect, err := grpc.Dial(address, grpc.WithInsecure())
	defer connect.Close()
	if err != nil {
		log.Fatalf("didn't connect: %s", err.Error())
	}

	client := desc.NewNoteServiceClient(connect)

	// call method CreateNote
	createNote, err := client.Create(context.Background(), &desc.CreateRequest{Note: &desc.NoteParams{
		Title:  "Title",
		Text:   "Text",
		Author: "Konstantin",
	}})

	if err != nil {
		log.Println("Call method CreateNote: ", err.Error())
	}

	log.Println("Create note: ", createNote.Id)

	// call method DeleteNote
	_, err = client.Delete(context.Background(), &desc.DeleteRequest{})

	if err != nil {
		log.Println("Call method DeleteNote: ", err.Error())
	}

	// call method GetListNote
	listNote, err := client.GetList(context.Background(), &desc.GetListRequest{Limit: 10, Offset: 2})

	if err != nil {
		log.Println("Call method GetListNote: ", err.Error())
	}

	log.Println("List Note: ", listNote.Notes)

	// call method GetNote
	note, err := client.Get(context.Background(), &desc.GetRequest{Id: 2})

	if err != nil {
		log.Println("Call method GetNote: ", err.Error())
	}

	log.Println("Note id: ", note.Id, "Note info: ", note.Note)

	// call method UpdateNote
	_, err = client.Update(context.Background(), &desc.UpdateRequest{})

	if err != nil {
		log.Println("Call method UpdateNote: ", err.Error())
	}
}
