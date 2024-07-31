-- name: CreateUser :one
INSERT INTO "user" (
  user_email,   
  user_password
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM "user"
WHERE id = $1 LIMIT 1;

-- name: ListUser :many
SELECT * FROM "user"
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE "user"
  set user_email = $2,
  user_password = $3
WHERE id = $1
RETURNING *;

-- name: DeleteUser :one
DELETE FROM "user"  
WHERE id = $1
RETURNING *;