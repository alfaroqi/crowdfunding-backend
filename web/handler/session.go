package handler

import (
	"bwastartup/user"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type sessionHandler struct {
	userService user.Service
}

func NewSessionHandler(userService user.Service) *sessionHandler {
	return &sessionHandler{userService}
}

func (h *sessionHandler) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "session_index.html", nil)
}

func (h *sessionHandler) Create(c *gin.Context) {
	var input user.LoginUserInput

	err := c.ShouldBind(&input)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return

	}

	user, err := h.userService.LoginUser(input)
	if err != nil || user.Role != "admin" {
		c.Redirect(http.StatusFound, "login")
		return
	}

	session := sessions.Default(c)
	session.Set("userID", user.ID)
	session.Set("userName", user.Name)
	session.Save()

	c.Redirect(http.StatusFound, "/users")
}

func (h *sessionHandler) Destroy(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.Redirect(http.StatusFound, "/login")
}
