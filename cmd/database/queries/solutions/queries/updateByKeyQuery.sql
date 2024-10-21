UPDATE
    "solutions"
SET
    key = COALESCE($2, key),
    file_name = COALESCE($3, file_name)
WHERE
    key = $1;