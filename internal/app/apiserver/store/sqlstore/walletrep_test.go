package sqlstore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	db_url := "host=localhost port=5432 user=postgres dbname=payments sslmode=disable password=2505"
	db, err := gorm.Open(postgres.Open(db_url), &gorm.Config{})

	assert.NoError(t, err, "failed to connect to the test database")
	assert.NotNil(t, db, "db instance can`t be nil")
}
