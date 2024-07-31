-- name: CreateRoom :one
INSERT INTO "room" (
  room_name
) VALUES (
  $1
)
RETURNING *;

-- name: GetRoom :one
SELECT * FROM "room"
WHERE id = $1 LIMIT 1;

-- name: ListRoom :many
SELECT * FROM "room"
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateRoom :one
UPDATE "room"
  set room_name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteRoom :one
DELETE FROM "room"  
WHERE id = $1
RETURNING *;            