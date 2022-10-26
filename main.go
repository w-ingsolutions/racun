package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/w-ingsolutions/sets"
)

/*
Arhitektonski
Konstrukcije
Masinski
Elektro
Telekomunikacije
Vodovod i kanalizacija
Termika
Akustika
Enterijer
Spoljno
Tehnoloski projekti
*/

func main() {
	app := fiber.New()

	// GET /api/register
	app.Get("/api/*", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("✋ %s", c.Params("*"))
		return c.SendString(msg) // => ✋ register
	}).Name("api")

	data, _ := json.MarshalIndent(app.GetRoute("api"), "", "  ")
	fmt.Print(string(data))
	// Prints:
	// {
	//    "method": "GET",
	//    "name": "api",
	//    "path": "/api/*",
	//    "params": [
	//      "*1"
	//    ]
	// }

	log.Fatal(app.Listen(":3000"))
}

func calculation(b float64, nivo int) (float64, float64) {

	var rezultat []float64

	rezultat[0] = sets.ArhitektonskiNivoi[0][0]/budzet(b) ^ sets.ArhitektonskiNivoi[0][1]

	return rezultat[0], rezultat[1]
}

func budzet(budzet float64) (koeficijent float64) {
	for b, k := range sets.Arhitektonski {
		if budzet <= b {
			koeficijent = k
		}
	}
	return
}
