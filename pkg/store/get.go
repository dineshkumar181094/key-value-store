package store

import (
	"errors"
	"log"
)
// This function will get the key from the gloabl store
func Get(store *map[string]string, key string) (string,error) {
	value, exists := (*store)[key]
	log.Print(*store)
	log.Print(value,exists)
	if !exists  {
		return "" , errors.New("KeyNotFound: Key Not Found")
	}
	return value, nil

}