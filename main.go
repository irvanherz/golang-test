package main

import (
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Result struct {
	From  int `json:"from"`
	To    int `json:"to"`
	Count int `json:"count"`
}

func strRev(str string) string {
	newStr := ""
	for i := len(str); i != 0; i-- {
		newStr += string(str[i-1])
	}
	return newStr
}

func isPalindrome(val int) bool {
	str1 := strconv.Itoa(val)
	str2 := strRev(str1)
	if strings.Compare(str1, str2) == 0 {
		return true
	} else {
		return false
	}
}

func findPalindrome(from, to int) int {
	var n int
	if from < 1 || to > int(math.Pow(10, 9)) {
		return -1
	}
	for i := from; i <= to; i++ {
		if isPalindrome(i) {
			n++
		}
	}
	return n
}

func hello(c echo.Context) error {
	from, _ := strconv.Atoi(c.QueryParam("from"))
	to, _ := strconv.Atoi(c.QueryParam("to"))
	cnt := findPalindrome(from, to)
	r := &Result{
		From:  from,
		To:    to,
		Count: cnt,
	}
	return c.JSON(http.StatusOK, r)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/", hello)
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
	// 	var from, to int

	// askInput:
	// 	fmt.Print("Input two number separated by a space: ")
	// 	_, err := fmt.Scanf("%d %d", &from, &to)
	// 	if err != nil {
	// 		fmt.Println("\nOops, error. Make sure you're input correct numbers")
	// 		goto askInput
	// 	}
	// 	fmt.Printf("Palindrome from %d to %d: %d occurences", from, to, findPalindrome(from, to))
}
