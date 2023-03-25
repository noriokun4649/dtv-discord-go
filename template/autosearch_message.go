package template

import (
	"bytes"
	"text/template"

	"github.com/kounoike/dtv-discord-go/db"
)

type AutoSearchMessageTemplateArgs struct {
	Program           db.Program
	Service           db.Service
	ProgramMessageURL string
}

var autoSearchMessageTemplateString = `================================================
{{ .Program.Name }}
{{ .Service.Name }}
{{ .Program.StartAt |toTimeStr }}～{{ .Program.Duration | toDurationStr }}
{{ .ProgramMessageURL }}
`

func GetAutoSearchMessage(program db.Program, service db.Service, programMessageURL string) (string, error) {
	funcMap := map[string]interface{}{
		"toTimeStr":     toTimeStr,
		"toDurationStr": toDurationStr,
		"toExtendStr":   toExtendStr,
	}
	tmpl, err := template.New("autosearch-message").Funcs(funcMap).Parse(autoSearchMessageTemplateString)
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	args := AutoSearchMessageTemplateArgs{
		Program:           program,
		Service:           service,
		ProgramMessageURL: programMessageURL,
	}
	err = tmpl.Execute(&b, args)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}
