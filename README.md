# microservice-achievements
Achievement microservice for our online game framework.

## Achievement endpoints

`GET` `/achievements` Returns json data about every achievement.

`GET` `/achievements/{id}` Returns json data about a specific achievement. `id=[string]`

`GET` `/health/live` Returns a Status OK when live.

`GET` `/health/ready` Returns a Status OK when ready or an error when dependencies are not available.

`POST` `/achievements` Add new achievement with specific data. </br>
__Data Params__
```json
{
  "name":        "string, required",
  "condition":   "string, required",
  "description": "string",
  "sprite_id":   "string",
}
```

`PUT` `/achievements` Update achievement data. </br>
__Data Params__
```json
{
  "id":          "string, required",
  "name":        "string",
  "condition":   "string",
  "description": "string",
  "sprite_id":   "string",
}
```

`DELETE` `/achievements/{id}` Delete an achievement.  `id=[string]`
