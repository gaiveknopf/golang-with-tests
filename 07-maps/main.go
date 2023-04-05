package main

type Dictionary map[string]string

const (
	ErrNotFound         = ErrDictionary("could not find the word you were looking for")
	ErrWordExists       = ErrDictionary("cannot add word because it already exists")
	ErrWordDoesNotExist = ErrDictionary("cannot update word because it does not exist")
)

type ErrDictionary string

func (e ErrDictionary) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	value, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return value, nil
}

func (d Dictionary) Insert(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
