-- name: CreatePasswordResetToken :one
INSERT INTO password_reset_tokens (user_id, token, expires_at)
VALUES ($1, $2, $3)
RETURNING id, token, expires_at;

-- name: GetPasswordResetToken :one
SELECT id, user_id, token, expires_at, used, created_at
FROM password_reset_tokens
WHERE token = $1 AND used = FALSE AND expires_at > NOW();

-- name: MarkPasswordResetTokenAsUsed :exec
UPDATE password_reset_tokens
SET used = TRUE
WHERE token = $1;

-- name: DeleteExpiredPasswordResetTokens :exec
DELETE FROM password_reset_tokens
WHERE expires_at < NOW() OR used = TRUE;
