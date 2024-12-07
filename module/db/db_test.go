package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func GetDB(t *testing.T) *database {
	db, err := New()
	assert.NoError(t, err)
	return db
}
