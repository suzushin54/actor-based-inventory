package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/oklog/ulid/v2"
	"github.com/suzushin54/actor-based-inventory/actors"
	"github.com/suzushin54/actor-based-inventory/service"
)

func main() {
	s, err := service.NewService("localhost:9092", "inventory")
	if err != nil {
		log.Fatal(err)
	}
	id := ulid.Make()

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		item := &actors.InventoryItem{
			ID:    id.String(),
			Name:  "商品1",
			Count: 10,
		}
		if err := s.AddInventoryItem(item); err != nil {
			return
		}
		fmt.Printf("Added item %s:%s to inventory\n", item.ID, item.Name)
	})

	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		s.UpdateInventoryItemCount(id, 333)
		fmt.Printf("Updated count of item %s to %d\n", id, 333)
	})

	http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
		inventoryItem, err := s.QueryInventoryItem(id)
		if err != nil {
			return
		}

		fmt.Printf("Queried item %s:%s from inventory\n", inventoryItem.ID, inventoryItem.Name)
		fmt.Printf("Current item count: %d\n", inventoryItem.Count)
	})

	http.HandleFunc("/remove", func(w http.ResponseWriter, r *http.Request) {
		s.RemoveInventoryItem(id)
	})

	log.Println("Starting inventory management server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
