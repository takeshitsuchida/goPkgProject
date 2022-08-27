package main

/* btcusd
import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

type JsonRPC2 struct {
	Version string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Result  interface{} `json:"result,omitempty"`
	Id      *int        `json:"id,omitempty"`
}
type SubscribeParams struct {
	Channel string `json:"channel"`
}

func main() {
	u := url.URL{Scheme: "wss", Host: "ws.lightstream.bitflyer.com", Path: "/json-rpc"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	if err := c.WriteJSON(&JsonRPC2{Version: "2.0", Method: "subscribe", Params: &SubscribeParams{"lightning_ticker_BTC_USD"}}); err != nil {
		log.Fatal("subscribe:", err)
		return
	}

	for {
		message := new(JsonRPC2)
		if err := c.ReadJSON(message); err != nil {
			log.Println("read:", err)
			return
		}

		if message.Method == "channelMessage" {
			log.Println(message.Params)
		}
	}
}
*/
/* talib
func main() {
	spy, _ := quote.NewQuoteFromYahoo("spy", "2017-01-01", "2017-07-12", quote.Daily, true)
	fmt.Print(spy.CSV())
	rsi2 := talib.Rsi(spy.Close, 2)
	fmt.Println(rsi2)
}
/*
/* ini
type ConfigList struct {
	Port      int
	DbName    string
	SQLDriver string
}

var Config ConfigList

func init() {
	cfg, _ := ini.Load("config.ini")
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustInt(),
		DbName:    cfg.Section("db").Key("name").MustString("example.sql"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
	}
}

func main() {
	fmt.Printf("%T %v\n", Config.Port, Config.Port)
	fmt.Printf("%T %v\n", Config.DbName, Config.DbName)
	fmt.Printf("%T %v\n", Config.SQLDriver, Config.SQLDriver)
}
*/
/* semaphore
var s *semaphore.Weighted = semaphore.NewWeighted(1)

func longPrecess(ctx context.Context) {
	if err := s.Acquire(ctx, 1); err != nil {
		fmt.Println(err)
		return
	}
	defer s.Release(1)
	fmt.Println("Wait...")
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
}

func main() {
	ctx := context.TODO()
	go longPrecess(ctx)
	go longPrecess(ctx)
	go longPrecess(ctx)
	time.Sleep(5 * time.Second)
}
*/
