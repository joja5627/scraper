package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := SetupRouter()
	router.Run()
}
