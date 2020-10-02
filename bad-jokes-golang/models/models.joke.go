package models

import "errors"

// Joke contains information about a single Joke
type Joke struct {
	ID    int    `json:"id" binding:"required"`
	Likes int    `json:"likes"`
	Joke  string `json:"joke" binding:"required"`
}

// list of jokes
var jokes = []Joke{
	Joke{1, 0, "Did you hear about the restaurant on the moon? Great food, no atmosphere."},
	Joke{2, 0, "What do you call a fake noodle? An Impasta."},
	Joke{3, 0, "How many apples grow on a tree? All of them."},
	Joke{4, 0, "Want to hear a joke about paper? Nevermind it's tearable."},
	Joke{5, 0, "I just watched a program about beavers. It was the best dam program I've ever seen."},
	Joke{6, 0, "Why did the coffee file a police report? It got mugged."},
	Joke{7, 0, "How does a penguin build it's house? Igloos it together."},
}

// ReturnAllJokes returns all jokes from list
func ReturnAllJokes() []Joke {
	return jokes
}

// GetJokeByID fetch joke by ID
func GetJokeByID(id int) (*Joke, error) {
	for _, a := range jokes {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Joke not found :(")
}

// LikeJoke allows users to vote on a joke
func LikeJoke(jokeid int) (*Joke, error) {
	// search for the joke with the matching id
	for _, a := range jokes {
		if a.ID == jokeid {
			a.Likes++
		}
	}
	return nil, errors.New("an unexpected error occurred")
}

// AddNewJoke add a new joke with jokes added
func AddNewJoke(joke string) (*Joke, error) {
	// Set the ID of a new joke to one more than the number of jokes
	a := Joke{ID: len(jokes) + 1, Joke: joke}

	// Add the joke to the list of jokes
	jokes = append(jokes, a)

	return &a, nil
}
