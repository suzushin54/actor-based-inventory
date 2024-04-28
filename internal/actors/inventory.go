package actors

import (
	"fmt"

	"github.com/asynkron/protoactor-go/actor"
)

// InventoryItem represents an item in an inventory.
type InventoryItem struct {
	ID    string
	Name  string
	Count int
}

// InventoryActor is an actor that manages an inventory of items.
type InventoryActor struct {
	Items map[string]*InventoryItem
}

type QueryInventoryItem struct {
	ItemID string
}

type AddInventoryItem struct {
	Item *InventoryItem
}

type UpdateInventoryItemCount struct {
	ItemID string
	Count  int
}

type RemoveInventoryItem struct {
	ItemID string
}

func NewInventoryActor() *InventoryActor {
	return &InventoryActor{
		Items: make(map[string]*InventoryItem),
	}
}

func (actor *InventoryActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *AddInventoryItem:
		actor.Items[msg.Item.ID] = msg.Item
		fmt.Printf("Added item %s to inventory\n", msg.Item.Name)
	case *RemoveInventoryItem:
		delete(actor.Items, msg.ItemID)
		fmt.Printf("Removed item %s from inventory\n", msg.ItemID)
	case *UpdateInventoryItemCount:
		item, ok := actor.Items[msg.ItemID]
		if !ok {
			fmt.Printf("Item %s not found\n", msg.ItemID)
			return
		}
		item.Count = msg.Count
		fmt.Printf("Updated count of item %s to %d\n", msg.ItemID, msg.Count)
	case *QueryInventoryItem:
		item, ok := actor.Items[msg.ItemID]
		if !ok {
			fmt.Printf("Item %s not found\n", msg.ItemID)
			return
		}
		fmt.Printf("Item %s:%s has count %d\n", item.ID, item.Name, item.Count)
		ctx.Respond(item)
	}

}
