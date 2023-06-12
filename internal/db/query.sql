-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetExpenses :many
SELECT * FROM expenses
WHERE user_id = $1 AND created_at >= $2 AND created_at <= $3;

-- name: GetExpensesSummary :many
SELECT type, SUM(price) FROM expenses
WHERE user_id = $1 AND created_at >= $2 AND created_at <= $3
GROUP BY type;

-- name: CreateExpense :one
INSERT INTO expenses (user_id, type, description, price, tax_percentage)
VALUES ($1, $2, $3, $4, $5) RETURNING *;