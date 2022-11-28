package routers

// パッケージのインポート
import (
	"github.com/gin-gonic/gin"
)

func Router(e *gin.Engine) {
	api := e.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			InitHelloWorld(v1)
		}
	}
}
