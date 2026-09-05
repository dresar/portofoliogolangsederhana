package handlers

import (
	"belajargolang/database"
	"belajargolang/internal/middleware"
	"belajargolang/internal/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// AdminLoginHandler handles admin login page
func AdminLoginHandler(c echo.Context) error {
	if middleware.IsAuthenticated(c) {
		return c.Redirect(http.StatusSeeOther, "/admin")
	}

	if c.Request().Method == http.MethodPost {
		password := c.FormValue("password")
		if password == middleware.AdminPassword {
			middleware.SetAuth(c, true)
			return c.Redirect(http.StatusSeeOther, "/admin")
		}
		return c.Render(http.StatusOK, "admin/login.html", map[string]interface{}{
			"Title": "Admin Login",
			"Error": "Invalid password",
		})
	}

	return c.Render(http.StatusOK, "admin/login.html", map[string]interface{}{
		"Title": "Admin Login",
	})
}

// AdminLogoutHandler handles admin logout
func AdminLogoutHandler(c echo.Context) error {
	middleware.SetAuth(c, false)
	return c.Redirect(http.StatusSeeOther, "/admin/login")
}

// AdminDashboardHandler handles admin dashboard
func AdminDashboardHandler(c echo.Context) error {
	var projects []models.Project
	database.DB.Find(&projects)

	return c.Render(http.StatusOK, "admin/dashboard.html", map[string]interface{}{
		"Title":    "Admin Dashboard",
		"Projects": projects,
	})
}

// AdminProjectsListHandler handles listing all projects
func AdminProjectsListHandler(c echo.Context) error {
	var projects []models.Project
	database.DB.Find(&projects)

	return c.Render(http.StatusOK, "admin/projects.html", map[string]interface{}{
		"Title":    "Manage Projects",
		"Projects": projects,
	})
}

// AdminProjectCreateHandler handles creating a new project
func AdminProjectCreateHandler(c echo.Context) error {
	if c.Request().Method == http.MethodPost {
		project := models.Project{
			Title:       c.FormValue("title"),
			Description: c.FormValue("description"),
			ImageURL:    c.FormValue("image_url"),
			TechStack:   c.FormValue("tech_stack"),
			DemoLink:    c.FormValue("demo_link"),
			GithubLink:  c.FormValue("github_link"),
		}

		if err := database.DB.Create(&project).Error; err != nil {
			return c.Render(http.StatusOK, "admin/project_form.html", map[string]interface{}{
				"Title":   "Create Project",
				"Project": project,
				"Error":   "Failed to create project",
			})
		}

		return c.Redirect(http.StatusSeeOther, "/admin/projects")
	}

	return c.Render(http.StatusOK, "admin/project_form.html", map[string]interface{}{
		"Title":   "Create Project",
		"Project": models.Project{},
	})
}

// AdminProjectEditHandler handles editing a project
func AdminProjectEditHandler(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/admin/projects")
	}

	var project models.Project
	if err := database.DB.First(&project, id).Error; err != nil {
		return c.Redirect(http.StatusSeeOther, "/admin/projects")
	}

	if c.Request().Method == http.MethodPost {
		project.Title = c.FormValue("title")
		project.Description = c.FormValue("description")
		project.ImageURL = c.FormValue("image_url")
		project.TechStack = c.FormValue("tech_stack")
		project.DemoLink = c.FormValue("demo_link")
		project.GithubLink = c.FormValue("github_link")

		if err := database.DB.Save(&project).Error; err != nil {
			return c.Render(http.StatusOK, "admin/project_form.html", map[string]interface{}{
				"Title":   "Edit Project",
				"Project": project,
				"Error":   "Failed to update project",
			})
		}

		return c.Redirect(http.StatusSeeOther, "/admin/projects")
	}

	return c.Render(http.StatusOK, "admin/project_form.html", map[string]interface{}{
		"Title":   "Edit Project",
		"Project": project,
	})
}

// AdminProjectDeleteHandler handles deleting a project
func AdminProjectDeleteHandler(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/admin/projects")
	}

	database.DB.Delete(&models.Project{}, id)
	return c.Redirect(http.StatusSeeOther, "/admin/projects")
}

