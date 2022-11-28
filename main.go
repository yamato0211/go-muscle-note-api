// パッケージ名
package main

// パッケージのインポート
import (
	"log"
	"test/routers"

	"github.com/gin-gonic/gin"
)

// go run main.go を実行時に最初に呼ばれる
func main() {
	e := gin.Default()
	routers.Router(e)

	if err := e.Run("0.0.0.0:8000"); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}
}
