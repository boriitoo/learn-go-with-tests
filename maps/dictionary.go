package maps

import "errors"

type Dictionary map[string]string

var (
	ErrWordNotFound = errors.New("key not found")
	ErrWordExists = errors.New("key already exists")
)

func (d Dictionary) Search(needle string) (string, error) {
	word, succeed := d[needle]

	if !succeed {
		return "", ErrWordNotFound
	}

	return word, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	if err != nil {
		d[key] = value;
		return nil
	}

	return ErrWordExists
}