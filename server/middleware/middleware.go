package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-lang-todo-react/server/models"
	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB connection string
// for localhost mongoDB
// const connectionString = "mongodb://localhost:27017"
const connectionString = "mongodb+srv://awaismumtaz0099:25213291231919@cluster0.3so1bcq.mongodb.net"

// Database Name
const dbName = "todo-list-go-lang"

// Collection name
const collName = "todolist"

// collection object/instance
var collection *mongo.Collection

// IN THIS FUNCTION WE ARE BASICALLY CONNECTING TO OUR DATABASE AND THE COLLECTION
func init() {

	// HERE WE ARE USING SOME CLIENT OPTIONS!
	clientOptions := options.Client().ApplyURI(connectionString)

	// HERE WE ARE CONNECTING TO MONGODB USING MONGO DRIVER!
	client, err := mongo.Connect(context.TODO(), clientOptions)

	// IF SOME ERROR OCCUR!
	if err != nil {
		log.Fatal(err)
	}

	// ELSE LETS CONNECT TO MONOGODB!
	err = client.Ping(context.TODO(), nil)

	// ELSE WE CAN SAY THAT WE ARE CONNECTIED TO MONGODB!
	fmt.Println("Connected to MongoDB!")

	// NOW HERE WE ARE ACCESSING THE COLLECTION IN OUR DATABASE AND STORED IN OUR VARIABLE!
	collection = client.Database(dbName).collection(collName)

	fmt.Println("Collection instance created!")
}

// FUNCTION TO GET ALL TASKS!
func GET_ALL_TASKS(w http.ResponseWriter, r *http.Response) {
	// SETTING THE HEADERS!
	w.Header().Set("Content-Type", "application/json")
	// ALLOWING ACCESS ORIGIN!
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// GETTING ALL TASKS!
	// payload := getAllTasks();
	// RETURNING JSON RESPONE BACK!
	// json.NewEncoder(w).Encode(payload);
}

// FUNCTION THAT WILL CREATE TASKS!
func CREATE_TASKS(w http.ResponseWriter, r *http.Request) {

	// FOR READING THE URL FORM DATA!
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	// ALLOWING ACCESS ORIGIN!
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// SETTING THE METHOD TO BE POST!
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	// SETTING THE CONTENT TYPING!
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// CREATING NEW VARIABLE THAT WILL HOLD THE DATA FROM REQUEST!
	var task models.TodoList
	// WE USE _ SO THAT MEANS WE GONNA NOT USE THIS VARIABLE!
	_ = json.NewDecoder(r.Body).Decode(r)
	// CALLING THE FUNCTION THAT WILL CREATE NEW TODO ITEM!
	// insertNewTodo(task);
	// RETURNING THE RESPONSE BACK!
	json.NewEncoder(w).Encode(task)
}

// FUNCTION THAT WILL MARK TASK TO BE COMPLETED THS TODO!
func TaskComplete(w http.ResponseWriter, r *http.Request) {
	// FOR READING THE URL FORM DATA!
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	// ALLOWING ACCESS ORIGIN!
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// SETTING THE METHOD TO BE PUT!
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	// SETTING THE CONTENT TYPING!
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// GETTING PARAMETER FROM THE REQUEST , WE ARE USING MUX HERE!
	params := mux.Vars()
	// completeTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// UndoTask undo the complete task route
func UndoTask(w http.ResponseWriter, r *http.Request) {

	// FOR READING THE URL FORM DATA!
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	// ALLOWING ACCESS ORIGIN!
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// SETTING THE METHOD TO BE PUT!
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	// SETTING THE CONTENT TYPING!
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	// undoTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// DeleteTask delete one task route
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	// FOR READING THE URL FORM DATA!
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	// ALLOWING ACCESS ORIGIN!
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// SETTING THE METHOD TO BE DELETE!
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	// SETTING THE CONTENT TYPING!
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	// deleteOneTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])
	// json.NewEncoder(w).Encode("Task not found")

}

// FUNCTION THAT WILL DELETE THAT ALL!
func DeleteAllTasks(w http.ResponseWriter, r *http.Request) {
	// ALLOWING ACCES ORIGIN!
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// SETTING THE METHOD TO BE DELETE!
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	// CREATING AND STORING SOME NEW DATA IN OUR VARAIBLE!
	// count := deleteAllTasks();
	// json.NewEncoder(w).Encode(count);

}

// FUNCTION THAT WILL GET ALL TASKS FROM THAT DATABASE AND RETURN BACKNTOTHE CALLING FUNCTION!
func getAllTasks() []primitive.M {
	// GETTING ALL DATA RELATED TO THAT COLLECTION!
	cur, err := collection.Find(context.Background(), bson.D{{}})

	// IF SOME ERROR OCCUR HANDLE THE ERROR!
	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M

	// LOOP THROUGH THE RESULTS OF A MONGODB CURSOR UNTIL THERE ARE NO MORE DOCUMENTS

	for cur.Next(context.Background()) {
		// DECLARE A VARIABLE TO STORE EACH DOCUMENT
		var result bson.M

		// 'bson' MEANS THAT WE ARE CONVERTING OBJECT TYPE DATA TO GO LANG STRUCT

		// DECODE THE CURRENT DOCUMENT INTO THE 'result' VARIABLE
		e := cur.Decode(&result)

		// CHECK IF THERE WAS AN ERROR DURING DECODING
		if e != nil {
			// LOG THE ERROR AND TERMINATE THE PROGRAM
			log.Fatal(e)
		}

		// APPEND THE DECODED DOCUMENT TO THE RESULTS SLICE
		results = append(results, result)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())
	return results // RETURNING THE DATA BACK!
}

func insertNewTodo(task models.TodoList) {
	// INSERTING NEW DOC TO OUR TODOLIST COLLECTION IN OUR DATABASE!
	insertedResult, err := collection.InsertOne(context.Background(), task)

	if err != nil {
		log.Fatal(err)
	}

	return insertedResult
}
