package routers

// パッケージのインポート
import (
	"github.com/gin-gonic/gin"
)

func InitHelloWorld(r *gin.RouterGroup) {
	r.GET("/", helloWorld)
}


func helloWorld(ctx *gin.Context) {
		ctx.String(200, "Hello World 8")
}