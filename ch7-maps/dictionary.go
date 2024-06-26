package main

import (
	"errors"
	"sort"
)

type Dictionary map[string]string

type DictionaryErr string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryItem struct {
	word       string
	definition string
}

func (d Dictionary) GetAll() (items []DictionaryItem) {
	for k, v := range d {
		items = append(items, DictionaryItem{k, v})
	}

	return items
}

func (d Dictionary) GetAllSortedByValue() (items []DictionaryItem) {
	var values []string
	for _, v := range d {
		values = append(values, v)
	}

	sort.Strings(values)
	for _, v := range values {
		items = append(items, DictionaryItem{v, d[v]})
	}

	return items
}

func (d Dictionary) GetAllSorted() (items []DictionaryItem) {
	var keys []string
	for k := range d {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, k := range keys {
		items = append(items, DictionaryItem{k, d[k]})
	}

	return items
}

func (d Dictionary) DeleteAll() {
	clear(d)
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
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

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch {
	case errors.Is(err, ErrNotFound):
		d[word] = definition
	case err == nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}
