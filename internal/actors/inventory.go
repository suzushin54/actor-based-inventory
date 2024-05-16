package actors

import (
	"log/slog"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/oklog/ulid/v2"
)

type ID ulid.ULID

func (id ID) String() string {
	return ulid.ULID(id).String()
}

// InventoryItem represents an item in an inventory.
type InventoryItem struct {
	ID    ID
	Name  string
	Count int
}

// InventoryActor is an actor that manages an inventory of items.
type InventoryActor struct {
	Items map[ID]*InventoryItem
}

type QueryInventoryItem struct {
	ItemID ID
}

type AddInventoryItem struct {
	Item *InventoryItem
}

type UpdateInventoryItem struct {
	Item *InventoryItem
}

type RemoveInventoryItem struct {
	ItemID ID
}

func NewInventoryActor() *InventoryActor {
	return &InventoryActor{
		Items: make(map[ID]*InventoryItem),
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
	case *UpdateInventoryItem:
		item, ok := actor.Items[msg.Item.ID]
		if !ok {
			slog.Warn("Item not found", "item_id", msg.Item.ID)
			return
		}
		item.Name = msg.Item.Name
		item.Count = msg.Item.Count

		slog.Info("Updated count of item in inventory actor", "item", item)
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
