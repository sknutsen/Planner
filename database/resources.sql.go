// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: resources.sql

package database

import (
	"context"
)

const createResource = `-- name: CreateResource :exec
INSERT INTO resources (
    plan_id,
    title,
    resource_type,
    content
) VALUES (
    ?,
    ?,
    ?,
    ?
)
`

type CreateResourceParams struct {
	PlanID       int64
	Title        string
	ResourceType int64
	Content      interface{}
}

func (q *Queries) CreateResource(ctx context.Context, arg CreateResourceParams) error {
	_, err := q.db.ExecContext(ctx, createResource,
		arg.PlanID,
		arg.Title,
		arg.ResourceType,
		arg.Content,
	)
	return err
}

const deleteResource = `-- name: DeleteResource :exec
DELETE FROM resources
WHERE id IN (SELECT r.id FROM resources as r
             INNER JOIN plans as p ON r.plan_id = p.id
             LEFT OUTER JOIN plan_access as pa ON p.id = pa.plan_id
             WHERE r.id = ? AND (p.user = ? OR pa.user = ?))
`

type DeleteResourceParams struct {
	ID     int64
	User   string
	User_2 string
}

func (q *Queries) DeleteResource(ctx context.Context, arg DeleteResourceParams) error {
	_, err := q.db.ExecContext(ctx, deleteResource, arg.ID, arg.User, arg.User_2)
	return err
}

const getResource = `-- name: GetResource :one
SELECT 
    r.id, r.title, r.resource_type, r.content, r.plan_id 
FROM resources as r
INNER JOIN plans as p ON r.plan_id = p.id
LEFT OUTER JOIN plan_access as pa ON p.id = pa.plan_id
WHERE r.id = ? AND (p.user = ? OR pa.user = ?)
`

type GetResourceParams struct {
	ID     int64
	User   string
	User_2 string
}

func (q *Queries) GetResource(ctx context.Context, arg GetResourceParams) (Resource, error) {
	row := q.db.QueryRowContext(ctx, getResource, arg.ID, arg.User, arg.User_2)
	var i Resource
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.ResourceType,
		&i.Content,
		&i.PlanID,
	)
	return i, err
}

const getResourcesByPlan = `-- name: GetResourcesByPlan :many
SELECT
  r.id, r.title, r.resource_type, r.content, r.plan_id
FROM
  resources AS r
WHERE
  r.plan_id IN (
    SELECT
      p.id
    FROM
      plans AS p
      LEFT OUTER JOIN plan_access AS pa ON p.id = pa.plan_id
    WHERE
      p.id = ?
      AND (
        p.user = ?
        OR pa.user = ?
      )
  )
ORDER BY r.title ASC
`

type GetResourcesByPlanParams struct {
	ID     int64
	User   string
	User_2 string
}

func (q *Queries) GetResourcesByPlan(ctx context.Context, arg GetResourcesByPlanParams) ([]Resource, error) {
	rows, err := q.db.QueryContext(ctx, getResourcesByPlan, arg.ID, arg.User, arg.User_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Resource
	for rows.Next() {
		var i Resource
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.ResourceType,
			&i.Content,
			&i.PlanID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateResource = `-- name: UpdateResource :exec
UPDATE resources 
SET title = ?, resource_type = ?, content = ?
WHERE id IN (SELECT r.id FROM resources as r
             INNER JOIN plans as p ON r.plan_id = p.id
             LEFT OUTER JOIN plan_access as pa ON p.id = pa.plan_id
             WHERE r.id = ? AND (p.user = ? OR pa.user = ?))
`

type UpdateResourceParams struct {
	Title        string
	ResourceType int64
	Content      interface{}
	ID           int64
	User         string
	User_2       string
}

func (q *Queries) UpdateResource(ctx context.Context, arg UpdateResourceParams) error {
	_, err := q.db.ExecContext(ctx, updateResource,
		arg.Title,
		arg.ResourceType,
		arg.Content,
		arg.ID,
		arg.User,
		arg.User_2,
	)
	return err
}
