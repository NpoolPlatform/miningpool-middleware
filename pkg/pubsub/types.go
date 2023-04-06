package pubsub

import (
	"github.com/google/uuid"
)

type MsgBase struct {
	MID    string
	Sender string
	RID    *uuid.UUID
	UnID   *uuid.UUID
}

type Resp struct {
	Code int
	Msg  string
}

type Msg struct {
	MsgBase
	Body string
}