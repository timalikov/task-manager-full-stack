// ex_6/controllers_test.go
package ex_6

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sql.DB {
	connStr := "host=localhost user=timalikov password=1234 dbname=go_assignment2 sslmode=disable"
	db, _ := sql.Open("postgres", connStr)
	return db
}

func TestGetUsersSQL(t *testing.T) {
	db := setupTestDB()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest("GET", "/sql/users?page=1&page_size=10&age=25", nil)

	GetUsersSQL(db)(c)

	assert.Equal(t, http.StatusOK, w.Code)

}
