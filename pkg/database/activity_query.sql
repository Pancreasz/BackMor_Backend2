-- name: ListActivities :many
SELECT * FROM activities;


-- name: InsertActivity :one
INSERT INTO activities (
    creator_id,
    title,
    description,
    start_time,
    end_time,
    max_participants,
    visibility,
    latitude,
    longitude,
    location
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
RETURNING *;

-- name: ListActivityMembers :many
SELECT
  am.user_id,
  u.display_name,
  am.role,
  am.joined_at
FROM activity_members am
JOIN users u ON u.id = am.user_id
WHERE am.activity_id = $1
ORDER BY am.joined_at ASC;

-- name: AddActivityMember :one
INSERT INTO activity_members (activity_id, user_id, role, joined_at)
VALUES ($1, $2, $3, NOW())
RETURNING activity_id, user_id, role, joined_at;

-- name: GetActivityByID :one
SELECT *
FROM activities
WHERE id = $1;

-- name: ListActivitiesByUser :many
SELECT a.*
FROM activities a
JOIN activity_members am ON a.id = am.activity_id
WHERE am.user_id = $1
ORDER BY a.start_time;

-- name: DeleteActivity :exec
DELETE FROM activities
WHERE id = $1;

-- name: DeleteActivityMember :exec
DELETE FROM activity_members
WHERE activity_id = $1 AND user_id = $2;

