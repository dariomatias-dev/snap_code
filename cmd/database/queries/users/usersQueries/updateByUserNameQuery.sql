UPDATE
    "users"
SET
    user_name = COALESCE(?, user_name)
WHERE
    user_name = ?;