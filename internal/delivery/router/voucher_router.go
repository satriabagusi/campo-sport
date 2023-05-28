package router

import (
	"github.com/gin-gonic/gin"
	"github.com/satriabagusi/campo-sport/internal/delivery/handler"
	"github.com/satriabagusi/campo-sport/internal/delivery/middleware"
	"github.com/satriabagusi/campo-sport/internal/usecase"
)

type VoucherRouter struct {
	voucherHandler handler.VoucherHandler
	publicRoute    *gin.RouterGroup
}

func (v *VoucherRouter) SetupRouter() {
	voucher := v.publicRoute.Group("/voucher")
	{
		voucher.Use(middleware.Auth())
		voucher.GET("/:id", v.voucherHandler.FindVoucherByID)
		voucher.GET("/search/", v.voucherHandler.FindVoucherByVoucherCode)
		voucher.GET("/", v.voucherHandler.GetAllVoucher)
	}

	admin := v.publicRoute.Group("/admin")
	{
		admin.Use(middleware.Auth())
		admin.POST("/", v.voucherHandler.InsertVoucher)
		admin.PUT("/", v.voucherHandler.UpdateVoucher)
		admin.DELETE("/:id", v.voucherHandler.DeleteVoucher)
	}

}

func NewVoucherRouter(publicRoute *gin.RouterGroup, voucherUsecase usecase.VoucherUsecase) {
	voucherHandler := handler.NewVoucherHandler(voucherUsecase)
	rt := VoucherRouter{
		voucherHandler,
		publicRoute,
	}
	rt.SetupRouter()
}
