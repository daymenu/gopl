package main

import (
	"errors"
	"io/ioutil"
	"net"
)

type Client struct {
	conn net.Conn
	buf  [8192]byte
}

func (c *Client) readPackage() (data []byte, err error) {
	n, err := c.conn.Read(c.buf[0:4])
	if n != 4 {
		err = errors.New("read header falied")
		return
	}
	return ioutil.ReadAll(c.conn)
}
