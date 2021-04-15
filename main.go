package main

import (
	"ginqb/initRouter"
)

func main() {
	router := initRouter.SetupRouter()
	_ = router.Run()
}
