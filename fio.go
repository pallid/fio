package fio

import (
	"fmt"
	"strings"
)

// New - создание экземпляра PartsName
func New() *PartsName {
	return &PartsName{}
}

//GetPartsOfFullName выполняет методы getPartsOfFullName и getLastNameAndInitials
func (p *PartsName) GetPartsOfFullName(fullName string) {
	p.setFullName(fullName)
	p.getLastNameAndInitials()
}

// setFullName устанавливает полное ФИО в поле FullName
func (p *PartsName) setFullName(fullName string) {
	p.FullName = fullName
}

// getPartsOfFullName раскладывает полное имя физического лица на составные части - фамилию, имя и отчество.
// Если в конце полного имени встречаются "оглы", "улы", "уулу", "кызы" или "гызы",
// то они также считаются частью отчества.
// Заполняются поля LastName, FirstName, MiddleName
func (p *PartsName) getPartsOfFullName() {
	s := strings.Split(p.FullName, " ")
	if len(s) >= 1 {
		p.LastName = s[0]
	}
	if len(s) >= 2 {
		p.FirstName = s[1]
	}
	if len(s) >= 3 {
		p.MiddleName = s[2]
	}
	if len(s) > 3 {
		if _, ok := additionalParts[s[3]]; ok {
			m := p.MiddleName
			p.MiddleName = fmt.Sprintf("%s %s", m, s[3])
		}
	}
}

// getLastNameAndInitials формирует краткое представление из полного имени физического лица.
// Заполняется поле LastNameAndInitials
func (p *PartsName) getLastNameAndInitials() {
	var lastNameAndInitials string
	if p.LastName == "" {
		p.getPartsOfFullName()
	}
	if p.FirstName == "" {
		lastNameAndInitials = p.LastName
	} else if p.MiddleName == "" {
		lastNameAndInitials = fmt.Sprintf("%s %s.", p.LastName, cut(p.FirstName, 1))
	} else {
		lastNameAndInitials = fmt.Sprintf("%s %s.%s.", p.LastName, cut(p.FirstName, 1), cut(p.MiddleName, 1))
	}
	p.LastNameAndInitials = lastNameAndInitials
}

func cut(text string, limit int) string {
	runes := []rune(text)
	return string(runes[:limit])
}
