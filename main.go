package main

import(
  "encoading/json"
  "log"
  "net/http"
  "strings"
  "time"
  "context"
  "os"
  "os/signal"

  "github.com/go-chi/chi"
  "github.com/go-chi/chi/middleware"
  "github.com/thedevsaddam/renderer"
  mgo "gopkg.in/mgo.v2"
  "gopkg.in/mgo/bson"
)

var rnd *renderer.Render
var db *mgo.Database

const(
  hostName        string = "localhost:27017"
  dbName          string = "demo_todo"
  collectionName  string = "todo"
  port            string = ":9000"
)


type(
  todoModel struct {
    ID      bson.ObjectId `bson:"_id,omitempty"`
    Title   string `bson:"title"`
    Completed bool `bson:"completed"`
    CreatedAt time.Time `bson:"createAt"`
  }
  todo struct {
    ID        string `json:"id"`
    Title     string `json:"title"`
    Completed string `json:"completed"`
    CreatedAt time.Time `json:"created_at"`
  }
)


func init(){
  rnd = renderer.New()
  sess, err := mgo.Dial(hostName)
  checkErr(err)
  sess.SetMode(mgo.Monotonic, true)
  db = sess.DB(dbName)
}


