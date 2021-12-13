package store

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

//success key
func TestSetExistingKey(t *testing.T) {
	store := make(map[string]string)
	store["testkey"] = "testvalue"
	key := "testkey"
	value := "override"

	_, err := Set(&store, key, value)
	expectedError := errors.New("KeyAlreadyExists: key already exists please use update method")
	assert.Equal(t, expectedError, err, "expected error and real error are not equal")

}

func TestSetNonExistingKey(t *testing.T) {
	store := make(map[string]string)

	key := "newkey"
	value := "newvalue"
	// store[key] = value

	_, err := Set(&store, key, value)
	assert.Equal(t, store[key], value, "set value failed")
	assert.Equal(t, err, nil, "error is not nil")

}
