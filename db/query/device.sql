-- name: CreateDevice :one
INSERT INTO "device" (
  device_name,
  room_id,
  user_id
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetDevice :one
SELECT * FROM "device"
WHERE id = $1 LIMIT 1;

-- name: ListDevice :many
SELECT * FROM "device"
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateDevice :one
UPDATE "device"
  set device_name = $2,
  room_id = $3,
  user_id = $4
WHERE id = $1
RETURNING *;

-- name: DeleteDevice :one
DELETE FROM "device"  
WHERE id = $1
RETURNING *;            