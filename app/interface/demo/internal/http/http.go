package server

import (
	"log"
	xgin "miopkg/http/gin"
	"github.com/gin-gonic/gin"
)


func (eng *Engine) NewHttpServer() error {
	server := xgin.StdConfig("http").Build()
	server.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, "Hello Gin")
	})
	da := *NewDemoInterface()
	apiv1 := server.Group("/v1")
	{
		// pa := ada.NewProductAdapter()
		apiv1.POST("/demo", da.CreateDemo)
	}
	//Upgrade to websocket
	server.Upgrade(xgin.WebSocketOptions("/ws", func(ws xgin.WebSocketConn, err error) {
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		for {
			mt, message, err := ws.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", message)
			err = ws.WriteMessage(mt, message)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	}))
	return eng.Serve(server)
}
