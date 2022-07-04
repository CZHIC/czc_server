package main

import (
	"czc-server/app/apitcp"
	"czc-server/app/apiudp"
	_ "czc-server/boot"
	"czc-server/library/logger"
	"czc-server/library/pubFuc"
	_ "czc-server/router"
	"time"

	"github.com/go-echarts/statsview"
	"github.com/go-echarts/statsview/viewer"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/v2/net/gtcp"
	"github.com/gogf/gf/v2/net/gudp"
)

func main() {
	if g.Cfg().GetString("GCRON.GetRtx") != "pro" {
		go func() {
			//增加一个系统监控
			mgr := statsview.New()
			viewer.SetConfiguration(viewer.WithTheme(viewer.ThemeMacarons), viewer.WithMaxPoints(1000))
			mgr.Start()
			//访问。 http://localhost:18066/debug/statsview
			// busy working....
			time.Sleep(time.Minute)
		}()
	}
	go g.Server().Run() // http

	//udp
	go gudp.NewServer("127.0.0.1:8899", func(conn *gudp.Conn) {
		logger.Println("服务启动：", conn.LocalAddr(), conn.RemoteAddr())
		defer conn.Close()
		for {
			msg, err := pubFuc.UdpRecvPkg(conn)
			if err == nil {
				if err := pubFuc.UdpSendPkg(conn, msg); err != nil {
					g.Log().Error(err)
				}
			}
			if err != nil {
				g.Log().Error(err)
			}
			apiudp.Router(conn, msg)
		}
	}).Run()

	// tcp
	gtcp.NewServer("127.0.0.1:8999", func(conn *gtcp.Conn) {
		defer conn.Close()

		for {
			msg, err := pubFuc.RecvPkg(conn)
			if err != nil {
				if err.Error() == "EOF" {
					glog.Println("client closed")
				}
				break
			}
			apitcp.Router(conn, msg)
		}

	}).Run()

}
