package server

import (
	"context"
	"miopkg/application"
	grpcsvr "miopkg/grpc/server"
	"miopkg/log"
	mallv1 "redbook/api/domain/mall/v1"
	"redbook/app/domain/mall/internal/biz"
	"redbook/app/domain/mall/internal/data"
)

type Engine struct {
	application.Application
}

func NewEngine() *Engine {
	eng := &Engine{}

	if err := eng.Startup(
		eng.NewGrpcServer,
	); err != nil {
		log.Panic("mall grpcsvt startup", log.Any("err", err))
	}
	return eng
}

type MallServer struct {
	mallv1.UnimplementedMallRpcServer
	server        *grpcsvr.Server
	collectionBiz *biz.CollectionBiz
	productBiz    *biz.ProductBiz
	cartBiz       *biz.CartBiz
	shopBiz       *biz.ShopBiz
	orderBiz      *biz.OrderBiz
}

func BuildMallServer(ctx context.Context, svr *grpcsvr.Server) *MallServer {
	d, _, err := data.NewData(ctx)
	if err != nil {
		log.Error("mall BuildMallServer NewData()")
		return nil
	}
	iCollectionRepo := data.NewCollectionRepo(d)
	iProductRepo := data.NewProductRepo(d)
	iCartRepo := data.NewCartRepo(d)
	iOrderRepo := data.NewOrderRepo(d)
	iShopRepo := data.NewShopRepo(d)

	collectionBiz := biz.NewCollectionBiz(iCollectionRepo)
	productBiz := biz.NewProductBiz(iProductRepo)
	cartBiz := biz.NewCartBiz(iCartRepo)
	shopBiz := biz.NewShopBiz(iShopRepo)
	orderBiz := biz.NewOrderBiz(iOrderRepo)

	//new 符合biz接口的repo
	return &MallServer{
		collectionBiz: collectionBiz,
		productBiz:    productBiz,
		cartBiz:       cartBiz,
		orderBiz:      orderBiz,
		shopBiz:       shopBiz,
		server:        svr,
	}
}
