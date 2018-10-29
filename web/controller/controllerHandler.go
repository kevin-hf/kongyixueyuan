/**
  author: kevin
 */
package controller

import (
	"net/http"
	"github.com/kongyixueyuan.com/kongyixueyuan/service"
)

type Application struct {
	Fabric *service.ServiceSetup
}

func (app *Application) IndexView(w http.ResponseWriter, r *http.Request){
	showView(w, r, "index.html", nil)
}

func (app *Application) SetInfoView(w http.ResponseWriter, r *http.Request)  {
	showView(w, r, "setInfo.html", nil)
}

// 根据指定的 key 设置/修改 value 信息
func (app *Application) SetInfo(w http.ResponseWriter, r *http.Request)  {
	// 获取提交数据
	name := r.FormValue("name")
	num := r.FormValue("num")

	// 调用业务层, 反序列化
	transactionID, err := app.Fabric.SetInfo(name, num)

	// 封装响应数据
	data := &struct {
		Flag bool
		Msg string
	}{
		Flag:true,
		Msg:"",
	}
	if err != nil {
		data.Msg = err.Error()
	}else{
		data.Msg = "操作成功，交易ID: " + transactionID
	}

	// 响应客户端
	showView(w, r, "setInfo.html", data)
}

// 根据指定的 Key 查询信息
func (app *Application) QueryInfo(w http.ResponseWriter, r *http.Request)  {
	// 获取提交数据
	name := r.FormValue("name")

	// 调用业务层, 反序列化
	msg, err := app.Fabric.GetInfo(name)

	// 封装响应数据
	data := &struct {
		Msg string
	}{
		Msg:"",
	}
	if err != nil {
		data.Msg = "没有查询到Jack对应的信息"
	}else{
		data.Msg = "查询成功: " + msg
	}
	// 响应客户端
	showView(w, r, "queryReq.html", data)
}
