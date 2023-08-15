package mydict

import "errors"

// Dictionary is a map of string to string
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")
var errWordExists = errors.New("That word already exists")

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = definition
	case nil:
		return errWordExists
	}
	return nil
}

// Update a word in the dictionary
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		return errNotFound
	case nil:
		d[word] = definition
	}
	return nil
}

// Delete a word from the dictionary
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
