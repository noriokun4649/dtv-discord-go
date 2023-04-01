// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: program_recording.sql

package db

import (
	"context"
)

const deleteProgramRecordingByProgramId = `-- name: DeleteProgramRecordingByProgramId :exec
DELETE FROM program_recording WHERE program_id = ?
`

func (q *Queries) DeleteProgramRecordingByProgramId(ctx context.Context, programID int64) error {
	_, err := q.db.ExecContext(ctx, deleteProgramRecordingByProgramId, programID)
	return err
}

const getProgramRecording = `-- name: GetProgramRecording :one
SELECT id, program_id, content_path, created_at, updated_at FROM program_recording WHERE id = ?
`

func (q *Queries) GetProgramRecording(ctx context.Context, id int32) (ProgramRecording, error) {
	row := q.db.QueryRowContext(ctx, getProgramRecording, id)
	var i ProgramRecording
	err := row.Scan(
		&i.ID,
		&i.ProgramID,
		&i.ContentPath,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getProgramRecordingByProgramId = `-- name: GetProgramRecordingByProgramId :one
SELECT id, program_id, content_path, created_at, updated_at FROM program_recording WHERE program_id = ?
`

func (q *Queries) GetProgramRecordingByProgramId(ctx context.Context, programID int64) (ProgramRecording, error) {
	row := q.db.QueryRowContext(ctx, getProgramRecordingByProgramId, programID)
	var i ProgramRecording
	err := row.Scan(
		&i.ID,
		&i.ProgramID,
		&i.ContentPath,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const insertProgramRecording = `-- name: InsertProgramRecording :exec
INSERT INTO program_recording(
    program_id,
    content_path
) VALUES (
    ?,
    ?
)
`

type InsertProgramRecordingParams struct {
	ProgramID   int64  `json:"programID"`
	ContentPath string `json:"contentPath"`
}

func (q *Queries) InsertProgramRecording(ctx context.Context, arg InsertProgramRecordingParams) error {
	_, err := q.db.ExecContext(ctx, insertProgramRecording, arg.ProgramID, arg.ContentPath)
	return err
}
