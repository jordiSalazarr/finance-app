package monthlyincome

import (
	"fmt"

	domainUsers "finances.jordis.golang/domain/members/users"
)

func MonthlyIncomeCommandHandler(userRepository domainUsers.UserRepository) error {

	users, err := userRepository.GetAll()
	if err != nil {
		fmt.Println("Error al obtener los usuarios:", err)
		return err
	}
	for _, user := range users {
		err = userRepository.UpdateCurrentBalance(user.Pk.Val, user.MonthlyIncome.Val)
		if err != nil {
			fmt.Println("Error al guardar el usuario:", err)
			return err
		}
		fmt.Println("Usuario actualizado:", user.Username.Val)

	}
	return nil
}
