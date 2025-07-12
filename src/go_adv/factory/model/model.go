package model

//student小写时其他包无法访问
type student struct{
	name string
	age int
}

func NewStudent(n string , a int) *student {
	return &student{
		name:n,
		age:a,
	}
}

func (s *student) GetAge() int{
	return s.age
}

