package server
import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/Nikita99100/WebSDRReceiver/server/proto"
)
func (r *Rest) GetReceiverSettings(c echo.Context) error {
	var freq proto.ReceiverSettings
	
	r.Handler.SetFreq(
	return nil
}
func (r *Rest) SetSettings(c echo.Context) error {
	return nil
}
func (r *Rest) SetMode(c echo.Context) error {
	return nil
}
func (r *Rest) SetFreq(c echo.Context) error {
	return nil
}
