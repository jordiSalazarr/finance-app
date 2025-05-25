package api

import (
	"fmt"

	monthlyincome "finances.jordis.golang/application/moves/transactions/commands/monthly-income"
	"github.com/robfig/cron/v3"
)

func (app *App) StartCrons() {
	go func() {
		c := cron.New()
		//for testing purposes should run every minute
		// "0 0 1 * *"
		_, err := c.AddFunc("* * * * *", func() {
			monthlyincome.MonthlyIncomeCommandHandler(app.UsersRepo)
			fmt.Println("Cron job executed")
		})

		if err != nil {
			fmt.Println("Error al agregar el cron job:", err)
			return
		}

		c.Start()

		select {}
	}()

}
