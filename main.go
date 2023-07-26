package main

import "cos_practice/router"

func main() {
	r := router.Router()
	r.Run("127.0.0.1:9003")
}
