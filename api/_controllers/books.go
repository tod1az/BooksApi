package controllers

import (
	dataBase "booksapi/api/_db"
	"database/sql"
	"log"
	"strconv"
)

type Book struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func GetBooks() []Book {
	db, err := dataBase.GetDBConnection()
	if err != nil{
		log.Fatal(err)
	}
	rows, queryErr := db.Query("SELECT * FROM BOOKS")

	if queryErr != nil{
		log.Fatal(err)
	}
	defer rows.Close()
	var books  []Book

	for rows.Next(){
		var newBook  Book
		err:= rows.Scan(&newBook.Id, &newBook.Author, &newBook.Title)
		if err != nil{
			log.Fatal(err)
		}
		books = append(books, newBook)
	}
	return books
}

func GetBookById(id string) *Book {
	db, err := dataBase.GetDBConnection()
	if err != nil{
		log.Fatal(err)
	}
	var book Book
	stmt, prepareErr := db.Prepare("SELECT ID, AUTHOR, TITLE FROM BOOKS WHERE ID =$1")
	if prepareErr != nil{
		log.Fatal(prepareErr)
	}
	defer stmt.Close()

	queryErr := stmt.QueryRow(id).Scan(&book.Id,&book.Author, &book.Title)

	if queryErr !=nil{
		return nil
	}
	return &book
}

func CreateBook(newBook *Book)(string, string){
	db, err := dataBase.GetDBConnection()
	if err != nil{
		log.Fatal(err)
	}
	stmt, stmtErr := db.Prepare("INSERT INTO BOOKS (TITLE, AUTHOR) VALUES($1,$2)")
	if stmtErr != nil{
		return "", "Database Error prepare"
	}
	res, execErr := stmt.Exec(&newBook.Title, &newBook.Author)
	if execErr !=nil{
		if execErr == sql.ErrNoRows{
			return "", "Not found"
		}
		return "", "DataBase Error exec"
	}
	rows, errAffected :=res.RowsAffected()
	if errAffected!=nil{
		return "" , "No rows afected"
	}
	return strconv.Itoa(int(rows)), ""
}

func UpdateBook(id string, bookFields *Book) (string, string) {
	db , err := dataBase.GetDBConnection()
	if err!= nil{
		return "", "Connection Error"
	}

	if bookFields.Author!= ""{
		stmt, stmtErr := db.Prepare("UPDATE BOOKS SET AUTHOR = $1 WHERE ID = $2")
		if stmtErr!=nil{
			return "", "Failed preparing query"
		}
		res, execErr := stmt.Exec(bookFields.Author, id)
		if execErr!=nil{
			return "","Failed executing query"
		}
		rows, rowsErr := res.RowsAffected()
		if rowsErr!=nil{
			return "","No rows afected"
		}
		return strconv.Itoa(int(rows)), ""
	}
	if bookFields.Title!=""{
		stmt, stmtErr := db.Prepare("UPDATE BOOKS SET TITLE = $1 WHERE ID = $2")
		if stmtErr!=nil{
			return "", "Failed preparing query"
		}
		res, execErr := stmt.Exec(bookFields.Title, id)
		if execErr!=nil{
			return "","Failed executing query"
		}
		rows, rowsErr := res.RowsAffected()
		if rowsErr!=nil{
			return "","No rows afected"
		}
		return strconv.Itoa(int(rows)), ""
	}
	return "" ,"No fields to update"
}

func DeleteBook(id string) (string , string) {
	db, err := dataBase.GetDBConnection()
	if err != nil{
		log.Fatal(err)
	}
	stmt, stmtErr := db.Prepare("DELETE FROM BOOKS WHERE ID=$1")
	if stmtErr != nil{
		return "", "Database Error prepare"
	}
	res, execErr := stmt.Exec(id)
	if execErr != nil{
		return "", "Database Error exec"
	}
	rows, errAffected := res.RowsAffected()
	if errAffected != nil{
		return "", "No columns afected"
	}
	return strconv.Itoa(int(rows)), ""
}