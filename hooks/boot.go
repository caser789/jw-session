package hooks

import (
	"context"

	"git.garena.com/duanzy/motto/motto"
	"git.garena.com/shopee/core-server/core-logic/clog"
	"github.com/caser789/jw-session/agents"
)

func Boot(payloads ...interface{}) {
	app := payloads[0].(motto.Application)
	clog.Initialize("log")

	templates := &struct {
		accountAgent agents.AccountAgent
	}{}

	dependencies := []struct {
		kind    interface{}
		tag     interface{}
		factory motto.DepFactory
		options *motto.DepOptions
	}{
		{
			// TODO: initialize toB service agent
			templates.accountAgent, nil, nil, &motto.DepOptions{Singleton: true},
		},
	}

	var err error
	for _, dep := range dependencies {
		if err = app.Container().Register(dep.kind, dep.tag, dep.factory, dep.options); err != nil {
			clog.Errorf(context.Background(), "boot|cannot_initialize_dependency|err=%v,dep=%+v", err, dep)
			return
		}
	}

	// TODO: set panic handler
}
