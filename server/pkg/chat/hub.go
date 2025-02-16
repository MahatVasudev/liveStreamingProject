package chat

type Hub struct {
	Clients    map[*Client]bool
	Broadcasts chan []byte
	Register   chan *Client
	UnRegister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Broadcasts: make(chan []byte),
		Register:   make(chan *Client),
		UnRegister: make(chan *Client),
		Clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
		case client := <-h.UnRegister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}

		case message := <-h.Broadcasts:
			for client := range h.Clients {
				select {
				case client.Send <- message:

				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}
		}
	}
}
