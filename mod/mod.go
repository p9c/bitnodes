package mod

type Order struct {
	TxnID string `json:"txnid"`

	Nick           string `json:"nick"`
	CoinSourceCode string `json:"csc"`
	Email          string `json:"email"`
	Pack           string `json:"pack"`

	Amount  string `json:"amount"`
	AmountI string `json:"amounti"`

	Timeout   uint32 `json:"timeout"`
	StatusURL string `json:"status_url"`
	QRCodeURL string `json:"qrcode_url"`

	Address string `json:"address"`

	Time       uint32 `json:"time"`
	Status     string `json:"status"`
	StatusText string `json:"status_text"`
	Currency   string `json:"currency"`

	ConfirmsNeeded string `json:"cnfnd"`
	Confirms       string `json:"confirms"`

	Fee  string `json:"fee"`
	FeeI string `json:"feei"`

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
