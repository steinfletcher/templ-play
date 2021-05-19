package todos

type Todo struct {
	ID   int64
	Name string
	Date string
}

func Get() []Todo {
	return []Todo{
		{
			ID:   1,
			Name: "Buy milk",
			Date: "2021-12-11T10:00:000Z",
		},
		{
			ID:   2,
			Name: "Some title",
			Date: "2021-12-11T10:00:000Z",
		},
	}
}
