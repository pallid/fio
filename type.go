package fio

// PartsName ...
type PartsName struct {
	FullName            string //Полное ФИО, Фамилия Имя Отчество
	LastName            string //Фамилия
	FirstName           string //Имя
	MiddleName          string //Отчество
	LastNameAndInitials string //Фамилия и инициалы, Фамилия И.О.
}

// additionalParts дополнительные части отчества
var additionalParts = map[string]bool{
	"оглы": true,
	"улы":  true,
	"уулу": true,
	"кызы": true,
	"гызы": true,
	"угли": true,
}
