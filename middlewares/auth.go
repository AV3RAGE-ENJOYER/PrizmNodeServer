package middlewares

import (
	"node/database"

	"github.com/gin-gonic/gin"
)

func ApiTokenAuthMiddleware(db *database.Database) gin.HandlerFunc {
	apiKeys, _ := db.GetApiKeys()

	return func(c *gin.Context) {
		providedApiKey := c.Request.URL.Query()["api_key"]

		if providedApiKey != nil {
			for _, apiKey := range apiKeys {
				if apiKey.ApiKey == providedApiKey[0] { // && apiKey.AllowedIPs == c.ClientIP() {
					c.Next()
					return
				}
			}
		}

		c.IndentedJSON(401, gin.H{
			"message": "Invalid Api Key",
		})
		c.Abort()
	}
}
