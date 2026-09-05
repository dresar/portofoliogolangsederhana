package middleware

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

const (
	AdminSessionCookie = "admin_session"
	AdminPassword      = "admin123" // Change this in production!
	SessionValue       = "authenticated"
)

// RequireAuth middleware checks if user is authenticated
func RequireAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cookie, err := c.Cookie(AdminSessionCookie)
			if err != nil || cookie.Value != SessionValue {
				return c.Redirect(http.StatusSeeOther, "/admin/login")
			}
			return next(c)
		}
	}
}

// SetAuth sets authentication cookie
func SetAuth(c echo.Context, authenticated bool) {
	if authenticated {
		cookie := new(http.Cookie)
		cookie.Name = AdminSessionCookie
		cookie.Value = SessionValue
		cookie.Expires = time.Now().Add(24 * time.Hour)
		cookie.Path = "/"
		cookie.HttpOnly = true
		c.SetCookie(cookie)
	} else {
		cookie := new(http.Cookie)
		cookie.Name = AdminSessionCookie
		cookie.Value = ""
		cookie.Expires = time.Now().Add(-1 * time.Hour)
		cookie.Path = "/"
		c.SetCookie(cookie)
	}
}

// IsAuthenticated checks if user is authenticated
func IsAuthenticated(c echo.Context) bool {
	cookie, err := c.Cookie(AdminSessionCookie)
	return err == nil && cookie.Value == SessionValue
}

