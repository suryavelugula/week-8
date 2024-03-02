package main_test

import ( 
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCalculatorEndpoints(t *testing.T) {
	r := setupRouter()

	testCases := []struct {
		method   string
		url      string
		expected string
	}{
		{"GET", "/", "Welcome to Simple Calculator"},
		{"GET", "/calculate?a=2&b=3&operation=add", "Result: 5"},
		{"GET", "/calculate?a=5&b=3&operation=subtract", "Result: 2"},
		{"GET", "/calculate?a=4&b=3&operation=multiply", "Result: 12"},
		{"GET", "/calculate?a=10&b=2&operation=divide", "Result: 5"},
	}

	for _, tc := range testCases {
		req, err := http.NewRequest(tc.method, tc.url, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		r.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		if rr.Body.String() != tc.expected {
			t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), tc.expected)
		}
	}
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to Simple Calculator")
	})

	r.GET("/calculate", func(c *gin.Context) {
		aStr := c.Query("a")
		bStr := c.Query("b")
		operation := c.Query("operation")

		// Convert query parameters to integers
		a, _ := strconv.Atoi(aStr)
		b, _ := strconv.Atoi(bStr)

		result := ""

		switch operation {
		case "add":
			result = strconv.Itoa(a + b)
		case "subtract":
			result = strconv.Itoa(a - b)
		case "multiply":
			result = strconv.Itoa(a * b)
		case "divide":
			if b == 0 {
				result = "Error: Division by zero"
			} else {
				result = strconv.Itoa(a / b)
			}
		}

		c.String(http.StatusOK, "Result: %s", result)
	})

	return r
}
