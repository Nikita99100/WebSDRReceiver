package main
import (
	"log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/Nikita99100/WebSDRReceiver/server/os"
	"github.com/Nikita99100/WebSDRReceiver/server/web"
	"github.com/Nikita99100/WebSDRReceiver/server/server"

)
const(
		serverPort = 80
		serverShutdownTimeout = 30 * time.Second
)
func main(){
	hdlr := handler.Handler{}

	router := web.NewRouter(
		e.Use(middleware.CORS()) // allow CORS
	)
	rest := server.Rest{
		Router:  router,
		Handler: &hdlr,
	}
	go web.Start(router, serverPort)x
	defer web.Stop(router, serverShutdownTimeout)
	<-os.NotifyAboutExit()
}