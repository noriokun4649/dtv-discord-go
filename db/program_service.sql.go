// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: program_service.sql

package db

import (
	"context"
)

const getProgramServiceByProgramID = `-- name: GetProgramServiceByProgramID :one
SELECT id, program_id, service_id, created_at, updated_at FROM ` + "`" + `program_service` + "`" + ` WHERE ` + "`" + `program_id` + "`" + ` = ?
`

func (q *Queries) GetProgramServiceByProgramID(ctx context.Context, programID int64) (ProgramService, error) {
	row := q.db.QueryRowContext(ctx, getProgramServiceByProgramID, programID)
	var i ProgramService
	err := row.Scan(
		&i.ID,
		&i.ProgramID,
		&i.ServiceID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getProgramServiceByServiceID = `-- name: GetProgramServiceByServiceID :one
SELECT id, program_id, service_id, created_at, updated_at FROM ` + "`" + `program_service` + "`" + ` WHERE ` + "`" + `service_id` + "`" + ` = ?
`

func (q *Queries) GetProgramServiceByServiceID(ctx context.Context, serviceID int64) (ProgramService, error) {
	row := q.db.QueryRowContext(ctx, getProgramServiceByServiceID, serviceID)
	var i ProgramService
	err := row.Scan(
		&i.ID,
		&i.ProgramID,
		&i.ServiceID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getServiceByProgramID = `-- name: GetServiceByProgramID :one
SELECT service.id, service.service_id, service.network_id, service.type, service.logo_id, service.remote_control_key_id, service.name, service.channel_type, service.channel, service.has_logo_data, service.created_at, service.updated_at
FROM ` + "`" + `service` + "`" + `
JOIN ` + "`" + `program_service` + "`" + ` on ` + "`" + `program_service` + "`" + `.` + "`" + `service_id` + "`" + ` = ` + "`" + `service` + "`" + `.` + "`" + `id` + "`" + `
WHERE ` + "`" + `program_service` + "`" + `.` + "`" + `program_id` + "`" + ` = ?
`

func (q *Queries) GetServiceByProgramID(ctx context.Context, programID int64) (Service, error) {
	row := q.db.QueryRowContext(ctx, getServiceByProgramID, programID)
	var i Service
	err := row.Scan(
		&i.ID,
		&i.ServiceID,
		&i.NetworkID,
		&i.Type,
		&i.LogoID,
		&i.RemoteControlKeyID,
		&i.Name,
		&i.ChannelType,
		&i.Channel,
		&i.HasLogoData,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const insertProgramService = `-- name: InsertProgramService :exec
INSERT INTO ` + "`" + `program_service` + "`" + ` (
    ` + "`" + `program_id` + "`" + `,
    ` + "`" + `service_id` + "`" + `
) VALUES (?, ?)
`

type InsertProgramServiceParams struct {
	ProgramID int64 `json:"programID"`
	ServiceID int64 `json:"serviceID"`
}

func (q *Queries) InsertProgramService(ctx context.Context, arg InsertProgramServiceParams) error {
	_, err := q.db.ExecContext(ctx, insertProgramService, arg.ProgramID, arg.ServiceID)
	return err
}
