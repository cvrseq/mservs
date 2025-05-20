package model


type Employee struct { 
	ID			int			`json:"id"`
	FirstName	string		`json:"first name"`
	SecondName	string		`json:"second name"`
	LastName	string		`json:"last name"`
	Grade		int			`json:"grade"`
}

