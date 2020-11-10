package structex

import "fmt"

type Address struct {
	City string
	name string
}

//匿名字段
type Human struct {
	id   int
	name string
	int
	string
	*Address
}

type Student struct {
	school string
	*Human
}

func (student *Student) Examination() {
	fmt.Printf("%s Examination at %s \n", student.Human.name, student.school)
}

func (humen *Human) Say() {
	fmt.Printf("%s say hello \n", humen.name)
}

//默认初始化
func NewHuman(id int, name string, code string, age int, city string, aname string) *Human {
	return &Human{
		id:     id,
		name:   name,
		int:    age,
		string: code,
		Address: &Address{
			City: city,
			name: aname,
		},
	}
}

//new 的方式初始化
func NewHuman2(id int, name string, code string, age int, city string, aname string) *Human {
	user := new(Human)
	user.id = id
	user.name = name
	user.int = age
	user.string = code
	user.Address = new(Address)
	user.City = city
	user.Address.name = aname
	return user
}

//默认强制构造方式
func NewHuman3(id int, name string, code string, age int, city string, aname string) *Human {
	var user Human
	user.id = id
	user.name = name
	user.int = age
	user.string = code
	user.Address = new(Address)
	user.City = city
	user.Address.name = aname
	return &user
}
func (user *Human) ToString() string {
	return fmt.Sprintf(" user.id = %v ,user.name = %v ,user.int = %v,user.string = %v,user.Address.City = %v,user.Address.name = %v \n", user.id, user.name, user.int, user.string, user.City, user.Address.name)
}

func TestNM() {

	user1 := NewHuman(11, "ccc", "code", 19, "gz", "江南")
	user2 := NewHuman2(12, "cc2c", "c2ode", 219, "gz", "江南1")
	user3 := NewHuman3(13, "cc32c", "c23ode", 21239, "gz", "江南3")

	fmt.Println("user1: \n", user1.ToString())
	fmt.Println("user2: \n", user2.ToString())
	fmt.Println("user3: \n", user3.ToString())

}
