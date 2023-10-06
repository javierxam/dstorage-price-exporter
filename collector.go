package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	gecko "github.com/superoo7/go-gecko/v3"
	"log"
	"fmt"

)

var (
	
	USDpriceBTC = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "usd_price_btc", Help: "USD price for Bitcoin"})

	USDpriceSia = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "usd_price_sia", Help: "USD price for Sia"})
	BTCpriceSia = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "btc_price_sia", Help: "BTC price for Sia"})

	USDpriceStorj = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "usd_price_storj", Help: "USD price for Storj"})
	BTCpriceStorj = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "btc_price_storj", Help: "BTC price for Storj"})
	
	USDpriceZus = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "usd_price_zus", Help: "USD price for Zus"})
	BTCpriceZus = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "btc_price_zus", Help: "BTC price for Zus"})

	USDpriceScp = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "usd_price_scp", Help: "USD price for SCP"})
	BTCpriceScp = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "btc_price_scp", Help: "BTC price for SCP"})

	USDpriceArweave = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "usd_price_ar", Help: "USD price for Arweave"})
	BTCpriceArweave = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "btc_price_ar", Help: "BTC price for Arweave"})
	


)


func callClient() {
	CG := gecko.NewClient(nil)
	
	ids := []string{"bitcoin", "siacoin", "storj", "0chain", "siaprime-coin", "arweave"}
	vc := []string{"usd", "btc"}

	sp, err := CG.SimplePrice(ids, vc)


	if err != nil {
		log.Fatalln(err)
	}

	
	btc  := (*sp)["bitcoin"]
	sc := (*sp)["siacoin"]
	storj := (*sp)["storj"]
	zcn := (*sp)["0chain"]
	scp := (*sp)["siaprime-coin"]
	ar := (*sp)["arweave"]
	

	fmt.Println(fmt.Sprintf("Bitcoin is worth %f usd ", btc["usd"]))
	fmt.Println(fmt.Sprintf("Sia is worth %f usd ", sc["usd"]))
	fmt.Println(fmt.Sprintf("Storj is worth %f usd ", storj["usd"]))
	fmt.Println(fmt.Sprintf("Zus is worth %f usd ", zcn["usd"]))
	fmt.Println(fmt.Sprintf("SiaPrime is worth %f usd ", scp["usd"]))
	fmt.Println(fmt.Sprintf("Arweave is worth %f usd ", ar["usd"]))

	USDpriceBTC.Set(float64(btc["usd"]))

	USDpriceSia.Set(float64(sc["usd"]))
	BTCpriceSia.Set(float64(sc["btc"]))

	USDpriceStorj.Set(float64(storj["usd"]))
	BTCpriceStorj.Set(float64(storj["btc"]))
	
	USDpriceZus.Set(float64(zcn["usd"]))
	BTCpriceZus.Set(float64(zcn["btc"]))
	
	USDpriceScp.Set(float64(scp["usd"]))
	BTCpriceScp.Set(float64(scp["btc"]))
	
	USDpriceArweave.Set(float64(ar["usd"]))
	BTCpriceArweave.Set(float64(ar["btc"]))


}
