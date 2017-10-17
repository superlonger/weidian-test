// todo.go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"os"
)

type MsgBuyerInfo struct {
	BuyerId string `json:"buyer_id"`
	Phone   string `json:"phone"`
	Name    string `json:"name"`
}
type MsgItem struct {
	ItemId     string `json:"item_id"`
	ItemName   string `json:"item_name"`
	Price      string `json:"price"`
	Quantity   string `json:"quantity"`
	TotalPrice string `json:"total_price"`
}
type MsgMessage struct {
	SellerName    string       `json:"seller_name"`
	Total         string       `json:"total"`
	OrderId       string       `json:"order_id"`
	OrderTypeDesc string       `json:"order_type_des"`
	Quantity      string       `json:"quantity"`
	Note          string       `json:"note"`
	Status        string       `json:"status"`
	BuyerInfo     MsgBuyerInfo `json:"buyer_info"`
	Items         []MsgItem    `json:"items"`
}
type MsgPayment struct {
	Message MsgMessage `json:"message"`
}

type MsgRecv struct {
	Type string `json:"type"`
}

//Message MsgMessage `json:"message"`
func main() {
	test()
	return
	// setup http log file
	logHttpFile, err := os.OpenFile("log/http.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		panic(err)
	}
	defer logHttpFile.Close()
	// create a new instance of Echo
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{${time_rfc3339}: "ip":"${remote_ip}","host":"${host}",` +
			`"method":"${method}","uri":"${uri}","status":${status},` +
			`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
			`"bytes_out":${bytes_out}}` + "\n",
		Output: logHttpFile,
	}))
	e.Static("/", "public")
	e.POST("/", recvMsg)
	e.Logger.Fatal(e.Start(":8000"))
}
func recvMsg(c echo.Context) error {
	var msgRecv MsgRecv
	content := c.FormValue("content")
	fmt.Println(content)
	err := json.Unmarshal([]byte(content), &msgRecv)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, `{"status":"success"}`)
}
func test() {
	//response := `{"message":{"seller_name":"包子铺","total":"1.00","express_fee":"0.00","buyer_identity_id":null,"f_phone":"","fx_fee_value":"","express_type":null,"express_note":"","modify_price_enable":"0","express":"","order_id":"800282104751141","order_type_des":"担保交易","confirm_expire":"","group_status":-1,"send_time":"","weixin":"","is_close":0,"argue_flag":0,"total_fee":"","status_desc":"待发货","quantity":"1","discount_list":[],"wei_supplier_price":"","note":"1213141516","sk":null,"trade_no":null,"status_ori":"20","status":"pay","seller_phone":"14716468888","express_fee_num":"0.00","refund_status_ori":"0","pay_time":"2017-10-15 02:50:53","is_wei_order":0,"real_income_price":null,"price":"1.00","buyer_info":{"buyer_id":"1204747211","idCardNo":"","region":"岚皋县","phone":"13891502008","post":null,"address":"陕西 安康市 岚皋县 城关镇 滔河路8号","name":"袁方","self_address":"城关镇滔河路8号","province":"陕西","city":"安康"},"items":[{"is_delivered":0,"img":"http://wd.geilicdn.com/bj-vshop-1259798741-1506798935233-1986573863_120_120.jpg?w=110&h=110&cp=1","warehouseId":null,"total_price":"1.00","merchant_code":"","deliver_id":"0","url":"https://weidian.com/i/2175400745","deliver_status_desc":"","price":"1.00","sku_merchant_code":null,"item_name":"test 1","item_id":"2175400745","fx_fee_rate":"","quantity":"1","can_deliver":1,"refund_info":{"refund_no":"","refund_fee":"","refund_status_str":"","item_id":"2175400745","item_sku_id":"0","refund_status":"","refund_kind":"0","refund_express_fee":"","can_refund":"1","refund_item_fee":"","refund_status_desc":""},"sku_id":"0","sku_title":""}],"user_phone":"14716466677","status2":null,"f_seller_id":"0","seller_id":"1259798741","add_time":"2017-10-15 02:48:27","refund_info":{"buyer_refund_fee":null,"refund_time":null},"return_code":null,"express_no":"","f_shop_name":"","original_total_price":"","order_type":"3"},"type":"weidian.order.already_payment"}`
	response := `{"message":{"seller_name":"包子铺","total":"1.00","express_fee":"0.00","buyer_identity_id":null,"f_phone":"","fx_fee_value":"","express_type":null,"express_note":"","modify_price_enable":"1","express":"","order_id":"800282040181797","order_type_des":"担保交易","confirm_expire":"","group_status":-1,"send_time":"","weixin":"","is_close":0,"argue_flag":0,"total_fee":"","status_desc":"待付款","quantity":"1","discount_list":[],"wei_supplier_price":"","note":"","sk":null,"trade_no":null,"status_ori":"10","status":"unpay","seller_phone":"14716468888","express_fee_num":"0.00","refund_status_ori":"0","pay_time":"","is_wei_order":0,"real_income_price":null,"price":"1.00","buyer_info":{"buyer_id":"1204747211","idCardNo":"","region":"岚皋县","phone":"13891502008","post":null,"address":"陕西 安康市 岚皋县 城关镇 滔河路8号","name":"袁方","self_address":"城关镇滔河路8号","province":"陕西","city":"安康"},"items":[{"is_delivered":0,"img":"http://wd.geilicdn.com/bj-vshop-1259798741-1506798935233-1986573863_120_120.jpg?w=110&h=110&cp=1","warehouseId":null,"total_price":"1.00","merchant_code":"","deliver_id":"0","url":"https://weidian.com/i/2175400745","deliver_status_desc":"","price":"1.00","sku_merchant_code":null,"item_name":"test 1","item_id":"2175400745","fx_fee_rate":"","quantity":"1","can_deliver":0,"refund_info":{"refund_no":"","refund_fee":"","refund_status_str":"","item_id":"2175400745","item_sku_id":"0","refund_status":"","refund_kind":"0","refund_express_fee":"","can_refund":"0","refund_item_fee":"","refund_status_desc":""},"sku_id":"0","sku_title":""}],"user_phone":"14716466677","status2":null,"f_seller_id":"0","seller_id":"1259798741","add_time":"2017-10-15 00:33:28","refund_info":{"buyer_refund_fee":null,"refund_time":null},"return_code":null,"express_no":"","f_shop_name":"","original_total_price":"","order_type":"3"},"type":"weidian.order.non_payment"}`
	var msgRecv MsgRecv
	err := json.Unmarshal([]byte(response), &msgRecv)
	if err != nil {
		return
	}
	fmt.Printf("%+v\n", msgRecv)
	switch msgRecv.Type {
	case "weidian.order.already_payment":
		{
			// 已付款
			var msgPayment MsgPayment
			err = json.Unmarshal([]byte(response), &msgPayment)
			if err != nil {
				return
			}
			fmt.Printf("%+v\n", msgPayment)
		}
	case "weidian.order.non_payment":
		{
			// 未付款
			var msgPayment MsgPayment
			err = json.Unmarshal([]byte(response), &msgPayment)
			if err != nil {
				return
			}
			fmt.Printf("%+v\n", msgPayment)
		}
	default:
		fmt.Println("unknown type")
	}
	return

}
