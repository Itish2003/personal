package service

import (
	"errors"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"itish.github.io/initializers"
	"itish.github.io/model"
)

func HomePage(c *gin.Context) {
	tmpl, err := template.ParseFiles("html/index.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(c.Writer, nil)
}

func BlogPost(c *gin.Context) {
	tmpl, err := template.ParseFiles("html/blogPage.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(c.Writer, nil)
}

func BlogPostDone(c *gin.Context) {
	tmpl, err := template.ParseFiles("html/blogPageDone.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(c.Writer, nil)
}

// func BlogEditPage(c *gin.Context) {
// 	tmpl, err := template.ParseFiles("html/blogPageEdit.html")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	tmpl.Execute(c.Writer, nil)
// }

// fetchBlogByTitle retrieves a blog post based on the title from the database.
func fetchBlogByTitle(title string) (*model.Blog, error) {
	var blog model.Blog
	if err := initializers.CONTENTDB.Find(&blog, "title = ?", title).Error; err != nil {
		log.Println("Error fetching blog by title:", err)
		return nil, err
	}
	if blog.ID == 0 {
		log.Println("Error fetching blog by title:")
		return nil, errors.New("blog not found")
	}
	return &blog, nil
}

func BlogEditPage(c *gin.Context) {
	title := c.Query("title")
	// Fetch the blog post with the specified title from the database
	// You'll need to implement this function based on your data model.
	blog, err := fetchBlogByTitle(title)
	if err != nil {
		// Handle error (e.g., blog not found)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error fetching blog for editing",
		})
		return
	}

	tmpl, err := template.ParseFiles("html/blogPageEdit.html")
	if err != nil {
		log.Println("Error parsing template:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}
	if err := tmpl.Execute(c.Writer, gin.H{
		"Title":   blog.Title,
		"Content": blog.Content,
	}); err != nil {
		log.Println("Error executing template:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return
	}

}

func LoginPage(c *gin.Context) {
	tmpl, err := template.ParseFiles("html/login.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(c.Writer, nil)
}

func SignUpDone(c *gin.Context) {
	tmpl, err := template.ParseFiles("html/signupDone.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(c.Writer, nil)
}

func SignUpPage(c *gin.Context) {
	tmpl, err := template.ParseFiles("html/signup.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(c.Writer, nil)
}

func LoginDone(c *gin.Context) {
	tmpl, err := template.ParseFiles("html/loginDone.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(c.Writer, nil)
}
