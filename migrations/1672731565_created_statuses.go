package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `{
			"id": "082x8g93l1ucie1",
			"created": "2023-01-03 07:39:25.431Z",
			"updated": "2023-01-03 07:39:25.431Z",
			"name": "statuses",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "chculdjd",
					"name": "name",
					"type": "text",
					"required": true,
					"unique": true,
					"options": {
						"min": 5,
						"max": 15,
						"pattern": ""
					}
				},
				{
					"system": false,
					"id": "reixsgxh",
					"name": "description",
					"type": "text",
					"required": false,
					"unique": false,
					"options": {
						"min": 5,
						"max": 25,
						"pattern": ""
					}
				}
			],
			"listRule": null,
			"viewRule": null,
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("082x8g93l1ucie1")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
