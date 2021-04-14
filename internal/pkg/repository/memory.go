package repository

import "github.com/1r0npipe/simple-web-broker/internal/pkg/model"

type MemRepo struct{}

func NewMemRepo() *MemRepo {
	return &MemRepo{}
}

func (mr *MemRepo) Get(getReq string) (string, error) {
	// TODO: impl
	return "", nil
}

func (mr *MemRepo) Put(putReq *model.PutValue) error {
	// TODO: impl
	return nil
}
