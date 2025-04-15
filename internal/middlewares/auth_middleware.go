package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/MicahParks/keyfunc"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

var jwks *keyfunc.JWKS
var clerkIssuer string

func init() {
    err := godotenv.Load("../../.env")
    if err != nil {
        log.Println("Could not load .env from root (../../.env)")
    }

    clerkIssuer = os.Getenv("CLERK_ISSUER")
    if clerkIssuer == "" {
        log.Fatalf("Failed to get CLERK_ISSUER from environment variables")
    }

    jwksURL := clerkIssuer + "/.well-known/jwks.json"
    jwks, err = keyfunc.Get(jwksURL, keyfunc.Options{
        RefreshInterval:     time.Hour,
        RefreshErrorHandler: func(err error) { println("Error refreshing JWKS:", err.Error()) },
    })
    if err != nil {
        panic("Failed to create JWKS from URL: " + err.Error())
    }
}


func ClerkAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			fmt.Print("ERORR")
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing token")
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenStr, jwks.Keyfunc)
		if err != nil || !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		fmt.Print("claims",claims)
		if !ok || claims["iss"] != clerkIssuer {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid issuer")
		}

		userID, ok := claims["sub"].(string)
		if !ok {
			return echo.NewHTTPError(http.StatusBadRequest, "User ID not found in token")
		}

		email, ok := claims["email"].(string)
		if !ok || email == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "Email not found in claims")
		}

		c.Set("user_id", userID)
		c.Set("email", email)

		return next(c)
	}
}