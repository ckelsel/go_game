package main

import (
	"fmt"
	"gg/utils"
	"gg/xnet"
	"io"
	"net"
	"strconv"
	"time"
)

func main() {
	fmt.Println("client start...")
	utils.Init()

	conn, err := net.Dial("tcp4", "192.168.1.106:"+strconv.Itoa(utils.GlobalObject.Port))
	if err != nil {
		fmt.Println("Dial err", err)
		return
	}

    for {
        time.Sleep(1 * time.Second)

        // msg := xnet.NewXMessage(0, []byte("hello v6"))
        dp := xnet.NewXDataPack()

        // bin, err := dp.Pack(msg)

        // _, err = conn.Write(bin)
        // if err != nil {
        //     fmt.Println("Write err", err)
        // }

        // 第一次读取包头
        buf := make([]byte, dp.GetHeaderLength())
        fmt.Println("ReadFull begin, length ", dp.GetHeaderLength())
        _, err := io.ReadFull(conn, buf)
        if err != nil {
            fmt.Println("ReadFull err ", err)
            return
        }
        fmt.Println("Readfull")

        // 解包读取ID和dataLength
        msgheader, err := dp.UnPack(buf)

        fmt.Println("length ", msgheader.GetLength())

        // 第二次读取数据，长度为dataLength
        if msgheader.GetLength() > 0 {
            data := make([]byte, msgheader.GetLength())
            _, err := io.ReadFull(conn, data)
            if err != nil {
                fmt.Println("ReadFull err ", err)
                return
            }
            msgheader.SetData(data)
        }

        fmt.Println("server replay, MsgID ", msgheader.GetID(), " MsgData ", msgheader.GetData())
    }
}
