package main

type MainProcess struct {
	MainProcessInfo      *mainProcessInfo
	ActivePositionsCount int64
	Action               chan Buy
}

type mainProcessInfo struct {}

// Actions possible for the MainProcess to execute
type Action chan struct {
	Buy  *Buy
	Sell *Sell
}

// Buy instruction details
type Buy struct {
}

// Buy instruction details
type Sell struct {
}


func main() {
	//login
	LoadCreds("./")

}
