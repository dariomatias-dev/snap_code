UPDATE
    "users"
SET
    user_name = COALESCE($2, user_name)
WHERE
    user_name = $1;