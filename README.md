# contra

## Development
Project structure:
- `internal/contrad`: Server code
  - `routes`: Transport logic (serialization, deserialization, status codes)
  - `scenario`: Business logic
  - `repo`: Data access logic
  - `database`: Database transport operations
- `internal/contra`: Client code
