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

		collection, err := dao.FindCollectionByNameOrId("082x8g93l1ucie1")
		if err != nil {
			return err
		}

		// add
		new_displayorder := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "jgxj30vy",
			"name": "displayorder",
			"type": "number",
			"required": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null
			}
		}`), new_displayorder)
		collection.Schema.AddField(new_displayorder)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("082x8g93l1ucie1")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("jgxj30vy")

		return dao.SaveCollection(collection)
	})
}
