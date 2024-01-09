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
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(mdlw.CorsMiddleware())
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(mdlw.AuthMiddleware())
	r.Route("/api", func(r chi.Router) {
		r.Get("/user", hls.UserDetails())
		// r.Get("/test", hls.GetPackages123())
		r.Get("/characters", hls.GetCharacters())
		r.Post("/characters", hls.CreateCharacter())
		r.Get("/packages/listing", hls.GetPackageListing())
		r.Get("/packages/generated", hls.GetGeneratedPackages())
		r.Get("/packages/:id", hls.GetPackageDetails())
		r.Post("/packages/buy", hls.BuyPackage())
		r.Get("/order/all", hls.GetAllOrders())
		r.Post("/checkout/sessions", hls.CreateCheckoutSession())
		r.Get("/checkout/sessions", hls.RetrieveCheckoutSession())
		// r.Get("/buckets/all", handlers.GetAllBuckets())
		// r.Get("/image/presigned", handlers.GetPresignedImgUrl())
	})
	return r
}
