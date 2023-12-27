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

	// call method CreateNote1
	createNote1, err := client.Create(context.Background(), &desc.CreateRequest{Note: &desc.NoteParams{
		Title:  "Title",
		Text:   "Text",
		Author: "Konstantin",
	}})

	if err != nil {
		log.Println("Call method CreateNote: ", err.Error())
	}

	log.Println("Create note: ", createNote1.Id)

	// call method CreateNote2
	createNote2, err := client.Create(context.Background(), &desc.CreateRequest{Note: &desc.NoteParams{
		Title:  "Title2",
		Text:   "Text2",
		Author: "Konstantin",
	}})

	if err != nil {
		log.Println("Call method CreateNote2: ", err.Error())
	}

	log.Println("Create note: ", createNote2.Id)

	// call method CreateNote3
	createNote3, err := client.Create(context.Background(), &desc.CreateRequest{Note: &desc.NoteParams{
		Title:  "Title3",
		Text:   "Text3",
		Author: "Konstantin",
	}})

	if err != nil {
		log.Println("Call method CreateNote3: ", err.Error())
	}

	log.Println("Create note: ", createNote3.Id)

	// call method DeleteNote
	_, err = client.Delete(context.Background(), &desc.DeleteRequest{Id: createNote1.GetId()})

	if err != nil {
		log.Println("Call method DeleteNote: ", err.Error())
	}

	// call method GetListNote
	listNote, err := client.GetList(context.Background(), &desc.GetListRequest{Limit: 10, Offset: 2})

	if err != nil {
		log.Println("Call method GetListNote: ", err.Error())
	}

	log.Println("List Note: ", listNote.Notes)

	//call method GetNote
	note, err := client.Get(context.Background(), &desc.GetRequest{Id: 2})
	if err != nil {
		log.Println("Call method GetNote: ", err.Error())
	}

	log.Println(note)

	log.Println("Note id: ", note.Id, "Note info: ", note.Note)

	// call method UpdateNote
	_, err = client.Update(context.Background(), &desc.UpdateRequest{
		Id: 3,
		Note: &desc.NoteParams{
			Title:  "newTitle",
			Text:   "newText",
			Author: "newAuthor",
		},
	})

	if err != nil {
		log.Println("Call method UpdateNote: ", err.Error())
	}
}
