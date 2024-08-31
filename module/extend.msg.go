package module

import (
	flatten2 "github.com/eden-quan/protoc-gen-validate-fx/flatten"
	"github.com/eden-quan/protoc-gen-validate-fx/flatten/meta"
	pgs "github.com/lyft/protoc-gen-star/v2"
)

func ProcessFlatten(msg pgs.Message) *Message {
	fields := make([]pgs.Field, 0)
	for _, f := range msg.Fields() {
		_, flatten := flatten2.ExtractFlatten(msg, f)
		if flatten {
			fieldMsg := f.Type().Embed()
			if fieldMsg == nil {
				continue
			}

			fieldMsg = ProcessFlatten(fieldMsg)
			subFields := make([]pgs.Field, 0)
			wrapMsg := NewMessage(msg, nil)
			for _, f := range fieldMsg.Fields() {
				f = NewField(f, wrapMsg)
				subFields = append(subFields, f)
				//fields = append(fields, f)
			}

			wrapMsg.AddFields(subFields...)
			fields = append(fields, subFields...)

		} else {
			fields = append(fields, f)
		}
	}

	return NewMessage(msg, fields)
}

type PgMessage pgs.Message

type Message struct {
	PgMessage

	ErrorDef *meta.ValidateErrorDefine
	fields   []pgs.Field
}

func NewMessage(m pgs.Message, fs []pgs.Field) *Message {
	define := flatten2.ExtractErrorDefine(m)
	return &Message{
		PgMessage: m,
		ErrorDef:  define,
		fields:    fs,
	}
}

func (m *Message) AddFields(fields ...pgs.Field) {
	m.fields = append(m.fields, fields...)
}

func (m *Message) Fields() []pgs.Field {
	return m.fields
}

func (m *Message) OneOfFields() (f []pgs.Field) {
	for _, o := range m.OneOfs() {
		f = append(f, o.Fields()...)
	}

	return f
}

func (m *Message) NonOneOfFields() (f []pgs.Field) {
	for _, fld := range m.fields {
		if !fld.InOneOf() {
			f = append(f, fld)
		}
	}
	return f
}

func (m *Message) SyntheticOneOfFields() (f []pgs.Field) {
	for _, o := range m.OneOfs() {
		if o.IsSynthetic() {
			f = append(f, o.Fields()...)
		}
	}

	return
}

func (m *Message) OneOfs() []pgs.OneOf {
	return m.PgMessage.OneOfs()
}

func (m *Message) RealOneOfs() (r []pgs.OneOf) {
	for _, o := range m.OneOfs() {
		if !o.IsSynthetic() {
			r = append(r, o)
		}
	}

	return
}
