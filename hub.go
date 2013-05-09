package main

type hub struct {
  // Registered connections.
  connections map[*connection]bool

  // Inbound messages from the connections.
  broadcast chan string

  // Register requests from the connections.
  register chan *connection

  // Unregister requests from connections.
  unregister chan *connection
}

func (h *hub) run() {
  for {
    select {
    case c := <-h.register:
      h.connections[c] = true
    case c := <-h.unregister:
      delete(h.connections, c)
      close(c.send)
    case m := <-h.broadcast:
      for c := range h.connections {
        select {
        case c.send <- m:
        default:
          delete(h.connections, c)
          close(c.send)
          go c.ws.Close()
        }
      }
    }
  }
}

