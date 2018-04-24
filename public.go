package main

import (
	"fmt"
	//	"strings"
	"log"
	"os"
	"runtime"
	//	"bytes"
	"errors"
)

//返回指针安全

func RePoint() *int {
	v := 1
	return &v
}

//匿名函数实现的闭包
func Incr() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

//格式化输出数字
func NumPrintf() {
	num := 0x34fd
	//[1]通知printf重复使用第一个操作数num
	// %d %o %x 十进制 八进制 十六进制
	//#表示在数字前添加进制标识
	fmt.Printf("%d %[1]o %#[1]o %[1]x %#[1]x %#[1]X\n", num)
}

func ArrayTest() {
	//不定长度数组
	var arr = [...]int{1, 2, 3, 4, 5}

	for i, _ := range arr {
		fmt.Printf("%d\n", arr[i])
	}
}

func SliceTest() {
	//s := make([]int, 10, 14)
	var s []int
	//由于不确定底层数组是否变更，通常将append调用结果再次赋值给传入append函数的slice
	s = append(s, 2)

	fmt.Printf("size=%d cap=%d\n", len(s), cap(s))

	for _, x := range s {
		fmt.Printf("%d\n", x)
	}

}

func MapTest() {
	maptest := make(map[string]int)
	maptest["id1"] = 1
	maptest["id2"] = 2
	maptest["id3"] = 3

	for key, value := range maptest {
		fmt.Printf("key=%s value=%d\n", key, value)
	}
}

//struct test
type Students struct {
	S_id int
	Name string
}

type Class struct {
	C_id    int
	Student []Students
}

func CreateClass(id int) Class {
	return Class{C_id: id}
}

func AddStudent(id int, name string, class Class) Class {
	class.Student = append(class.Student, Students{id, name})
	return class
}

func RemoveStudent(id int, class Class) Class {
	for i, value := range class.Student {
		if value.S_id == id {
			class.Student[i] = class.Student[len(class.Student)-1]
			class.Student = class.Student[:len(class.Student)-1]
			break
		}
	}
	return class
}
func PrintStudent(class Class) {
	incr := Incr()
	var ret int
	for count := 0; count < len(class.Student); {
		ret = incr()
		for _, value := range class.Student {
			if value.S_id == ret {
				fmt.Printf("id=%d name=%s\n", value.S_id, value.Name)
				count++
				break
			}
		}
	}
}

//nil是一个合法的方法接收者
func (s *Students) SetName(name string) error {
	if s == nil {
		return errors.New("method receiver is nil")
	}
	(*s).Name = name
	return nil
}
func (s Students) GetName() string {
	return s.Name
}

func StructTest() {
	class := CreateClass(1)
	class = AddStudent(1, "name1", class)
	class = AddStudent(2, "name2", class)
	class = AddStudent(3, "name3", class)
	class = AddStudent(4, "name4", class)
	class = AddStudent(5, "name5", class)
	PrintStudent(class)

	class = RemoveStudent(3, class)
	//class = RemoveStudent(4, class)
	PrintStudent(class)
}

func MethodTest() {
	s := Students{1, "aaa"}
	//此时方法接收者形参为T类型指针，实参为T类型,此时编译器默认进行&T的隐式转换
	//另一种方式(&s).SetName("bbbb")
	//编译器默认隐式转换同样适用于*T到T的转换
	s.SetName("bbbb")
	fmt.Printf("name = %s\n", s.GetName())
	//方法变量
	set := s.SetName
	get := s.GetName //此时已绑定该接收者，因此通过set更新接收者对其无效，get返回的是此时绑定接收者s的值
	set("cccc")
	fmt.Printf("name = %s\n", s.GetName())
	fmt.Printf("name = %s\n", get())
	//方法表达式

	fset := (*Students).SetName
	fget := Students.GetName

	fset(&s, "dddd")
	fmt.Printf("name = %s\n", fget(s))

	//nil
	var ps *Students
	err := ps.SetName("cccc")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("name = %s\n", ps.GetName())
}

//错误输出
func ErrorPrint() {
	fmt.Fprintf(os.Stderr, "print error msg\n")
	err := fmt.Errorf("print error msg")
	fmt.Println(err)
	log.Fatalf("print error msg\n")
}

//边长函数
func ArgFunc(args ...int) int {
	sum := 0
	for _, val := range args {
		sum += val
	}
	return sum
}

//defer 延迟调用
//panic 引发运行时错误，默认退出程序
//recover:从运行时错误恢复，正常运行
func HandleCore() {
	type core struct{}

	defer func() {
		switch p := recover(); p {
		case nil:
			break
		case core{}: //捕捉特定类型运行时错误，然后恢复
			fmt.Fprintf(os.Stderr, "core\n")
			func() {
				var buf [4096]byte
				n := runtime.Stack(buf[:], false) //保存当前运行栈信息
				os.Stdout.Write(buf[:n])
			}()
			break
		default:
			panic(p) //其他异常执行重新引发，执行默认处理操作
			break
		}
	}()

	fmt.Println("not core")
	fmt.Println("set core")
	panic(core{}) //引发一个特定类型的运行时错误
	//panic("core exit")
}
