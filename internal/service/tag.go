package service

import (
	"context"

	"github.com/goexl/gox/field"
	"github.com/itcoursee/core/tag"
	"github.com/itcoursee/core/vo"
	"github.com/pangum/example/internal/config"
	"github.com/pangum/logging"
)

// Tag 标签
type Tag struct {
	config *config.Tag
	logger logging.Logger
}

func newTag(config *config.Tag, logger logging.Logger) *Tag {
	return &Tag{
		config: config,
		logger: logger,
	}
}

func (t *Tag) Get(_ context.Context, req *tag.GetReq) (rsp *tag.GetRsp, err error) {
	t.logger.Info("调用获取方法", field.New("req", req))

	rsp = new(tag.GetRsp)
	rsp.Tag = new(vo.Tag)
	rsp.Tag.Id = t.config.Id
	rsp.Tag.Name = t.config.Name
	rsp.Tag.Label = t.config.Label
	rsp.Tag.Keywords = t.config.Keywords
	rsp.Tag.Description = t.config.Description

	return
}
