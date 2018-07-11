package controllers

import (
	"blog/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetUserStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token, err := c.Cookie("token"); err == nil && config.ValidateToken(token) == nil {
			c.Set("is_logged_in", true)
		} else {
			c.Set("is_logged_in", false)
		}
	}
}

func EnsureLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)

		if !loggedIn {
			//if token, err := c.Cookie("token"); err != nil || token == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func EnsureNotLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)

		if loggedIn {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
