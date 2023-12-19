package internal

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/krispekla/pro-profile-ai-api/internal/handler"
)

func routes(app *Application) *chi.Mux {
	hls := handler.NewHandler(app.Db, app.ErrorLog, app.InfoLog, app.R2Config)
	// mdlw := customMiddleware.NewMiddleware(app.JwtSecret, app.InfoLog, app.ClientError)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	// r.Use(CorsMiddleware())
	r.Use(middleware.Heartbeat("/ping"))
	// r.Use(mdlw.AuthMiddleware())
	r.Route("/api", func(r chi.Router) {
		r.Get("/user", hls.UserDetails())
		// r.Get("/test", hls.GetPackages123())
		r.Get("/characters", hls.GetCharacters())
		r.Post("/characters", hls.CreateCharacter())
		r.Get("/packages", hls.GetPackages())
		r.Get("/packages/:id", hls.GetPackageDetails())
		r.Post("/packages/buy", hls.BuyPackage())
		r.Post("/checkout/sessions", hls.CreateCheckoutSession())
		r.Get("/checkout/sessions", hls.RetrieveCheckoutSession())
		// r.Get("/buckets/all", handlers.GetAllBuckets())
		// r.Get("/image/presigned", handlers.GetPresignedImgUrl())
	})
	return r
}