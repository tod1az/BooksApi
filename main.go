package main

import (
	"booksapi/controllers"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Error struct{
  Message string `json:"message"` 
}

type Result struct{
  Message string `json:"message"`
}

func main(){
	r := mux.NewRouter()
    r.HandleFunc("/books", getBooksHandler).Methods("GET")
    r.HandleFunc("/books/{id}", getBookHandler).Methods("GET")
    r.HandleFunc("/books", createBookHandler).Methods("POST")
    r.HandleFunc("/books/{id}", updateBookHandler).Methods("PUT")
    r.HandleFunc("/books/{id}", deleteBookHandler).Methods("DELETE")
    http.ListenAndServe(":8000", r)
}

func getBooksHandler(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(controllers.GetBooks())
}

func getBookHandler(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    book := controllers.GetBookById(params["id"])
    if book == nil {
      var err Error
      err.Message = "Book does not exist"
      // agrega el codigo de status de la peticion
      w.WriteHeader(http.StatusNotFound)
      json.NewEncoder(w).Encode(err)
        return
    }
    json.NewEncoder(w).Encode(book)
}

func createBookHandler(w http.ResponseWriter, r *http.Request) {
    var book controllers.Book
    _ = json.NewDecoder(r.Body).Decode(&book)
   newId, err:= controllers.CreateBook(&book)
     if err != ""{
           var newErr Error
             newErr.Message = err
        json.NewEncoder(w).Encode(newErr)
        return
     }
     var result Result
     result.Message = "Book inserted in db: "+newId
    json.NewEncoder(w).Encode(result)
}

func updateBookHandler(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var newBook controllers.Book
    _ = json.NewDecoder(r.Body).Decode(&newBook)
    id, err := controllers.UpdateBook(params["id"],&newBook)
    if err!= ""{
      error  := Error { Message : err}
      json.NewEncoder(w).Encode(error)
      return
    }
    res := Result { Message: id}
    json.NewEncoder(w).Encode(res)
}

func deleteBookHandler(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    newId, err:= controllers.DeleteBook(params["id"])

    if err != ""{
      var error Error
      error.Message = "Not found"
      json.NewEncoder(w).Encode(error)
      return
    }
    var result Result 
     result.Message = newId
     json.NewEncoder(w).Encode(result)
}
