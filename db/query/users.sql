-- name: CreateUser :one
INSERT INTO "users" (
    username,
    password,
    email
) VALUES (
    $1, $2, sqlc.narg('email')
) RETURNING *;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;