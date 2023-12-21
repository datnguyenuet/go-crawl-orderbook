package websocket_client

import (
	"time"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type WebsocketClient struct {
	logger  *zap.SugaredLogger
	URL     string
	conn    *websocket.Conn
	running bool
}

func NewWebsocketClient(url string) *WebsocketClient {
	return &WebsocketClient{
		URL: url,
	}
}

func (c *WebsocketClient) Start() {
	c.running = true

	for c.running {
		if err := c.connect(); err != nil {
			c.logger.Error("Error connecting to websocket: %v. Retrying in 5 seconds...", err)
			time.Sleep(5 * time.Second)
		}
	}

	c.logger.Debug("WebSocket client stopped.")
}

func (c *WebsocketClient) Close() {
	c.running = false
	if c.conn != nil {
		if err := c.conn.Close(); err != nil {
			return
		}
	}
}

func (c *WebsocketClient) connect() error {
	conn, _, err := websocket.DefaultDialer.Dial(c.URL, nil)
	if err != nil {
		return err
	}
	c.conn = conn

	return nil
}
