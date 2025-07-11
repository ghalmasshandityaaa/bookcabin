package route

func (r *Route) SetupAircraftRoute() {
	r.Log.Info("setting up aircraft routes")

	r.App.Post("/api/check", r.VoucherHandler.Check)
	r.Log.Info("mapped {/api/check, POST} route")

	r.App.Post("/api/generate", r.VoucherHandler.Generate)
	r.Log.Info("mapped {/api/generate, POST} route")
}
