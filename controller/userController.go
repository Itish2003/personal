package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"itish.github.io/initializers"
	"itish.github.io/model"
)

func BlogCreate(c *gin.Context) {
	var body struct {
		Title   string `form:"title"`
		Content string `form:"content"`
	}

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while binding the request body",
		})
		return
	}

	blog := model.Blog{
		Title:   body.Title,
		Content: body.Content,
	}

	result := initializers.CONTENTDB.Create(&blog)
	if result.Error != nil {
		log.Println("Error while creating user:", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while creating user",
		})
		return
	}

	title := blog.Title
	content := blog.Content
	createdAt := blog.CreatedAt
	c.HTML(http.StatusOK, "blogPageDone.html", gin.H{
		"Title":     title,
		"Content":   content,
		"CreatedAt": createdAt,
	})
}

func BlogEdit(c *gin.Context) {
	var body struct {
		Title   string `form:"title"`
		Content string `form:"content"`
	}

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while binding the request body",
		})
		return
	}

	var user *model.Blog
	initializers.CONTENTDB.Find(&user, "title=?", body.Title)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while finding the blog",
		})
		return
	}

	blog := model.Blog{
		Title:   body.Title,
		Content: body.Content,
	}

	result := initializers.CONTENTDB.Where("title = ?", body.Title).Updates(model.Blog{Content: body.Content})
	if result.Error != nil {
		log.Println("Error while updating content:", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while updating content",
			"error":   result.Error.Error(), // Print the specific error message.
		})
		return
	}

	title := blog.Title
	content := blog.Content
	c.HTML(http.StatusOK, "blogPageEdit.html", gin.H{
		"Title":   title,
		"Content": content,
	})
}

func SignUp(c *gin.Context) {
	var body struct {
		Username string `form:"username"`
		Email    string `form:"email"`
		Password string `form:"password"`
	}

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while binding the request body",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while hashing password",
		})
		return
	}

	user := model.User{
		Username: body.Username,
		Email:    body.Email,
		Password: string(hash),
	}

	result := initializers.DB.Create(&user)
	if result.Error != nil {
		log.Println("Error while creating user:", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while creating user",
		})
		return
	}

	c.Redirect(http.StatusSeeOther, "/v1/signupdone")

}

func Login(c *gin.Context) {
	var body struct {
		Username string `form:"username"`
		Email    string `form:"email"`
		Password string `form:"password"`
	}

	if c.ShouldBind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while reading the body",
		})
		return
	}

	var user *model.User
	initializers.DB.Find(&user, "username=?", body.Username)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while finding the user",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid password",
		})
		return
	}

	email := user.Email
	c.HTML(http.StatusOK, "loginDone.html", gin.H{
		"Email": email,
	})
}
