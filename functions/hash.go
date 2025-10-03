package groupino

import (
	"crypto/sha256"
	"fmt"
	"os"
)
const (
	origin = "aaa7ae8ca3bca72c2642ecc577e2a8a8bee4e7492b1c8f967f6262e7437f32be"
	real   = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    
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
