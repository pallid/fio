
Пример использования

```go
package main

import (
	"fmt"

	"github.com/pallid/fio"
)

func main() {
	s := []string{"Иванов Сергей Александрович", "Алиев Ахмед Октай оглы"}
	for _, str := range s {
		p := fio.New()
		p.GetPartsOfFullName(str)
		fmt.Printf("Устанавливаемое значение: %s\n", str)
		fmt.Printf("Фамилия: %s\n", p.LastName)
		fmt.Printf("Имя: %s\n", p.FirstName)
		fmt.Printf("Отчество: %s\n", p.MiddleName)
		fmt.Printf("ФамилияИО: %s\n", p.LastNameAndInitials)
	}
}
```

```
Устанавливаемое значение: Иванов Сергей Александрович
Фамилия: Иванов
Имя: Сергей
Отчество: Александрович
ФамилияИО: Иванов С.А.
Устанавливаемое значение: Алиев Ахмед Октай оглы
Фамилия: Алиев
Имя: Ахмед
Отчество: Октай оглы
ФамилияИО: Алиев А.О.
```