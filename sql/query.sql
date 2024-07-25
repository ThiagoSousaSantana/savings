-- name: CreateExpense :exec
INSERT INTO expense (description, value) VALUES ($1, $2);

-- name: ListExpenses :many
SELECT * FROM expense;

-- name: GetExpenseById :one
SELECT * FROM expense WHERE id = $1;

-- name: UpdateExpense :one
UPDATE expense SET description = $2, value = $3 WHERE id = $1 RETURNING *;

-- name: DeleteExpense :exec
DELETE FROM expense WHERE id = $1;

