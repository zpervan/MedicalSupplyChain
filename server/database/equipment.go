package database

type Item struct {
	Id           string
	Name         string
	Description  string
	ReceivedDate string
	TakenDate    string
	ModifiedDate string
	Quantity     int
	Category     string
}

// Insert a new equipment item into the database.
//
// Returns an error if the creation was not successful.
func (d Database) InsertEquipment(item Item) error {
	return nil
}

// Fetch all equipment items from the database.
//
// Returns an error if the equipment cannot be fetched, i.e. database is not reachable.
func (d Database) FetchEquipment() error {
	return nil
}

// Update the equipment item by item ID.
//
// Returns an error if the update was not successful.
func (d Database) UpdateEquipment(itemId string, item Item) error {
	return nil
}

// Delete the equipment item with the specified item ID.
//
// Returns an error if the update was not successful.
func (d Database) DeleteEquipment(id string) error {
	return nil
}