package room

import (
	"github.com/MahatVasudev/liveStreamingProject/server/payloads"
	w "github.com/MahatVasudev/liveStreamingProject/server/pkg/webrtc"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
)

func (h *Handler) Welcome(c *fiber.Ctx) error {

	return c.JSON(map[string]string{"data": "Welcome to our room services"})
}

func (h *Handler) CreateRoom(c *fiber.Ctx) error {

	p := payloads.ResponseCreate{
		ID:   uuid.NewString(),
		Data: "created new room",
	}

	return c.JSON(p)
}

func (h *Handler) RoomWebsocket(c *websocket.Conn) {
	uuid := c.Params("uuid")

	if uuid == "" {
		return
	}
	_, _, room := createOrGetRoom(uuid)
	w.RoomConn(c, room.Peers)
}

func createOrGetRoom(guuid string) (string, string, w.Rooms) {
	panic("unimplemented")
}

func roomViewerConn(c *websocket.Conn, p *w.Peers) {

}
