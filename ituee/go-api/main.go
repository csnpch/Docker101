package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)


type CatlogNodes struct {
	CatlogNodes []Catlog `json:"catlog_nodes"`
}
 
type Catlog struct {
	productId string `json: "productId"`
	quantity   int    `json: "quantity"`
}


var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
func randSeq(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}


func writeDatafile() {
	
	file, err := os.Create(pathUploads + "/dataFile.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	
	_, err2 := file.WriteString(randSeq(60) + "\n")

	if err2 != nil {
		log.Fatal(err)
	}
	
}

func readDataFile() string {
	content, err := os.ReadFile(pathUploads + "/dataFile.txt")
    if err != nil {
		os.Create(pathUploads + "/dataFile.txt")
        fmt.Printf("Error : %v", err)
	}
    fmt.Println(string(content))
	return string(content);
}

var pathUploads = ""

func main() {
	err := godotenv.Load() // or Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	
	if os.Getenv("ENVIRONMENT") == "production" {
		pathUploads = os.Getenv("PATH_UPLOADS")
	} else {
		pathUploads = "./../uploads"
	}
	println("pathUploads = " + pathUploads)
	os.Mkdir(pathUploads, 755)

	app := fiber.New()

	app.Get("/getStatus", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": fiber.StatusOK,
			"data": fiber.Map{
				"message": "API Connected",
			},
		})
	})

	app.Get("/pushData/", func(c *fiber.Ctx) error {
		writeDatafile()
		return c.JSON(fiber.Map{
			"status": fiber.StatusOK,
			"message": "success",
			"data": nil,
		})
	})

	app.Get("/getData", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": fiber.StatusOK,
			"message": "success",
			"data": fiber.Map{
				"dataInFile": readDataFile(),
			},
		})
	})

	app.Listen(":3000")
}
