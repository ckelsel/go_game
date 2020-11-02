package xnet

import (
	"fmt"
	"io"
	"net"
	"testing"
	"time"
	"xkm/utils"
)

// go test xkm/xnet -v
// 封包、拆包的单元测试
func TestXDataPack(t *testing.T) {
	utils.Init()

	// 模拟服务端
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		t.Error("Listen err ", err)
		return
	}

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				t.Error("Accept err ", err)
				break
			}

			go func(conn net.Conn) {
				dp := NewXDataPack()

				for {
					msgHeader := make([]byte, dp.GetHeaderLength())

					// header
					_, err := io.ReadFull(conn, msgHeader)
					if err != nil {
						t.Error("ReadFull err", err)
						break
					}

					unPackMsg, err := dp.UnPack(msgHeader)
					if err != nil {
						t.Error("Unpack err", err)
						break
					}

					if unPackMsg.GetLength() == 0 {
						fmt.Println("msg1 without data")
						continue
					}

					msg1 := unPackMsg.(*XMessage)

					// data
					msg1.Data = make([]byte, msg1.GetLength())
					_, err = io.ReadFull(conn, msg1.Data)
					if err != nil {
						t.Error("ReadFull err", err)
						break
					}

					fmt.Println("Recv msg1, ID ", msg1.GetID(), ", Length ", msg1.GetLength(), ", Data ", msg1.GetData())
				}

			}(conn)
		}
	}()

	// 模拟客户端
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		t.Error("Dial err ", err)
		return
	}

	msg1 := &XMessage{}

	msg1.SetID(1)
	msg1.SetLength(4)
	msg1.SetData([]byte("1234"))

	dp := NewXDataPack()

	packMsg1, err := dp.Pack(msg1)
	if err != nil {
		t.Error("Pack err ", err)
	}

	msg2 := &XMessage{}

	msg2.SetID(2)
	msg2.SetLength(2)
	msg2.SetData([]byte("ab"))

	packMsg2, err := dp.Pack(msg2)
	if err != nil {
		t.Error("Pack err ", err)
	}

	// 黏包
	packMag := append(packMsg1, packMsg2...)
	_, err = conn.Write(packMag)
	if err != nil {
		t.Error("Write err ", err)
	}

	// 延时等待
	time.Sleep(1 * time.Second)
}
