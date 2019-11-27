package main

import (
	"fmt"

	"github.com/pallid/fio"
)

func main() {
	s := []string{"Иванов Сергей Александрович", "Алиев Ахмед Октай оглы"}
	for _, str := range s {
		p := &fio.PartsName{}
		p.SetFullName(str)
		p.GetPartsOfFullName()
		p.GetLastNameAndInitials()
		fmt.Printf("Устанавливаемое значение: %s\n", str)
		fmt.Printf("Фамилия: %s\n", p.LastName)
		fmt.Printf("Имя: %s\n", p.FirstName)
		fmt.Printf("Отчество: %s\n", p.MiddleName)
		fmt.Printf("ФамилияИО: %s\n", p.LastNameAndInitials)
	}
}
