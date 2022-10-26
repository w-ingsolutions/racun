package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/w-ingsolutions/racun/sets"
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
	app.Get("/:budzet<float>", func(c *fiber.Ctx) error {
		budzet := c.Params("budzet")
		b, _ := strconv.ParseFloat(budzet, 8)
		xxx, _ := calculation(b, 1)

		fmt.Println("xxxxxxxxxxxxxxx")
		fmt.Println(xxx)

		fmt.Println("BRAVO TECHO")

		return c.SendString(fmt.Sprint(xxx)) // => ✋ register
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

	rezultat := []float64{0.0, 0.0}

	rezultat[0] = sets.ArhitektonskiNivoi[0][0] / (math.Pow(budzet(b), sets.ArhitektonskiNivoi[0][1]))
	return rezultat[0], rezultat[1]
}

func budzet(budzet float64) (koeficijent float64) {
	for i, b := range sets.Arhitektonski {
		if i < len(sets.Arhitektonski)-1 {
			if budzet >= b.Budzet && budzet <= sets.Arhitektonski[i+1].Budzet {
				koeficijent = b.Koeficijent
			}
		}
	}
	return
}
