package api

func (app *App) StartCrons() {
	// go func() {
	// 	c := cron.New()
	// 	//for testing purposes should run every minute
	// 	// "0 0 1 * *"
	// 	_, err := c.AddFunc("* * * * *", func() {
	// 		monthlyincome.MonthlyIncomeCommandHandler(app.UsersRepo)
	// 		app.Logger.Info("Cron job executed: Monthly income command handler called")
	// 	})

	// 	if err != nil {
	// 		fmt.Println("Error al agregar el cron job:", err)
	// 		return
	// 	}

	// 	c.Start()

	// 	select {}
	// }()

}
