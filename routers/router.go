package routers

import (
	"PrometheusAlert/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/alerttest", &controllers.MainController{},"post:AlertTest")
	//prometheus
	beego.Router("/prometheus/alert", &controllers.PrometheusController{},"post:PrometheusAlert")
	beego.Router("/prometheus/router", &controllers.PrometheusController{},"post:PrometheusRouter")
	//beego.Router("/prometheus/dingding", &controllers.PrometheusController{},"post:PrometheusRouter")
	//beego.Router("/prometheus/weixin", &controllers.PrometheusController{},"post:PrometheusRouter")
	//beego.Router("/prometheus/feishu", &controllers.PrometheusController{},"post:PrometheusRouter")
	//beego.Router("/prometheus/txdx", &controllers.PrometheusController{},"post:PrometheusRouter")
	//beego.Router("/prometheus/txdh", &controllers.PrometheusController{},"post:PrometheusRouter")
	//beego.Router("/prometheus/hwdx", &controllers.PrometheusController{},"post:PrometheusRouter")
	//beego.Router("/prometheus/alydx", &controllers.PrometheusController{},"post:PrometheusRouter")
	//beego.Router("/prometheus/alydh", &controllers.PrometheusController{},"post:PrometheusRouter")
	//beego.Router("/prometheus/rlydh", &controllers.PrometheusController{},"post:PrometheusRouter")
    //graylog2
	//beego.Router("/graylog2/phone", &controllers.Graylog2Controller{},"post:GraylogTxdh")
	beego.Router("/graylog2/dingding", &controllers.Graylog2Controller{},"post:GraylogDingding")
	beego.Router("/graylog2/weixin", &controllers.Graylog2Controller{},"post:GraylogWeixin")
	beego.Router("/graylog2/feishu", &controllers.Graylog2Controller{},"post:GraylogFeishu")
	beego.Router("/graylog2/txdx", &controllers.Graylog2Controller{},"post:GraylogTxdx")
	beego.Router("/graylog2/txdh", &controllers.Graylog2Controller{},"post:GraylogTxdh")
	beego.Router("/graylog2/hwdx", &controllers.Graylog2Controller{},"post:GraylogHwdx")
	beego.Router("/graylog2/alydx", &controllers.Graylog2Controller{},"post:GraylogALYdx")
	beego.Router("/graylog2/alydh", &controllers.Graylog2Controller{},"post:GraylogALYdh")
	beego.Router("/graylog2/rlydh", &controllers.Graylog2Controller{},"post:GraylogRLYdh")
	//graylog3
	//beego.Router("/graylog3/phone", &controllers.Graylog3Controller{},"post:GraylogTxdh")
	beego.Router("/graylog3/dingding", &controllers.Graylog3Controller{},"post:GraylogDingding")
	beego.Router("/graylog3/weixin", &controllers.Graylog3Controller{},"post:GraylogWeixin")
	beego.Router("/graylog3/feishu", &controllers.Graylog3Controller{},"post:GraylogFeishu")
	beego.Router("/graylog3/txdx", &controllers.Graylog3Controller{},"post:GraylogTxdx")
	beego.Router("/graylog3/txdh", &controllers.Graylog3Controller{},"post:GraylogTxdh")
	beego.Router("/graylog3/hwdx", &controllers.Graylog3Controller{},"post:GraylogHwdx")
	beego.Router("/graylog3/alydx", &controllers.Graylog3Controller{},"post:GraylogALYdx")
	beego.Router("/graylog3/alydh", &controllers.Graylog3Controller{},"post:GraylogALYdh")
	beego.Router("/graylog3/rlydh", &controllers.Graylog3Controller{},"post:GraylogRLYdh")
    //grafana
	//beego.Router("/grafana/phone", &controllers.GrafanaController{},"post:GrafanaTxdh")
	beego.Router("/grafana/dingding", &controllers.GrafanaController{},"post:GrafanaDingding")
	beego.Router("/grafana/weixin", &controllers.GrafanaController{},"post:GrafanaWeixin")
	beego.Router("/grafana/feishu", &controllers.GrafanaController{},"post:GrafanaFeishu")
	beego.Router("/grafana/txdx", &controllers.GrafanaController{},"post:GrafanaTxdx")
	beego.Router("/grafana/txdh", &controllers.GrafanaController{},"post:GrafanaTxdh")
	beego.Router("/grafana/hwdx", &controllers.GrafanaController{},"post:GrafanaHwdx")
	beego.Router("/grafana/alydx", &controllers.GrafanaController{},"post:GrafanaALYdx")
	beego.Router("/grafana/alydh", &controllers.GrafanaController{},"post:GrafanaALYdh")
	beego.Router("/grafana/rlydh", &controllers.GrafanaController{},"post:GrafanaRlydh")
	beego.Router("/tengxun/status", &controllers.TengXunStatusController{},"post:TengXunStatus")
    //zabbix
	beego.Router("/zabbix/alert", &controllers.ZabbixController{},"post:ZabbixAlert")
}
