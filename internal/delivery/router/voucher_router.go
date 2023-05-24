package router

import (
	"github.com/gin-gonic/gin"
	"github.com/satriabagusi/campo-sport/internal/delivery/handler"
	"github.com/satriabagusi/campo-sport/internal/usecase"
)

type VoucherRouter struct {
	voucherHandler handler.VoucherHandler
	publicRoute    *gin.RouterGroup
}

func (v *VoucherRouter) SetupRouter() {
	voucher := v.publicRoute.Group("/voucher")
	{
		//voucher.Use(middleware.Authentication())
		voucher.POST("/", v.voucherHandler.InsertVoucher)
		voucher.PUT("/:id", v.voucherHandler.UpdateVoucher)
		voucher.DELETE("/:id", v.voucherHandler.DeleteVoucher)
		voucher.GET("/:id", v.voucherHandler.FindVoucherByID)
		voucher.GET("/voucher", v.voucherHandler.FindVoucherByVoucherCode)
		voucher.GET("/", v.voucherHandler.GetAllVoucher)
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
