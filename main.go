package main

import (
	"math"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type input struct {
	JariJariLingkaran float64 `json:"jari-jari-lingkaran"`
	SisiPersegi       float64 `json:"sisi-persegi"`
	AlasSegitiga      float64 `json:"alas-segitiga"`
	TinggiSegitiga    float64 `json:"tinggi-segitiga"`
}

type output struct {
	LuasLingkaran     float64 `json:"luas-Lingkaran"`
	LuasPersegi       float64 `json:"luas-Persegi"`
	LuasSegitiga      float64 `json:"luas-Segitiga"`
	KelilingLingkaran float64 `json:"keliling-Lingkaran"`
	KelilingPersegi   float64 `json:"keliling-Persegi"`
	KelilingSegitiga  float64 `json:"keliling-Segitiga"`
}

func hitungLuas(input input) output {
	luasLingkaran := math.Pi * input.JariJariLingkaran * input.JariJariLingkaran
	luasPersegi := input.SisiPersegi * input.SisiPersegi
	luasSegitiga := 0.5 * input.AlasSegitiga * input.TinggiSegitiga

	return output{
		LuasLingkaran: luasLingkaran,
		LuasPersegi:   luasPersegi,
		LuasSegitiga:  luasSegitiga,
	}
}

func hitungKeliling(input input) output {
	kelilingLingkaran := 2 * math.Pi * input.JariJariLingkaran
	kelilingPersegi := 4 * input.SisiPersegi
	kelilingSegitiga := 2*input.AlasSegitiga + 2*input.TinggiSegitiga

	return output{
		KelilingLingkaran: kelilingLingkaran,
		KelilingPersegi:   kelilingPersegi,
		KelilingSegitiga:  kelilingSegitiga,
	}
}

func hitungLuasDanKeliling(c echo.Context) error {
	var inputData input
	if err := c.Bind(&inputData); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid input")
	}

	luasData := hitungLuas(inputData)
	kelilingData := hitungKeliling(inputData)

	finalOutput := output{
		LuasLingkaran:     luasData.LuasLingkaran,
		LuasPersegi:       luasData.LuasPersegi,
		LuasSegitiga:      luasData.LuasSegitiga,
		KelilingLingkaran: kelilingData.KelilingLingkaran,
		KelilingPersegi:   kelilingData.KelilingPersegi,
		KelilingSegitiga:  kelilingData.KelilingSegitiga,
	}

	return c.JSON(http.StatusOK, finalOutput)
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/hitung", hitungLuasDanKeliling)

	e.Logger.Fatal(e.Start(":1323"))
}
