package net

import (
	"io"
	"net"
)

func Serve(network, addr string) error {
	listener, err := net.Listen(network, addr)
	if err != nil {
		// 比较常见的是端口被占用
		return err
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}

		go func() {
			if err := handleConn(conn); err != nil {
				_ = conn.Close()
			}
		}()
	}
}

func handleConn(conn net.Conn) error {
	bs := make([]byte, 8)
	_, err := conn.Read(bs)
	if err == io.EOF || err == io.ErrUnexpectedEOF || err == net.ErrClosed {
		return err
	}

	// 这种是可以挽救的
	// 但通常建议是遇到错误就把连接关掉
	if err != nil {
		return err
	}

	//if n != 8 {
	//	return errors.New("micro: 没读够数据")
	//}

	res := handleMsg(bs)
	_, err = conn.Write(res)
	if err == io.EOF || err == io.ErrUnexpectedEOF || err == net.ErrClosed {
		return err
	}

	// 这种是可以挽救的
	// 但通常建议是遇到错误就把连接关掉
	if err != nil {
		return err
	}

	//if n != 8 {
	//	return errors.New("micro: 没写完数据")
	//}

	return nil
}

func handleMsg(bs []byte) []byte {
	return nil
}
