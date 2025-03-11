package controllers

// import (
// 	"context"
// 	"crypto/tls"
// 	"fmt"
// 	"log"

// 	"github.com/AbdulWasay1207/notes-sharing-app/models"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// const connectionString = "mongodb+srv://notes:Wasay@cluster0.nwnx0.mongodb.net/?retryWrites=true&w=majority&tls=true"
// const dbName = "TempNotes"
// const colName = "notes"

// // Most Important
// var collection *mongo.Collection

// // connect with mongodb

// func init() {
// 	//client options
// 	clientOptions := options.Client().ApplyURI(connectionString)

// 	clientOptions.SetTLSConfig(&tls.Config{
// 		InsecureSkipVerify: true, // Only use this for testing; not recommended for production
// 	})

// 	//connect to mongodb
// 	client, err := mongo.Connect(context.TODO(), clientOptions)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("MONGODB CONNECTION SUCCESS")

// 	collection = client.Database(dbName).Collection(colName)

// 	fmt.Println("Collection instance is ready")
// }

// // Mongodb helpers

// // insert 1 record
// func insertOneNote(note models.Notes) string {
// 	inserted, err := collection.InsertOne(context.Background(), note)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	insertedID, _ := inserted.InsertedID.(primitive.ObjectID)
// 	fmt.Println("Inserted 1 Note in db with id : ", inserted.InsertedID)
// 	idString := string(insertedID[:])

// 	return idString
// }

// // delete 1 record
// func deleteOneNote(noteId string) error {
// 	id, err := primitive.ObjectIDFromHex(noteId)
// 	if err != nil {
// 		return fmt.Errorf("invalid id")
// 	}
// 	filter := bson.M{"_id": id}
// 	_, err = collection.DeleteOne(context.Background(), filter)
// 	if err != nil {
// 		return fmt.Errorf("id not found")
// 	}
// 	fmt.Println("Note deleted with id : ", noteId)
// 	return nil
// }

// func getAllNotes() []primitive.M {
// 	curr, err := collection.Find(context.Background(), bson.D{{}})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var notes []primitive.M
// 	for curr.Next(context.Background()) {
// 		var note bson.M
// 		err := curr.Decode(&note)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		notes = append(notes, note)
// 	}
// 	defer curr.Close(context.Background())
// 	return notes
// }

// func getOneNote(noteId string) (primitive.M, error) {
// 	id, err := primitive.ObjectIDFromHex(noteId)
// 	if err != nil {
// 		return nil, fmt.Errorf("invalid id")
// 	}
// 	var note primitive.M
// 	err = collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&note)
// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			return nil, fmt.Errorf("note not found")
// 		}
// 		return nil, fmt.Errorf("error fetching note: %v", err)
// 	}

// 	return note, nil
// }
