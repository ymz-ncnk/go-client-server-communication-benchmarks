package nhj

import (
	"bufio"
	"net"
)

type bufferedConn struct {
	r *bufio.Reader
	w *bufio.Writer
	net.Conn
}

func (c *bufferedConn) Read(b []byte) (int, error) {
	return c.r.Read(b)
}

func (c *bufferedConn) Write(b []byte) (n int, err error) {
	n, err = c.w.Write(b)
	if err != nil {
		return
	}
	err = c.w.Flush()
	return
}

func (c *bufferedConn) Close() (err error) {
	err = c.w.Flush()
	if err != nil {
		return
	}
	return c.Conn.Close()
}
