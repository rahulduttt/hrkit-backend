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
		new_description := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "5ujegrwq",
			"name": "description",
			"type": "text",
			"required": true,
			"unique": false,
			"options": {
				"min": 10,
				"max": 100,
				"pattern": ""
			}
		}`), new_description)
		collection.Schema.AddField(new_description)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("r5nb86123fmssfv")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("5ujegrwq")

		return dao.SaveCollection(collection)
	})
}
