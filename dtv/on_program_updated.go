package dtv

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"

	"github.com/kounoike/dtv-discord-go/db"
	"github.com/kounoike/dtv-discord-go/discord"
	"github.com/kounoike/dtv-discord-go/template"
	"github.com/pkg/errors"
	"golang.org/x/exp/slog"
)

func (dtv *DTVUsecase) OnProgramsUpdated(serviceId uint) {
	service, err := dtv.mirakc.GetService(serviceId)
	_ = service
	if err != nil {
		slog.Error("mirakc GetService Error", err)
		return
	}
	_, err = dtv.discord.SendMessage(discord.InformationCategory, discord.LogChannel, fmt.Sprintf("programs updated: %s", service.Name))
	if err != nil {
		slog.Error("discord SendMessage error", err)
		return
	}
	programs, err := dtv.mirakc.ListPrograms(serviceId)
	if err != nil {
		slog.Error("mirakc ListPrograms Error", err)
		return
	}

	ctx := context.Background()
	for _, p := range programs {
		if p.Name == "" {
			continue
		}
		program, err := dtv.queries.GetProgram(ctx, p.ID)
		if errors.Cause(err) == sql.ErrNoRows {
			msg, err := template.GetProgramMessage(p, *service)
			if err != nil {
				slog.Error("template GetProgramMessage error", err)
				return
			}
			msgID, err := dtv.discord.SendMessage("録画-番組情報", service.Name, msg)
			if err != nil {
				slog.Error("discord SendMessage error", err)
				return
			}
			err = dtv.queries.InsertProgram(ctx, p)
			if err != nil {
				slog.Error("InsertProgram error", err)
				return
			}
			err = dtv.queries.InsertProgramMessage(ctx, db.InsertProgramMessageParams{MessageID: msgID, ProgramID: p.ID})
			if err != nil {
				slog.Error("InsertProgramMessage error", err)
				return
			}
			params := db.InsertProgramServiceParams{
				ProgramID: p.ID,
				ServiceID: service.ID,
			}
			err = dtv.queries.InsertProgramService(ctx, params)
			if err != nil {
				slog.Error("InsertProgramService error", err)
				return
			}
		} else {
			pJson, err := p.Json.MarshalJSON()
			if err != nil {
				continue
			}
			programJson, err := program.Json.MarshalJSON()
			if err != nil {
				continue
			}
			if !bytes.Equal(pJson, programJson) {
				dtv.queries.UpdateProgram(ctx, p)
			}
		}
	}
}
