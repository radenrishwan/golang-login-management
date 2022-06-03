package database

import (
	"github.com/stretchr/testify/assert"
	"login-management/database"
	"testing"
)

func TestDatabaseSuccess(t *testing.T) {
	db := database.GetDatabase()

	assert.NotEmpty(t, db)
}
