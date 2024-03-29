package service

import (
	"github.com/1r0npipe/simple-web-broker/internal/pkg/model"
)

type (
	repo interface {
		Get(key string) (string, error)
		Put(putReq *model.PutValue) error
	}

	Service struct {
		repo repo
	}
)

func New(repo repo) *Service {
	return &Service{repo: repo}
}
