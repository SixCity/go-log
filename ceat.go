package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

type AliaRemoteReaStrict struct {
	Ono         string   `json:"ono"`
	OrderItem   []Item   `json:"order_item"`
	OrderRefund []Refund `json:"order_refund"`
}
type Item struct {
	Ono string `json:"ono"`
	Oid int    `json:"oid"`
}
type Refund struct {
	Ono     string `json:"ono"`
	Item    int    `json:"item"`
	Content string `json:"content"`
	Imps    string `json:"imps"`
	Status  string `json:"status"`
}

func Celt() {
	var m AliaRemoteReaStrict
	m.Ono = "12345"
	m.OrderItem = append(m.OrderItem, Item{Ono: "Shanghai_VPN", Oid: 1})
	m.OrderItem = append(m.OrderItem, Item{Ono: "Beijing_VPN", Oid: 2})
	for i := 1; i < 6; i++ {
		str := []byte("物品")
		str = strconv.AppendInt(str, int64(i), 10)
		orders := Item{Ono: string(str), Oid: i}
		m.OrderItem = append(m.OrderItem, orders)
	}
	bytes, _ := json.Marshal(m)
	fmt.Printf("json:m,%s\n", bytes)

	var js AliaRemoteReaStrict

	err := json.Unmarshal(bytes, &js)

	fmt.Printf("json:type,%s\n", reflect.TypeOf(bytes))

	if err != nil {
		fmt.Printf("format err:%s\n", err.Error())
	}
}
