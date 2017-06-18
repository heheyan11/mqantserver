package main

import (
	"github.com/liangdas/mqant"
	"server/chat"
	"server/gate"
	"server/login"
	"server/hitball"
	"server/user"
	"webapp"
	"github.com/liangdas/mqant/module/modules"
	"server/tracing"
	"sourcegraph.com/sourcegraph/appdash"
	appdashtracer "sourcegraph.com/sourcegraph/appdash/opentracing"
	"github.com/opentracing/opentracing-go"
)
var(
	collector *appdash.RemoteCollector= nil

	// Here we use the local collector to create a new opentracing.Tracer
	tracer opentracing.Tracer= nil
)
func DefaultTracer()opentracing.Tracer{
	return tracer
}
//func ChatRoute( app module.App,Type string,hash string) (*module.ServerSession){
//	//演示多个服务路由 默认使用第一个Server
//	log.Debug("Hash:%s 将要调用 type : %s",hash,Type)
//	servers:=app.GetServersByType(Type)
//	if len(servers)==0{
//		return nil
//	}
//	return servers[0]
//}

func main() {
	app := mqant.CreateApp()
	app.DefaultTracer(func()opentracing.Tracer {
		if app.GetSettings().Tracing.Enable{
			if collector==nil{
				collector=appdash.NewRemoteCollector("127.0.0.1:7701")
				tracer=appdashtracer.NewTracer(collector)
			}
			return tracer
		}else{
			return nil
		}
	})
	//app.Route("Chat",ChatRoute)
	app.Run(true, //只有是在调试模式下才会在控制台打印日志, 非调试模式下只在日志文件中输出日志
		modules.MasterModule(),
		hitball.Module(),
		gate.Module(),  //这是默认网关模块,是必须的支持 TCP,websocket,MQTT协议
		login.Module(), //这是用户登录验证模块
		chat.Module(),
		user.Module(),
		webapp.Module(),
		tracing.Module(),
	)  //这是聊天模块

}

