package main

import "fmt"
import "database/sql"
import _ "github.com/mattn/go-sqlite3"

type DataBaseController struct{}

//From the query result, converts each row to a member of an array, or well, the object Song

func (dbc DataBaseController) FilterSongs(parameters map[string]string) []*Song {
	query := queryConstructor(parameters)
	rows, err := runQuery(query)
	songs := []*Song{}

	for rows.Next() {
		var artist string
		var name string
		var genre string
		var length int

		err = rows.Scan(&artist, &name, &genre, &length)
		if err != nil {
			continue
		}
		song := &Song{Artist: artist, Name: name, Genre: genre, Length: length}
		songs = append(songs, song)
	}

	return songs
}

//Executes received query and returns the rows and possible errors thrown as a result from the query

func runQuery(query string) (*sql.Rows, error) {
	db, err := sql.Open("sqlite3", "./jrdd.db")
	defer db.Close()

	checkErr(err)
	return db.Query(query)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//It builds the query from the parameters given.
//It ignores any parameter that its key is not declared in columnName
//It lso ignores any parameter sent empty ("")
//Returns tha constructed query

func queryConstructor(params map[string]string) string {
	query := "SELECT  songs.artist, songs.song, genres.name, songs.length FROM songs LEFT OUTER JOIN genres ON songs.genre = genres.ID AND songs.genre LIKE genres.id"
	parameters := parameterFilter(params)

	if len(parameters) > 0 {
		query += " WHERE "
	}

	for column, value := range parameters {
		query += parameterConstructor(column, value)
		delete(parameters, column)
		if len(parameters) > 0 {
			query += " AND "
		}
	}

	return query
}

//Contains the name of the recpective column name in the DataBase

var columnName = map[string]string{
	"name":   "songs.song",
	"artist": "songs.artist",
	"genre":  "genres.name",
}

//Generates a map with the valid parameters found

func parameterFilter(parameters map[string]string) map[string]string {
	result := make(map[string]string)
	for key, value := range parameters {
		column := columnName[key]
		if value != "" && column != "" {
			result[column] = value
		}
	}
	return result
}

//Given a value and the column it belongs, generates the comparison line for the query

func parameterConstructor(column string, parameter string) string {
	if parameter == "" {
		return ""
	} else {
		return fmt.Sprintf("%s LIKE \"%s\"", column, parameter)
	}
}
