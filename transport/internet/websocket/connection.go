package websocket

import (
	"io"
	"net"
	"time"

	"github.com/gorilla/websocket"
	"v2ray.com/core/common/errors"
)

// connection is a wrapper for net.Conn over WebSocket connection.
type connection struct {
	wsc    *websocket.Conn
	reader io.Reader
}

// Read implements net.Conn.Read()
func (c *connection) Read(b []byte) (int, error) {
	for {
		reader, err := c.getReader()
		if err != nil {
			return 0, err
		}

		nBytes, err := reader.Read(b)
		if errors.Cause(err) == io.EOF {
			c.reader = nil
			continue
		}
		return nBytes, err
	}
}

func (c *connection) getReader() (io.Reader, error) {
	if c.reader != nil {
		return c.reader, nil
	}

	_, reader, err := c.wsc.NextReader()
	if err != nil {
		return nil, err
	}
	c.reader = reader
	return reader, nil
}

func (c *connection) Write(b []byte) (int, error) {
	if err := c.wsc.WriteMessage(websocket.BinaryMessage, b); err != nil {
		return 0, err
	}
	return len(b), nil
}

func (c *connection) Close() error {
	c.wsc.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""), time.Now().Add(time.Second*5))
	return c.wsc.Close()
}

func (c *connection) LocalAddr() net.Addr {
	return c.wsc.LocalAddr()
}

func (c *connection) RemoteAddr() net.Addr {
	return c.wsc.RemoteAddr()
}

func (c *connection) SetDeadline(t time.Time) error {
	if err := c.SetReadDeadline(t); err != nil {
		return err
	}
	return c.SetWriteDeadline(t)
}

func (c *connection) SetReadDeadline(t time.Time) error {
	return c.wsc.SetReadDeadline(t)
}

func (c *connection) SetWriteDeadline(t time.Time) error {
	return c.wsc.SetWriteDeadline(t)
}
