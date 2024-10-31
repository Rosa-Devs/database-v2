# Object Storage


## L1

Storage - main folder of the db engine
|
|_database - folder where db folder is storing
|_manifest - folder where db manifests is storing


## L2

database
|
|_130288c2-ab49-482b-a6f1-b6ef2b59d914 - database folder
    |_messages - pool folder
        |_3ffd0a10-885e-413c-861b-fa25a4631447.json -- message record
|_61e1e83c-9172-471b-96c5-82afee8d2111 - database folder

manifest
|
|
|_130288c2-ab49-482b-a6f1-b6ef2b59d914.json - manifest file for the db
|_61e1e83c-9172-471b-96c5-82afee8d2111.json - manifest file for the db

### manifest model

```json
{
    "id": "130288c2-ab49-482b-a6f1-b6ef2b59d914", 
    "name": "Chat for Anime lovers", 
    "encryption_algorithm": "AES-256-GCM",
    "secret": "jioprtwe4gjikoptgrxgg", 
    "created_at": "2024-10-17T12:00:00Z",
    "checksum": "e99a18c428cb38d5f260853678922e03",
}
```

### record model

```json
{
    // Base model
    "id": "3ffd0a10-885e-413c-861b-fa25a4631447",
    "pool_name": "messages",
    "created_at": "2024-10-17T23:27:11.205Z",
    "updated_at": "2024-10-17T23:27:11.205Z",
    "updated_by": "user_id",
    "deleted_at": "null or time"

    // Other data by user
}
```
