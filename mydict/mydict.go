package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not Found")
	errWordExist  = errors.New("That word already exists")
	errCantUpdate = errors.New("Cant update non-existing word")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exist := d[word]
	if exist {
		return value, nil
	}
	return "", errNotFound
}

// Add a word
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExist
	}
	return nil

	// if err == errNotFound {
	// 	d[word] = def
	// } else if err == nil {
	// 	return errWordExist
	// }
}

// Update a word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
