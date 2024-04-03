package main

import (
	"fmt"
	"github.com/oklog/ulid/v2"
	"github.com/suzushin54/actor-based-inventory/actors"
	"github.com/suzushin54/actor-based-inventory/service"
)

func main() {
	s := service.NewService()

	id := ulid.Make()
	item := &actors.InventoryItem{
		ID:    id.String(),
		Name:  "商品1",
		Count: 10,
	}

	s.AddInventoryItem(item)

	fmt.Printf("Added item %s:%s to inventory\n", item.ID, item.Name)

	s.RemoveInventoryItem("item1")
}
