package structex

import (
	"encoding/json"
	"fmt"
	"reflect"
)

//在定义struct成员时候区分大小写，若首字母大写则该成员为公有成员(对外可见)，否则是私有成员(对外不可见)。
//go中的结构体内存布局和c结构体布局类似，每个成员的内存分布是连续的，
type People struct {
	id   int    "tag 说明，可以通过反射获取"
	Name string `json:"name"`
	Code string `json:"code"`
}

//匿名成员
type Man struct {
	People
	Age  int
	Sex  int
	Name string "继承，与父级一样的名称，那会优先最近的"
}

//返回新的对象，对象辅助使用：
func NewPeople(id int, code string, name string) *People {
	//这个就类似构造函数了
	return &People{id: id, Code: code, Name: name}
}

//构造函数,传引用既可以原地初始化
func (self *Man) ManInit(id int, code string, name string, age int, sex int) {
	self.id = id
	self.Code = code
	self.Name = name
	self.Age = age
	self.Sex = sex
}

//通过new 创建对象在赋值。
func NewPeopleA(id int, code string, name string) *People {
	var p *People = new(People)
	p.Name = name
	p.id = id
	p.Code = code
	return p
}

//直传，实验与传地址的，没传地址。修改是还是值传递
func ChangeNameA(p People) {
	newname := fmt.Sprintf("new B %s ", p.Name)
	fmt.Printf("ChangeNameA object :%v,new name:%s \n", p, newname)
	p.Name = newname
	fmt.Printf("ChangeNameA side people:%v \n", p)
}

//传地址
func ChangeNameB(p *People) {
	newname := fmt.Sprintf("new A %s ", (*p).Name)
	fmt.Printf("ChangeNameB object :%v,new name:%s \n", p, newname)
	(*p).Name = newname
	fmt.Printf("ChangeNameB inside people:%v \n", p)

}

func (p *People) ToJson() string {
	data, err := json.Marshal(p)
	if err != nil {
		return fmt.Sprintf("json encode failed err:%v\n", err)
	}
	return (string(data)) // {"name":"new A name1 ","code":"code1"}
}

//测试属性的就近原则
func TestProperty() {
	var man = new(Man)
	man.Name = "Nmaess" //优先选择man中的name
	man.id = 12
	fmt.Printf("p.name:%s,m.name:%s \n", man.People.Name, man.Name) // 与被继承一样的属性，会就近使用 ,Nmaess
	fmt.Printf("p.id:%d,m.id:%d \n", man.People.id, man.id)         // 继承的属性赋值的是一致 12,12
}

//测试结构的
func TestAddrs() {
	var man *Man = new(Man)
	man.ManInit(1, "code1", "name1", 19, 1)
	fmt.Printf("man.Name %p\n", &man.Name)
	fmt.Printf("man.Age %p\n", &man.Age)
	fmt.Printf("man.Code %p\n", &man.Code)
	fmt.Printf("man.Sex %p\n", &man.Sex)

	typ := reflect.TypeOf(Man{})
	fmt.Printf("Man is %d bytes long\n", typ.Size())
	// We can run through the fields in the structure in order
	n := typ.NumField()
	for i := 0; i < n; i++ {
		field := typ.Field(i)
		fmt.Printf("%s at offset %v, size=%d, align=%d\n",
			field.Name, field.Offset, field.Type.Size(),
			field.Type.Align())
	}
}
