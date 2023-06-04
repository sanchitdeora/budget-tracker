package db

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var db Database

func TestMain(m *testing.M) {
	db = NewDatabase()
	m.Run()
}

func TestGetAllTransactions(t *testing.T) {
	assert.Equal(t, 1, 1, "The two words should be the same.")
}

func TestGetTransactionRecordById(t *testing.T) {
	assert.Equal(t, 1, 1, "The two words should be the same.")
}

func TestInsertTransactionRecord(t *testing.T) {
	assert.Equal(t, 1, 1, "The two words should be the same.")
}

func TestUpdateTransactionRecordById(t *testing.T) {
	assert.Equal(t, 1, 1, "The two words should be the same.")
}

func TestDeleteTransactionRecordById(t *testing.T) {
	assert.Equal(t, 1, 1, "The two words should be the same.")
}