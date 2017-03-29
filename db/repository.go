package main

import (
	"errors"
	"log"

	"gopkg.in/mgo.v2"
	"github.com/rebecabordini/movies-api/movie"
)

const MovieCollection = "movie"

var ErrDuplicatedMovie = errors.New("Duplicated Movie")

type MovieRepository struct {
	session *mgo.Session
}

func (r *MovieRepository) Create(m *movie.Movie) error {
	session := r.session.Clone()
	defer session.Close()

	collection := session.DB("").C(MovieCollection)
	err := collection.Insert(m)
	if mongoErr, ok := err.(*mgo.LastError); ok {
		if mongoErr.Code == 11000 {
			return ErrDuplicatedMovie
		}
	}
	return err
}

func NewMovieRepository(session *mgo.Session) *MovieRepository {
	return &MovieRepository{session}
}

func main() {
	session, err := mgo.Dial("localhost:27017/go-course")

	if err != nil {
		log.Fatal(err)
	}

	repository := NewMovieRepository(session)

	person := movie.Person{ Name: "Sofia Copolla", Age: 42 }
	movie := movie.Movie{ Title: "Vicky Cristina Barcelona", Director: person, Year: "2004", Stars:[]movie.Person{person} }

	err = repository.Create(&movie)

	if err == ErrDuplicatedMovie {
		log.Printf("%s is already created\n", movie.Title)
	}
}