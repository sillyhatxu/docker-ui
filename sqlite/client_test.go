package sqlite

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var dataSourceName = "./test.db"

func TestTest(t *testing.T) {
	Test()
}

func TestOpenDB(t *testing.T) {
	db, err := OpenDB(dataSourceName)
	db.SetMaxOpenConns(10)
	assert.Nil(t, err)
	assert.NotNil(t, db)
}
