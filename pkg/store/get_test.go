package store

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

//success key
func TestGetExistingKey(t *testing.T) {
	var store = make(map[string]string)
	store["testkey"] = "testvalue"
	key := "testkey"

	result, err := Get(&store, key)
	assert.Equal(t, "testvalue", result, "unable to get the value")
	assert.Equal(t, nil, err, "err is not nil")
}

func TestNonExistingKey(t *testing.T) {
	var store = make(map[string]string)
	key := "nonexisting"
	_, err := Get(&store, key)
	expectedErrors := errors.New("KeyNotFound: Key Not Found")
	assert.Equal(t, expectedErrors, err, "expected error and real error are not equal")

}
