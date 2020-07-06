package ws

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func (r *subscriptionResolver) NotificationAdded(ctx context.Context, id string, userChan *chan User) error {
	var userDoc User
	var change bson.M
	cs, err := r.users.Watch([]bson.M{}, mgo.ChangeStreamOptions{MaxAwaitTimeMS: time.Hour, FullDocument: mgo.FullDocument("updateLookup")})
	if err != nil {
		return err
	}
	if cs.Err() != nil {
		fmt.Println(err)
	}
	go func() {
		start := time.Now()
		for {
			ok := cs.Next(&change)
			if ok {
				byts, _ := bson.Marshal(change["fullDocument"].(bson.M))
				bson.Unmarshal(byts, &userDoc)
				userDoc.ID = bson.ObjectId(userDoc.ID).Hex()
				if userDoc.ID == id {
					*userChan <- userDoc
				}
			}
			if time.Since(start).Minutes() >= 60 {
				break
			}
			continue
		}
	}()
	return nil
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
