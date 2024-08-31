package main

import (
	"github.com/eden-quan/protoc-gen-validate-fx/module"
	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {

	//in, _ := io.ReadAll(os.Stdin)
	//f, _ := os.OpenFile("./test.data", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0766)
	//_, _ = f.Write(in)
	//_ = f.Close()
	//return

	//// load test file
	//tempFile, err := os.OpenFile("./test/test.data", os.O_RDONLY, 0766)
	//if err != nil {
	//	fmt.Println("load test file failed:", err)
	//	return
	//}
	//defer func() { _ = tempFile.Close() }()
	//// replace test file as stdin
	//os.Stdin = tempFile

	optional := uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
	pgs.
		Init(pgs.DebugEnv("DEBUG_PGV"), pgs.SupportedFeatures(&optional)).
		RegisterModule(module.ValidatorForLanguage("go")).
		RegisterPostProcessor(pgsgo.GoFmt()).
		Render()
}
