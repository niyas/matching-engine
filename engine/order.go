package engine

import "encoding/json"

type Order struct {
  Amount uint64 
  Price  uint64 
  ID     string 
  Side   int8   
}

func (order *Order) FromJson(msg [] byte) error {
  return json.Unmarshal(msg, order)
}

func (order *Order) ToJson() []byte {
  str, _ := json.Marshal(order)
  return str
}
