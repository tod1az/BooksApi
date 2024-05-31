package handler

import (
	controllers "booksapi/api/_controllers"
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

func Main(){
	r := mux.NewRouter()
    r.HandleFunc("/books", GetBooksHandler).Methods("GET")
    r.HandleFunc("/books/{id}", GetBookHandler).Methods("GET")
    r.HandleFunc("/books", CreateBookHandler).Methods("POST")
    r.HandleFunc("/books/{id}", UpdateBookHandler).Methods("PUT")
    r.HandleFunc("/books/{id}", DeleteBookHandler).Methods("DELETE")
    http.ListenAndServe(":8000", r)
}

func GetBooksHandler(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(controllers.GetBooks())
}

func GetBookHandler(w http.ResponseWriter, r *http.Request) {
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

func CreateBookHandler(w http.ResponseWriter, r *http.Request) {
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

func UpdateBookHandler(w http.ResponseWriter, r *http.Request) {
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

func DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
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
