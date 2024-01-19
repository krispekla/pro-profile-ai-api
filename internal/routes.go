package internal

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/krispekla/pro-profile-ai-api/internal/handler"
	customMiddleware "github.com/krispekla/pro-profile-ai-api/internal/middleware"
)

func routes(app *Application) *chi.Mux {
	hls := handler.NewHandler(app.Db, app.ErrorLog, app.InfoLog, app.R2Service.Config)
	mdlw := customMiddleware.NewMiddleware(app.Config.JwtSecret, app.InfoLog)
	r := chi.NewRouter()
	pubRoutes := map[string]struct{}{
		"/api/payment/webhook": {},
		"/api/user/webhook":    {},
	}
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(mdlw.CorsMiddleware())
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(mdlw.AuthMiddleware(pubRoutes))
	r.Route("/api", func(r chi.Router) {
		r.Get("/user", hls.UserDetails())
		r.Post("/user/webhook", hls.UserRegistrationWebhook())
		r.Get("/characters", hls.GetCharacters())
		r.Post("/characters", hls.CreateCharacter())
		r.Get("/packages/listing", hls.GetPackageListing())
		r.Get("/packages/generated", hls.GetGeneratedPackages())
		r.Get("/packages/:id", hls.GetPackageDetails())
		r.Post("/packages/buy", hls.BuyPackage())
		r.Get("/order/all", hls.GetAllOrders())
		r.Post("/payment/checkout", hls.CreateCheckoutSession())
		r.Get("/payment/checkout", hls.RetrieveCheckoutSession())
		r.Post("/payment/webhook", hls.StripeWebhookHandler())
		// r.Get("/buckets/all", handlers.GetAllBuckets())
		// r.Get("/image/presigned", handlers.GetPresignedImgUrl())
	})

	r.Route("/admin", func(r chi.Router) {
		// r.Use(mdlw.AdminMiddleware())
		// r.Get("/product/list", hls.GetProductListing())
		// r.Post("/product", hls.CreateProduct())
		// r.Put("/product/:id", hls.UpdateProduct())
		// r.Delete("/product/:id", hls.DeleteProduct())
		// r.Get("/product/coupon", hls.GetCoupons())
		// r.Post("/product/coupon", hls.CreateCoupon())
		// r.Put("/product/coupon/:id", hls.UpdateCoupon())
		// r.Delete("/product/coupon/:id", hls.DeleteCoupon())
	})
	return r
}
