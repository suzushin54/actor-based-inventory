package actors

import (
	"log/slog"

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
		slog.Info("Added item to inventory actor", "item_id", msg.Item.ID)
	case *RemoveInventoryItem:
		delete(actor.Items, msg.ItemID)
		slog.Info("Removed item from inventory actor", "item_id", msg.ItemID)
	case *UpdateInventoryItemCount:
		item, ok := actor.Items[msg.ItemID]
		if !ok {
			slog.Warn("Item not found", "item_id", msg.ItemID)
			return
		}
		item.Count = msg.Count
		slog.Info("Updated count of item in inventory actor", "item_id", msg.ItemID, "count", msg.Count)
	case *QueryInventoryItem:
		item, ok := actor.Items[msg.ItemID]
		if !ok {
			slog.Warn("Item not found", "item_id", msg.ItemID)
			return
		}
		slog.Info("Item found in inventory actor", "item_id", msg.ItemID)
		ctx.Respond(item)
	}

}
