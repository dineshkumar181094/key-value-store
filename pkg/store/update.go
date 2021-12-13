package store

import "errors"
func Update(store *map[string]string, key string, value string) (bool,error) {
	_ , exists := (*store)[key]
	if !exists  {
		return false , errors.New("KeyNotFound: Key Not Found")
	}
	(*store)[key] = value
	return true, nil

}