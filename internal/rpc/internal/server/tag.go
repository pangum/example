package server

import (
	"github.com/pangum/example/internal/rpc/internal/unimplemented"
	"github.com/pangum/example/internal/service"
)

type Tag struct {
	unimplemented.TagRpc
	*service.Tag
}

func newTag(service *service.Tag) *Tag {
	return &Tag{
		Tag: service,
	}
}
