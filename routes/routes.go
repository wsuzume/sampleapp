package routes

import (
	"sampleapp/sessions"
	"sampleapp/config"

	"net/http"
	"github.com/gin-gonic/gin"
)

func Home(ctx *gin.Context) {
	session := sessions.GetDefaultSession(ctx)
	buffer, exists := session.Get("user")
	if exists {
		user := buffer.(*config.DummyUserModel)
		println("Home sweet home")
		println("sessionID: " + session.ID)
		println("username: " + user.Username)
		println("email: " + user.Email)
	} else {
		println("Unhappy home")
		println("sessionID: " + session.ID)
	}

	session.Save()
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
}

func LogIn(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", gin.H{})
}

func SignUp(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "signup.html", gin.H{})
}

func NoRoute(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
}
