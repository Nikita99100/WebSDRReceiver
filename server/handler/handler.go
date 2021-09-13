package handler
import(
	"github.com/Nikita99100/WebSDRReceiver/server/proto"	
)
type Handler struct {
	Receiver *proto.ReceiverSettings
}

func (h *Handler) SetFreq(freq float32) error {
	h.Receiver.Freq=freq
	return nil
}
func (h *Handler) GetReceiverSettings() (*proto.ReceiverSettings, error) {
	return h.Receiver, nil
}
