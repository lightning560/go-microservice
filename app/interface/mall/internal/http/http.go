package server

import (
	"redbook/app/interface/mall/internal/biz"
	"redbook/app/interface/mall/internal/data"

	xgin "miopkg/http/gin"

	"github.com/gin-contrib/cors"
)

type MallInterface struct {
	shopBiz       *biz.ShopBiz
	productBiz    *biz.ProductBiz
	collectionBiz *biz.CollectionBiz
	cartBiz       *biz.CartBiz
	orderBiz      *biz.OrderBiz
}

func NewMallInterface(shopBiz *biz.ShopBiz, productBiz *biz.ProductBiz, collectionBiz *biz.CollectionBiz, cartBiz *biz.CartBiz, orderBiz *biz.OrderBiz) *MallInterface {
	return &MallInterface{
		shopBiz:       shopBiz,
		productBiz:    productBiz,
		collectionBiz: collectionBiz,
		cartBiz:       cartBiz,
		orderBiz:      orderBiz,
	}
}

func (engine *Engine) NewHttpServer() error {
	mallRpc := data.NewMallRpcClient()
	userRpc := data.NewUserRpcClient()
	dataInstance := data.NewData(mallRpc, userRpc)

	shopRepo := data.NewShopRepo(dataInstance)
	productRepo := data.NewProductRepo(dataInstance)
	collectionRepo := data.NewCollectionRepo(dataInstance)
	cartRepo := data.NewCartRepo(dataInstance)
	orderRepo := data.NewOrderRepo(dataInstance)

	shopBiz := biz.NewShopBiz(shopRepo)
	productBiz := biz.NewProductBiz(productRepo)
	collectionBiz := biz.NewCollectionBiz(collectionRepo)
	cartBiz := biz.NewCartBiz(cartRepo)
	orderBiz := biz.NewOrderBiz(orderRepo)

	mallInterface := NewMallInterface(shopBiz, productBiz, collectionBiz, cartBiz, orderBiz)

	server := xgin.StdConfig("http").Build()
	server.Use(cors.Default())
	apiv1 := server.Group("/v1/mall")
	{
		apiv1.POST("/shop", mallInterface.CreateShop)
		apiv1.GET("/shop/:id", mallInterface.GetShopById)

		apiv1.POST("/product", mallInterface.CreateProduct)
		apiv1.GET("/product/:id", mallInterface.GetProductById)

		apiv1.POST("/collection", mallInterface.CreateCollection)
		apiv1.GET("/collection/:id", mallInterface.GetCollectionById)
		apiv1.GET("/collection/card/:id", mallInterface.GetCollectionCardById)
		apiv1.POST("/collection/cards", mallInterface.ListCollectionCardByIds)
		apiv1.GET("/collection/shop/:id/:offset/:limit/:by/:order", mallInterface.ListCollectionCardByShopId)
		apiv1.PUT("/collection/sku", mallInterface.UpdateCollectionSku)
		apiv1.PUT("/collection/state", mallInterface.UpdateCollectionState)

		apiv1.POST("/cart/item", mallInterface.CreateCartItem)
		apiv1.GET("/carts/user/:id", mallInterface.ListCartItemByUid)
		apiv1.PUT("/cart/item", mallInterface.UpdateCartItemQuantity)
		// apiv1.DELETE("/cart/item", mallInterface.DeleteCartItem)

	}
	return engine.Serve(server)
}
