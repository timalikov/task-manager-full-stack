package ex_6

import (
	"assignment_2/ex_1"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetUsersSQL(t *testing.T) {
	w := httptest.NewRecorder()
	_, db := ex_1.ConnectDB()
	defer db.Close()
	c, _ := gin.CreateTestContext(w)

	// Call your handler
	GetUsersSQL(db)(c)

	// Validate response
	assert.Equal(t, http.StatusOK, w.Code)
}
