package module

import pgs "github.com/lyft/protoc-gen-star/v2"

type Field struct {
	pgs.Field
	msg pgs.Message
}

func NewField(f pgs.Field, msg pgs.Message) *Field {
	return &Field{
		Field: f,
		msg:   msg,
	}
}

func (f *Field) UpdateMessage(msg pgs.Message) {
	f.msg = msg
}

func (f *Field) Message() pgs.Message {
	return f.msg
}
