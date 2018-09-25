package engine

import "encoding/json"

type Trade struct {
	TakerOrderID string 
	MakerOrderID string 
	Amount       uint64 
	Price        uint64
}

func (trade *Trade) FromJson(msg [] byte) error {
  return json.Unmarshal(msg, trade)
}

func (trade *Trade) ToJson() []byte {
  str, _ := json.Marshal(trade)
  return str
}
