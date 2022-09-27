package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Obtener la arquitectura y el sistema operativo
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
