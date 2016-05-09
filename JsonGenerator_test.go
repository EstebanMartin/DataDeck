package main

import "testing"

type EmptyDataInteractor struct{}

func (t EmptyDataInteractor) FilterSongs(parameters map[string]string) []*Song {
	song := Song{Artist: "AC/DC", Name: "Thunderstruck", Genre: "Rock", Length: 192}

	isCorrectArtist := parameters["artist"] == "AC/DC" || parameters["artist"] == ""
	isCorrectName := parameters["name"] == "Thunderstruck" || parameters["name"] == ""
	isCorrectGenre := parameters["genre"] == "Rock" || parameters["genre"] == ""

	if isCorrectName && isCorrectArtist && isCorrectGenre {
		return []*Song{&song}
	}

	return []*Song{}
}

func TestIncorrectParameter(t *testing.T) {
	di := EmptyDataInteractor{}
	param := map[string]string{"artist": "io"}
	json := JsonGenerator{data: di}.RespondSearchSongs(param)

	expected := "[]"
	result := string(json)

	if expected != result {
		t.Fatalf("Expected %s, got %s", expected, result)
	}
}

func TestCorrectArtistWrongName(t *testing.T) {
	di := EmptyDataInteractor{}
	param := map[string]string{"artist": "AC/DC", "name": "pio pio"}
	json := JsonGenerator{data: di}.RespondSearchSongs(param)

	expected := "[]"
	result := string(json)

	if expected != result {
		t.Fatalf("Expected %s, got %s", expected, result)
	}
}

func TestCorrectArtist(t *testing.T) {
	di := EmptyDataInteractor{}
	param := map[string]string{"artist": "AC/DC"}
	json := JsonGenerator{data: di}.RespondSearchSongs(param)

	expected := "[{\"artist\":\"AC/DC\",\"song\":\"Thunderstruck\",\"genre\":\"Rock\",\"length\":192}]"
	result := string(json)

	if expected != result {
		t.Fatalf("Expected %s, got %s", expected, result)
	}
}
