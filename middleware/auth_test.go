package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Agarwalsahil/TodoAPI/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware_Success(t *testing.T) {
	token, _ := utils.GenerateToken("test@example.com")

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(AuthMiddleware)

	// test route that only succeeds if middleware allows it
	r.GET("/protected", func(c *gin.Context) {
		email := c.GetString("CreatedBy")
		c.JSON(http.StatusOK, gin.H{"email": email})
	})

	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestAuthMiddleware_MissingToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(AuthMiddleware)

	r.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "Shouldn't reach here"})
	})

	req, _ := http.NewRequest("GET", "/protected", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnauthorized, resp.Code)
}
