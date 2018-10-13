package api

import (
	"context"
)

type DBEntityChecker interface {
	dbCheckEntity(ctx context.Context) (bool, error)
}

type DBEntityCreator interface {
	dbCreate(ctx context.Context) (error)
}

func DBCheckEntity(entity DBEntityChecker, ctx context.Context) (bool, error){
	if entityExists, err := entity.dbCheckEntity(ctx); err != nil {
		return entityExists, err
	} else {
		return entityExists, nil
	}
}

func DBCreate(entity DBEntityCreator, ctx context.Context) error {
	if err:= entity.dbCreate(ctx); err != nil {
		return err
	}
	return nil
}

