package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// inicializa router com config default
	r := gin.Default()

	// define rotas
	// funções no mesmo package podem ser acessadas sem import
	initializeRoute(r)

	// inicializa router
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
