# Healthtracker project
## Healthtracker REST API
```
GET /health/view/:id
POST /health/daily
PUT /health/view/:id
DELETE /health/view/:id
```
## DB Structure
```
Table healthtracker{
    id bigserial [primary key]
    created_at timestamp
    walking text
    hydrate text
    sleep text
    version int
}
```
