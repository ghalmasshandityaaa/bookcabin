package route

func (r *Route) SetupVoucherRoute() {
	r.Log.Info("setting up voucher routes")

	r.App.Get("/api/aircraft/seats", r.AircraftHander.ListSeats)
	r.Log.Info("mapped {/api/aircraft/seats, GET} route")
}
