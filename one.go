package main
 
import (
    "github.com/gin-gonic/gin"
    "html/template"
    "net/http"
    "strconv"
)
 
func main() {
    r := gin.Default()
 
    r.SetHTMLTemplate(template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<title>Simple Calculator</title>
</head>
<body>
<h1>Simple Calculator</h1>
<form action="/calculate" method="GET">
<label for="a">Number A:</label>
<input type="number" id="a" name="a"><br><br>
<label for="b">Number B:</label>
<input type="number" id="b" name="b"><br><br>
<button type="submit" name="operation" value="add">Addition</button>
<button type="submit" name="operation" value="subtract">Subtraction</button>
<button type="submit" name="operation" value="multiply">Multiplication</button>
<button type="submit" name="operation" value="divide">Division</button>
</form>
<br>
            {{if .Result}}
<h2>Result: {{.Result}}</h2>
            {{end}}
</body>
</html>
    `)))
 
    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "", nil)
    })
 
    r.GET("/calculate", func(c *gin.Context) {
        a, _ := strconv.Atoi(c.Query("a"))
        b, _ := strconv.Atoi(c.Query("b"))
        operation := c.Query("operation")
        var result string
 
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
 
        c.HTML(http.StatusOK, "", gin.H{"Result": result})
    })
 
    // Start the server
    port := "8888"
    r.Run(":" + port)
}