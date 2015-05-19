package models

func CreateModel() []Phones {

	var persons []Phones

	p1 := Phones{
		Id: 1,
		Phone: "12-45",
		Person: "Иванов Иван",
		Street: "ул. Пушкина",
		Building: "12",
		Appartment: "21",
	}

	p2 := Phones{
		Id: 2,
		Phone: "89-23",
		Person: "Петров Петр",
		Street: "ул. Кукушкина",
		Building: "45",
		Appartment: "67",
	}

	p3 := Phones{
		Id: 3,
		Phone: "12-96",
		Person: "Сидоров Сергей",
		Street: "ул. Зайцева",
		Building: "56",
		Appartment: "1",
	}

	p4 := Phones{
		Id: 4,
		Phone: "15-45",
		Person: "Смирнова Анна",
		Street: "ул. Артема",
		Building: "89",
		Appartment: "3",
	}

	p5 := Phones{
		Id: 5,
		Phone: "14-47",
		Person: "Иванова Светлана",
		Street: "ул. Гагарина",
		Building: "11",
		Appartment: "34",
	}

	persons = append(persons, p1)
	persons = append(persons, p2)
	persons = append(persons, p3)
	persons = append(persons, p4)
	persons = append(persons, p5)

	return persons
}