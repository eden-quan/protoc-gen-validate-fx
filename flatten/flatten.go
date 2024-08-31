package flatten

import (
	pgs "github.com/lyft/protoc-gen-star/v2"
	"github.com/sinsay/protoc-gen-validate-fx/flatten/meta"
)

type FlatInfo struct {
	Rule    *meta.FlattenRules
	Message pgs.Message
	Field   pgs.Field
}

func ExtractErrorDefine(msg pgs.Message) *meta.ValidateErrorDefine {
	bizCode := int64(10000000)
	httpCode := int64(400)
	define := &meta.ValidateErrorDefine{
		BizCode:  &bizCode,
		HttpCode: &httpCode,
	}

	exists, _ := msg.Extension(meta.E_MsgError, define)
	if exists {
		return define
	}

	_, _ = msg.File().Extension(meta.E_FileError, define)
	return define
}

func ExtractFlatten(msg pgs.Message, f pgs.Field) (*FlatInfo, bool) {

	var flatten bool
	var ext *meta.FlattenRules
	_, err := f.Extension(meta.E_Flatten, &flatten)
	if err == nil {
		_, err = f.Extension(meta.E_FlattenRule, ext)
	}

	if err != nil {
		return nil, false
	}

	if ext == nil {
		// create default ext
		m := int32(0)
		ext = &meta.FlattenRules{Reserved: &meta.Reserved{
			Min: &m, // useless now
			Max: &m, // useless now
		}}
	} else {
		flatten = true
	}

	return &FlatInfo{
		Rule:    ext,
		Message: msg,
		Field:   f,
	}, flatten
}
