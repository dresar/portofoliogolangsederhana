package handlers

import (
	"belajargolang/database"
	"belajargolang/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// HomeHandler handles the home page
func HomeHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "public/home.html", map[string]interface{}{
		"Title": "Home - Portfolio",
	})
}

// AboutHandler handles the about page
func AboutHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "public/about.html", map[string]interface{}{
		"Title": "About Me - Portfolio",
	})
}

// SkillsHandler handles the skills page
func SkillsHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "public/skills.html", map[string]interface{}{
		"Title": "Skills - Portfolio",
	})
}

// ExperienceHandler handles the experience page
func ExperienceHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "public/experience.html", map[string]interface{}{
		"Title": "Experience - Portfolio",
	})
}

// ProjectsHandler handles the projects page
func ProjectsHandler(c echo.Context) error {
	var projects []models.Project
	database.DB.Find(&projects)

	return c.Render(http.StatusOK, "public/projects.html", map[string]interface{}{
		"Title":    "Projects - Portfolio",
		"Projects": projects,
	})
}

// ServicesHandler handles the services page
func ServicesHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "public/services.html", map[string]interface{}{
		"Title": "Services - Portfolio",
	})
}

// BlogHandler handles the blog page
func BlogHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "public/blog.html", map[string]interface{}{
		"Title": "Blog - Portfolio",
	})
}

// ContactHandler handles the contact page
func ContactHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "public/contact.html", map[string]interface{}{
		"Title": "Contact - Portfolio",
	})
}
