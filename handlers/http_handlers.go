package handlers

import (
	"app/models"
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Fruits []models.Fruit
var MaxID int

func CreateFruit(c *gin.Context) {
	var fruit models.Fruit
	if err := c.ShouldBindJSON(&fruit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Asigna un nuevo ID a la fruta
	fruit.ID = MaxID + 1
	MaxID = fruit.ID

	Fruits = append(Fruits, fruit)
	c.JSON(http.StatusOK, gin.H{"data": fruit})
}

func GetFruits(c *gin.Context) {
	apiFruits, err := getFruitsFromAPI()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	allFruits := append(Fruits, apiFruits...)
	c.JSON(http.StatusOK, allFruits)
}

func GetFruitByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	apiFruits, err := getFruitsFromAPI()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	allFruits := append(Fruits, apiFruits...)

	for _, fruit := range allFruits {
		if fruit.ID == id {
			c.IndentedJSON(http.StatusOK, fruit)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": " Fruit not found"})
}

func getFruitsFromAPI() ([]models.Fruit, error) {
	apiUrl := os.Getenv("FRUITYVICE_API_URL")
	if apiUrl == "" {
		apiUrl = "https://www.fruityvice.com" // valor por defecto
	}
	resp, err := http.Get(apiUrl + "/api/fruit/all")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var fruits []models.Fruit
	err = json.NewDecoder(resp.Body).Decode(&fruits)
	if err != nil {
		return nil, err
	}

	// Actualiza MaxID con las frutas obtenidas de la API
	for _, fruit := range fruits {
		if fruit.ID > MaxID {
			MaxID = fruit.ID
		}
	}

	return fruits, nil
}
