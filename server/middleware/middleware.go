package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"../models"
	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
);

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
