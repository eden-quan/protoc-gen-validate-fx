package main

import (
	pgs "github.com/lyft/protoc-gen-star/v2"
	"google.golang.org/protobuf/types/pluginpb"

	"github.com/eden-quan/protoc-gen-validate-fx/module"
)

func main() {
	optional := uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
	pgs.
		Init(pgs.DebugEnv("DEBUG_PGV"), pgs.SupportedFeatures(&optional)).
		RegisterModule(module.ValidatorForLanguage("cc")).
		Render()
}
