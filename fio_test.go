package fio

import "testing"

var fioTests = []struct {
	input               string
	firstName           string
	middleName          string
	lastName            string
	fullName            string
	lastNameAndInitials string
}{
	{
		input:               "Иванов Сергей Александрович",
		firstName:           "Сергей",
		middleName:          "Александрович",
		lastName:            "Иванов",
		fullName:            "Иванов Сергей Александрович",
		lastNameAndInitials: "Иванов С.А.",
	},
	{
		input:               "Алиев Ахмед Октай оглы",
		firstName:           "Ахмед",
		middleName:          "Октай оглы",
		lastName:            "Алиев",
		fullName:            "Алиев Ахмед Октай оглы",
		lastNameAndInitials: "Алиев А.О.",
	},
	{
		input:               "Степанов Андрей",
		firstName:           "Андрей",
		middleName:          "xxx",
		lastName:            "Степанов",
		fullName:            "Степанов Андрей",
		lastNameAndInitials: "Степанов А.",
	},
	{
		input:               "Афанасенко",
		firstName:           "",
		middleName:          "",
		lastName:            "Афанасенко",
		fullName:            "Афанасенко",
		lastNameAndInitials: "Афанасенко",
	},
}

func TestSetFullName(t *testing.T) {
	for _, tt := range fioTests {
		p := &PartsName{}
		p.SetFullName(tt.input)
		if p.FullName != tt.fullName {
			t.Fatalf("SetFullName. Expected [%v], Actual [%v]", tt.fullName, p.FullName)
		}
	}
}

func BenchmarkSetFullName(b *testing.B) {
	p := &PartsName{}
	for i := 0; i < b.N; i++ {
		for _, tt := range fioTests {
			p.SetFullName(tt.input)
		}
	}
}

func TestGetPartsOfFullName(t *testing.T) {
	for _, tt := range fioTests {
		p := &PartsName{FullName: tt.input}
		p.GetPartsOfFullName()
		if p.LastName != tt.lastName {
			t.Fatalf("GetPartsOfFullName. Expected [%v], Actual [%v]", tt.lastName, p.LastName)
		}
		if p.FirstName != tt.firstName {
			t.Fatalf("GetPartsOfFullName. Expected [%v], Actual [%v]", tt.firstName, p.FirstName)
		}
		if p.MiddleName != tt.middleName {
			t.Fatalf("GetPartsOfFullName. Expected [%v], Actual [%v]", tt.middleName, p.MiddleName)
		}
	}
}

func BenchmarkGetPartsOfFullName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range fioTests {
			p := &PartsName{FullName: tt.input}
			p.GetPartsOfFullName()
		}
	}
}

func TestGetLastNameAndInitials(t *testing.T) {
	for _, tt := range fioTests {
		p := &PartsName{FullName: tt.input}
		p.GetLastNameAndInitials()
		if p.LastNameAndInitials != tt.lastNameAndInitials {
			t.Fatalf("GetLastNameAndInitials. Expected [%v], Actual [%v]", tt.lastNameAndInitials, p.LastNameAndInitials)
		}
	}
}

func BenchmarkGetLastNameAndInitials(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range fioTests {
			p := &PartsName{FullName: tt.input}
			p.GetLastNameAndInitials()
		}
	}
}
