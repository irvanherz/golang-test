package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/irvanherz/golang-test/functions"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

func soal0(c echo.Context) error {
	type PalindromeResult struct {
		From   int   `json:"from"`
		To     int   `json:"to"`
		Result []int `json:"result"`
		Count  int   `json:"count"`
	}

	from, err1 := strconv.Atoi(c.QueryParam("from"))
	to, err2 := strconv.Atoi(c.QueryParam("to"))
	if err1 != nil || err2 != nil {
		r := ErrorResponse{
			Success: false,
			Error:   "Incomplete parameter",
		}
		return c.JSON(400, r)
	} else {
		pals := functions.GetPalindromes(from, to)
		r := SuccessResponse{
			Success: true,
			Data: PalindromeResult{
				From:   from,
				To:     to,
				Result: pals,
				Count:  len(pals),
			},
		}
		return c.JSON(200, r)
	}
}

func soal1(c echo.Context) error {
	type SortResult struct {
		RawInput  string   `json:"raw_input"`
		Input     []string `json:"input"`
		RawResult string   `json:"raw_result"`
		Result    []string `json:"result"`
		Count     int      `json:"count"`
	}

	raw_input := c.QueryParam("input")
	if raw_input == "" {
		r := ErrorResponse{
			Success: false,
			Error:   "Incomplete parameter",
		}
		return c.JSON(400, r)
	} else {
		raw_input = strings.TrimSpace(raw_input)
		regex := regexp.MustCompile("\\s")
		input := regex.Split(raw_input, -1)
		result := functions.SortBooks(input)
		raw_result := strings.Join(result, " ")
		r := SuccessResponse{
			Success: true,
			Data: SortResult{
				RawInput:  raw_input,
				Input:     input,
				RawResult: raw_result,
				Result:    result,
				Count:     len(result),
			},
		}
		return c.JSON(200, r)
	}
}

func soal2(c echo.Context) error {
	type ScanResult struct {
		Input  string `json:"input"`
		Result []int  `json:"result"`
		Count  int    `json:"count"`
	}

	input := c.QueryParam("input")
	if input == "" {
		r := ErrorResponse{
			Success: false,
			Error:   "Incomplete parameter",
		}
		return c.JSON(400, r)
	} else {
		mys := new(functions.Mystery)
		mys.Init(input)
		result := mys.Scan()
		r := SuccessResponse{
			Success: true,
			Data: ScanResult{
				Input:  input,
				Result: result,
				Count:  len(result),
			},
		}
		return c.JSON(200, r)
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/go/soal0", soal0)
	e.GET("/go/soal1", soal1)
	e.GET("/go/soal2", soal2)
	// Start server
	e.Logger.Fatal(e.Start(":5000"))
}
