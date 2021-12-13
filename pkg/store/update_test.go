package store

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

//success key
func TestUpdateExistingKey(t *testing.T) {
	store := make(map[string]string)
	store["testkey"] = "testvalue"
	key := "testkey"
	value := "override"

	_, err := Update(&store, key, value)
	assert.Equal(t, store["testkey"], "override", "Expected key and real are not equal")
	assert.Equal(t, err, nil, "error is not nil")

}

func TestUpdateNonExistingKey(t *testing.T) {
	store := make(map[string]string)
	store["testkey"] = "testvalue"
	key := "newkey"
	value := "override"

	_, err := Update(&store, key, value)
	expectedError := errors.New("KeyNotFound: Key Not Found")
	assert.Equal(t, expectedError, err, "Expected key and real are not equal")

}
