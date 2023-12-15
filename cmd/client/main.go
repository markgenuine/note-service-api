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
	res1, err := client.CreateNote(context.Background(), &desc.CreateNoteRequest{Note: &desc.NoteParams{
		Title:  "Title",
		Text:   "Text",
		Author: "Konstantin",
	}})

	if err != nil {
		log.Println("Call method CreateNote: ", err.Error())
	}

	log.Println("Create note: ", res1.Id)

	// call method DeleteNote
	res2, err := client.DeleteNote(context.Background(), &desc.DeleteNoteRequest{Id: 5})

	if err != nil {
		log.Println("Call method DeleteNote: ", err.Error())
	}

	log.Println("Delete note: ", res2.Id, " Delete: ", res2.Status)

	// call method GetListNote
	res3, err := client.GetListNote(context.Background(), &desc.GetListNoteRequest{Limit: 10})

	if err != nil {
		log.Println("Call method GetListNote: ", err.Error())
	}

	log.Println("List Note: ", res3.Notes)

	// call method GetNote
	res4, err := client.GetNote(context.Background(), &desc.GetNoteRequest{Id: 2})

	if err != nil {
		log.Println("Call method GetNote: ", err.Error())
	}

	log.Println("Note id: ", res4.Id, "Note info: ", res4.Note)

	// call method UpdateNote
	res5, err := client.UpdateNote(context.Background(), &desc.UpdateNoteRequest{Id: 23})

	if err != nil {
		log.Println("Call method UpdateNote: ", err.Error())
	}

	log.Println("Note id: ", res5.Id, "Update: ", res5.Status)

}
