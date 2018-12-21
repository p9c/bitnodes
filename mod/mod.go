package mod

type Order struct {
	ID string `json:"id"`

	Nick string `json:"id"`
	ID   string `json:"id"`

	Address    string `json:"address"`
	TxnID      string `json:"txn_id"`
	Status     string `json:"status"`
	StatusText string `json:"status_text"`
	Currency   string `json:"currency"`
	Confirms   string `json:"confirms"`
	Amount     string `json:"amount"`
	AmountI    string `json:"amounti"`
	Fee        string `json:"fee"`
	FeeI       string `json:"feei"`
	DestTag    string `json:"dest_tag"`

	Currency1        string `json:"currency1"`
	Currency2        string `json:"currency2"`
	Amount1          string `json:"amount1"`
	Amount2          string `json:"amount2"`
	Fee              string `json:"fee"`
	BuyerName        string `json:"buyer_name"`
	Email            string `json:"email"`
	ItemName         string `json:"item_name"`
	ItemNumber       string `json:"item_number"`
	Invoice          string `json:"invoice"`
	Custom           string `json:"custom"`
	SendTX           string `json:"send_tx"` // the tx id of the payment to the merchant. only included when 'status' >= 100 and the payment mode is set to ASAP or nightly or if the payment is paypal passthru
	ReceivedAmount   string `json:"received_amount"`
	ReceivedConfirms string `json:"received_confirms"`
}

type Subscription struct {
	ID   string `json:"id"`
	User string `json:"user"`
	Date string `json:"date"`
}

type Node struct {
	GPSID       string `json:"id"`
	NodeID      string `json:"nodeid" form:"nodeid"`
	Coin        string `json:"coin" form:"coin"`
	RPCUser     string `json:"rpcuser" form:"rpcuser"`
	RPCPassword string `json:"rpcpassword" form:"rpcpassword"`
	Addr        string `json:"addr" form:"addr"`
	IP          string `json:"ip" form:"ip"`
	Port        int64  `json:"port" form:"port"`
	Published   bool   `json:"published" form:"published"`
	BitNode     bool   `json:"bitnode" form:"bitnode"`
}

// gpsID := utils.IDGEN()

type Tutorial struct {
	ID      string `json:"id"`
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
}

type Announcement struct {
	ID      string `json:"id"`
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
}
