package main

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type SinoPacOrderStatus struct {
	Action   string  `json:"action"`
	Code     string  `json:"code"`
	ID       string  `json:"id"`
	Price    float64 `json:"price"`
	Quantity int64   `json:"quantity"`
	Status   string  `json:"status"`
}

type VolumeClose struct {
	Code   string  `json:"code"`
	Close  float64 `json:"close"`
	Volume int64   `json:"volume"`
}

type UpdateVolumeArrBody struct {
	StockNumArr []string `json:"stock_num_arr"`
}

func main() {
	var res []VolumeClose
	a := "asfdasfsafds"
	arr := append([]string{}, "2330", "2609")
	tmp := UpdateVolumeArrBody{
		StockNumArr: arr,
	}
	client := resty.New()
	resp, err := client.R().
		SetBody(&tmp).SetResult(&a).
		Post("http://tradebot.tocandraw.com:3333/basic/volumeclos")
	if err != nil {
		fmt.Println(err)
	}
	// if err := json.Unmarshal(resp.Body(), &res); err != nil {
	// 	panic(err)
	// }
	b := resp.Result().(*string)
	fmt.Println(resp.StatusCode(), *b)
	fmt.Println(res)
}
