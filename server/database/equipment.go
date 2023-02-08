package database

import (
	"database/sql"
	"fmt"
	ms "server/proto"
	"time"
)

// Insert a new equipment item into the database.
//
// Returns an error if the creation was not successful.
//
// TODO: Implement ID and received date functionality
func (d Database) InsertEquipment(item *ms.Item) error {
	if item == nil {
		return fmt.Errorf("cannot insert new item into equipment. reason: passed item data is empty")
	}

	_, err := d.handler.Exec("INSERT INTO equipment(user_id, equipment_inventory_id, equipment_category_id, name, description, received_date) VALUES (NULL, NULL, NULL, $1, $2, $3)",
		item.Name, item.Description, time.Now().Format(time.RFC1123))
	if err != nil {
		d.logger.Error("cannot insert new item. reason: " + err.Error())
		return err
	}

	return nil
}

// Fetch all equipment items from the database.
//
// Returns an error if the equipment cannot be fetched, i.e. database is not reachable.
//
// TODO: Implement proper ID functionalities while fetching
func (d Database) FetchEquipment() ([]*ms.Item, error) {
	rows, err := d.handler.Query("SELECT * FROM equipment;")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var items []*ms.Item
	// TODO: Implement ID functionality
	var dummyId sql.NullString

	for rows.Next() {
		item := &ms.Item{}
		err := rows.Scan(&item.Id, &dummyId, &dummyId, &dummyId, &item.Name, &item.Description, &item.ReceivedDate)

		if err != nil {
			d.logger.Error("could not fetch equipment row. reason: " + err.Error())
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

// Update the equipment item by item ID.
//
// Returns an error if the update was not successful.
func (d Database) UpdateEquipment(itemId string, item *ms.Item) error {
	return nil
}

// Delete the equipment item with the specified item ID.
//
// Returns an error if the update was not successful.
func (d Database) DeleteEquipment(id string) error {
	return nil
}
