UPDATE
    "solutions"
SET
    key = COALESCE(?, key),
    file_name = COALESCE(?, file_name)
WHERE
    key = ?;