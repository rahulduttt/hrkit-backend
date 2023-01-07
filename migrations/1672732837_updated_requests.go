package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("r5nb86123fmssfv")
		if err != nil {
			return err
		}

		// add
		new_status := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "dgldgqtc",
			"name": "status",
			"type": "relation",
			"required": true,
			"unique": false,
			"options": {
				"maxSelect": 1,
				"collectionId": "082x8g93l1ucie1",
				"cascadeDelete": false
			}
		}`), new_status)
		collection.Schema.AddField(new_status)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("r5nb86123fmssfv")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("dgldgqtc")

		return dao.SaveCollection(collection)
	})
}
