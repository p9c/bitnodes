package mod

type Order struct {
	ID        string `json:"id"`
	Price     string `json:"price"`
	Confirmed bool   `json:"confirmed"`
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
