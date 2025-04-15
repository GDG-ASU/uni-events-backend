package utils

import "github.com/labstack/echo/v4"

func GetClerkUserID(c echo.Context) string {
    if userID, ok := c.Get("user_id").(string); ok {
        return userID
    }
    return ""
}
