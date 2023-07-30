package biz

import (
	"context"
	"fmt"
	"log"
)

type Demo struct {
	Id   int
	Age  int32
	Name string
}

type IDemoRepo interface {
	// GetProdict(ctx context.Context, uid int64) (*Demo, error)
	CreateDemo(ctx context.Context, p *Demo) (*Demo, error)
	// DeleteDemo(ctx context.Context, uid int64) error
}

type DemoBiz struct {
	repo IDemoRepo
}

func NewDemoBiz(repo IDemoRepo) *DemoBiz {
	return &DemoBiz{repo: repo}
}

func (uc *DemoBiz) CreateDemo(ctx context.Context, age int32, name string) (p *Demo, err error) {
	fmt.Println("age:", age, "+name:", name)
	var (
		bp = &Demo{Age: age, Name: name}
	)
	log.Println(bp)
	out, err := uc.repo.CreateDemo(ctx, bp)
	log.Println("biz cp fin")
	if err != nil {
		return nil, err
	}
	log.Println(out)
	return out, nil
}
