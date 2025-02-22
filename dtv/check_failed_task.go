package dtv

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/kounoike/dtv-discord-go/db"
	"github.com/kounoike/dtv-discord-go/discord"
	"github.com/kounoike/dtv-discord-go/tasks"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (dtv *DTVUsecase) CheckFailedTask(ctx context.Context) error {
	dtv.logger.Debug("Start CheckFailedTask")
	if dtv.inspector == nil {
		return nil
	}
	taskInfoList, err := dtv.inspector.ListArchivedTasks("default")
	if err != nil {
		return err
	}
	for _, taskInfo := range taskInfoList {
		_, err := dtv.queries.GetEncodeTaskByTaskID(ctx, taskInfo.ID)
		if errors.Cause(err) == sql.ErrNoRows {
			var payload tasks.ProgramEncodePayload
			err := json.Unmarshal(taskInfo.Payload, &payload)
			if err != nil {
				dtv.logger.Warn("task payload json.Unmarshal error", zap.Error(err))
				continue
			}
			err = dtv.queries.InsertEncodeTask(ctx, db.InsertEncodeTaskParams{TaskID: taskInfo.ID, Status: "fail"})
			if err != nil {
				dtv.logger.Warn("failed to InsertEncodeTask", zap.Error(err))
				continue
			}
			dtv.discord.SendMessage(discord.InformationCategory, discord.RecordingFailedChannel, fmt.Sprintf("**エンコード失敗**`%s`のエンコードが失敗しました", payload.ContentPath))
		}
	}
	return nil
}
