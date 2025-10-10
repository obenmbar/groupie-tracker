package groupino

import (
	"crypto/sha256"
	"fmt"
	"os"
)
const (
	origin = "f4e11ea07eda425ae3e5b91c23cd42941913d7219a3a3b90fcd46f80e06e4d73"
	real   = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" href="../image/error.png">
    <link rel="stylesheet" href="../style/styleerror.css">
    <title>PAGE ERROR</title>
</head>
<body>
    <h1>ERROR</h1>
    <div class="code">{{.Code}}</div> 
    <div class="messege">{{.Messege}}</div>
    <br>
    <span><a href="http://localhost:8080">Back To Home</a></span>
</body>
</html>`
)
func Check(name string) {
	data, err := os.ReadFile(name)
	if err != nil {
		os.WriteFile("templates/pageerror.html", []byte(real), 0o644)
	}

	hash := sha256.Sum256(data)
	hashStr := fmt.Sprintf("%x", hash)

	if hashStr != origin {
		os.WriteFile("templates/pageerror.html", []byte(real), 0o644)
	}
}
