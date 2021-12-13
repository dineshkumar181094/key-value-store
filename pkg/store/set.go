package store

import (
	"errors"
	"log"
)
func Set(store *map[string]string, key string, value string) (bool,error) {
	_, exists := (*store)[key]
	log.Print(*store)
	if exists  {
		return false , errors.New("KeyAlreadyExists: key already exists please use update method")
	}
	(*store)[key] = value
	return true, nil

}