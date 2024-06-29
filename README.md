# contra

## Usage
### Add new provider
```shell
curl --location 'http://0.0.0.0:8080/provider' \
--header 'Content-Type: application/json' \
--data '{
    "name": "testProvider"
}'
```

## Development
Project structure:
- `internal/contrad`: Server code
  - `routes`: Transport logic (serialization, deserialization, status codes)
  - `scenario`: Business logic
  - `repo`: Data access logic
  - `database`: Database transport operations
- `internal/contra`: Client code
