-- name: CreatePayment :one
INSERT INTO payments (player_id, stripe_id, amount, status, payment_date)
VALUES ($1, $2, $3, $4, NOW())
RETURNING *;

-- name: GetPaymentById :one
SELECT * FROM payments WHERE id = $1;

-- name: GetPaymentByStripeId :one
SELECT * FROM payments WHERE stripe_id = $1;

-- name: ListPayments :many
SELECT * FROM payments
ORDER BY payment_date DESC;

-- name: ListPaymentsByPlayer :many
SELECT * FROM payments
WHERE player_id = $1
ORDER BY payment_date DESC;

-- name: ListPaymentsByStatus :many
SELECT * FROM payments
WHERE status = $1
ORDER BY payment_date DESC;

-- name: UpdatePaymentStatus :exec
UPDATE payments
SET status = $1
WHERE id = $2;

-- name: DeletePayment :exec
DELETE FROM payments
WHERE id = $1;

-- name: GetPaymentWithPlayer :one
SELECT py.*, p.user_id, u.email, u.first_name, u.last_name
FROM payments py
INNER JOIN players p ON py.player_id = p.id
INNER JOIN users u ON p.user_id = u.id
WHERE py.id = $1;

-- name: ListPaymentsWithPlayerInfo :many
SELECT py.*, u.first_name, u.last_name, u.email
FROM payments py
INNER JOIN players p ON py.player_id = p.id
INNER JOIN users u ON p.user_id = u.id
ORDER BY py.payment_date DESC;

-- name: GetPlayerPaymentSummary :one
SELECT 
    player_id,
    COUNT(*) as total_payments,
    SUM(CASE WHEN status = 'completed' THEN amount ELSE 0 END) as total_paid,
    SUM(CASE WHEN status = 'pending' THEN amount ELSE 0 END) as total_pending,
    SUM(CASE WHEN status = 'failed' THEN amount ELSE 0 END) as total_failed
FROM payments
WHERE player_id = $1
GROUP BY player_id;
