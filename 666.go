package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

type Card struct {
	Code   string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Status int64  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"` // 卡状态 -1 未启用 0 已使用 1 未使用 -9 正在处理中
}

func main() {

	cardCodes := []string{"111", "666"}

	cardsSli := make([]*Card, 0)

	for _, cardCode := range cardCodes {
		var cardItem Card
		cardItem.Code = cardCode
		cardItem.Status = 1

		cardsSli = append(cardsSli, &cardItem)
	}

	fmt.Println(cardsSli)

	spew.Dump(cardsSli)

}
