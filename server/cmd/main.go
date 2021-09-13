package main
import (
	//"log"
	"time"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/Nikita99100/WebSDRReceiver/server/os"
	"github.com/Nikita99100/WebSDRReceiver/server/web"
	"github.com/Nikita99100/WebSDRReceiver/server/handler"
	"github.com/Nikita99100/WebSDRReceiver/server/server"
	"github.com/Nikita99100/WebSDRReceiver/server/proto"
)
const(
		serverPort = "80"
		serverShutdownTimeout = 30 * time.Second
)
func main(){
	
	hdlr := handler.Handler{
		Receiver:&proto.ReceiverSettings{},
	}

	router := web.NewRouter(
		cors,
	)
	rest := server.Rest{
		Router:  router,
		Handler: &hdlr,
	}
	rest.Route()
	go web.Start(router, serverPort)
	defer web.Stop(router, serverShutdownTimeout)
	<-os.NotifyAboutExit()
}

func cors(e *echo.Echo) {
	e.Use(middleware.CORS()) // allow CORS
}