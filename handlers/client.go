package handlers

import (
	"fmt"
	"node/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetClientHandler(db *database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		providedId := c.Request.URL.Query()["id"]

		if providedId != nil {
			id, err := strconv.Atoi(providedId[0])

			if err == nil {
				client, err := db.GetClientById(id)

				if err == nil {
					c.IndentedJSON(200, gin.H{
						"message": client,
					})
					return
				}

				c.IndentedJSON(404, gin.H{
					"message": "User does not exist",
				})
				return
			}
		}

		c.IndentedJSON(400, gin.H{
			"message": "Id has not been provided",
		})
	}
}

func GetClientsHandler(db *database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		clients, err := db.GetClients()

		if err != nil {
			fmt.Println(err)
			c.IndentedJSON(500, gin.H{
				"message": "An error has occured",
			})
			return
		}

		c.IndentedJSON(200, gin.H{
			"message": clients,
		})
	}
}

func AddClient(db *database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		params := c.Request.URL.Query()

		clientId := params["id"]
		clientExpiryDate := params["expiry_date"]

		if clientId != nil && clientExpiryDate != nil {
			id, _ := strconv.Atoi(clientId[0])
			expiryDate, _ := strconv.Atoi(clientExpiryDate[0])

			r, err := db.AddClient(
				id, expiryDate)

			if err != nil {
				fmt.Println(r)
				c.IndentedJSON(500, gin.H{
					"message": "An error has occured",
				})
				return
			}

			c.IndentedJSON(200, gin.H{
				"message": "Client has been added",
			})
			return
		}

		c.IndentedJSON(400, gin.H{
			"message": "Required params have not been passed",
		})
	}
}

func UpdateExpiryDateHandler(db *database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		params := c.Request.URL.Query()

		clientId := params["id"]
		clientExpiryDate := params["expiry_date"]

		if clientId != nil && clientExpiryDate != nil {
			id, _ := strconv.Atoi(clientId[0])
			expiryDate, _ := strconv.Atoi(clientExpiryDate[0])

			r, err := db.UpdateExpiryDate(id, expiryDate)

			if err != nil {
				fmt.Println(r)
				c.IndentedJSON(500, gin.H{
					"message": "An error has occured",
				})
				return
			}

			c.IndentedJSON(200, gin.H{
				"message": "Expiry date has been updated",
			})
			return
		}

		c.IndentedJSON(400, gin.H{
			"message": "Required params have not been passed",
		})
	}
}
