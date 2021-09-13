package server
import (
	"github.com/pkg/errors"
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/Nikita99100/WebSDRReceiver/server/proto"
)
func (r *Rest) GetReceiverSettings(c echo.Context) error {
	out, err := r.Handler.GetReceiverSettings()
		if err != nil {
			res := errors.Wrap(err, "can't get receiver settings")
			return c.JSON(http.StatusBadRequest, res)
		}
	return c.JSON(http.StatusOK, out)
}
func (r *Rest) SetSettings(c echo.Context) error {
	return nil
}
func (r *Rest) SetMode(c echo.Context) error {
	return nil
}
func (r *Rest) SetFreq(c echo.Context) error {
	var freq proto.ReceiverSettings
	if err := c.Bind(&freq); err != nil {
		res := errors.Wrap(err, "failed to parse the request")
		return c.JSON(http.StatusBadRequest, res)
	}
	err := r.Handler.SetFreq(freq.Freq)
	return c.JSON(http.StatusBadRequest,err)
}
