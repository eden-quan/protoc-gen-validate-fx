package module

import (
	pgs "github.com/lyft/protoc-gen-star/v2"
	"google.golang.org/protobuf/types/descriptorpb"
)

type PgFile pgs.File

type File struct {
	PgFile
	file        pgs.File
	allMessages []*Message
	messages    []*Message
}

func NewFile(file pgs.File) *File {
	f := &File{
		PgFile:      file,
		file:        file,
		messages:    make([]*Message, 0),
		allMessages: make([]*Message, 0),
	}

	return f
}

func (f *File) InputPath() pgs.FilePath {
	return f.file.InputPath()
}

func (f *File) Descriptor() *descriptorpb.FileDescriptorProto {
	return f.file.Descriptor()
}

func (f *File) TransitiveImports() []pgs.File {
	return f.file.TransitiveImports()
}

func (f *File) UnusedImports() []pgs.File {
	return f.file.UnusedImports()
}

func (f *File) Dependents() []pgs.File {
	return f.file.Dependents()
}

func (f *File) Services() []pgs.Service {
	return f.file.Services()
}

func (f *File) SyntaxSourceCodeInfo() pgs.SourceCodeInfo {
	return f.file.SyntaxSourceCodeInfo()
}

func (f *File) PackageSourceCodeInfo() pgs.SourceCodeInfo {
	return f.file.PackageSourceCodeInfo()
}

func (f *File) AllMessages() []pgs.Message {
	msgList := make([]pgs.Message, 0)
	for _, m := range f.allMessages {
		msgList = append(msgList, m)
	}

	return msgList
}

func (f *File) Messages() []pgs.Message {
	msgList := make([]pgs.Message, 0)
	for _, m := range f.messages {
		msgList = append(msgList, m)
	}

	return msgList
}
