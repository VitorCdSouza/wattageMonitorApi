-- name: CreateReading :one
INSERT INTO "reading" (
  reading_wattage,
  reading_hour,
  device_id
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetReading :one
SELECT * FROM "reading"
WHERE id = $1 LIMIT 1;

-- name: ListReading :many
SELECT * FROM "reading"
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteReading :one
DELETE FROM "reading"  
WHERE id = $1
RETURNING *;            