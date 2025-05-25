package api

import (
	"fmt"

	monthlyincome "finances.jordis.golang/application/moves/transactions/commands/monthly-income"
	"github.com/robfig/cron/v3"
)

func (app *App) StartCrons() {
	go func() {
		c := cron.New()

		_, err := c.AddFunc("0 3 1 * *", func() {
			monthlyincome.MonthlyIncomeCommandHandler(app.UsersRepo)
			app.Logger.Info("Cron job executed: Monthly income command handler called")
		})

		if err != nil {
			fmt.Println("Error al agregar el cron job:", err)
			return
		}

		c.Start()

		select {}
	}()

}
