package event

import (
	"github.com/mikestefanello/hooks"
	"its.ac.id/base-go/pkg/app/common"
)

var HookEvent = hooks.NewHook[*common.Event]("event")

func init() {

}
