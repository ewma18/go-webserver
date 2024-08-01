package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

// NoSurf is the csrf protection middleware
func NoSurf(context *fiber.Ctx) {
	/*csrfHandler := nosurf.New(context.Next)
	config := config.GetAppConfig()

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   config.InProduction,
		SameSite: http.SameSiteStrictMode,
	})*/

}
