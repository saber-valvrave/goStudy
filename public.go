package main

import (
	"fmt"
	//	"strings"
	"os"
	"log"
	"runtime"
//	"bytes"
)

//返回指针安全

func RePoint() *int {
	v := 1
	return &v
}

func Incr() func () int {
	x := 0
	return func () int {
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
	flag := true
	for count:= 0; count < len(class.Student); {
		for _, value := range class.Student {
			if flag {
				ret = incr()
			}	
			if value.S_id == ret {
				fmt.Printf("id=%d name=%s\n", value.S_id, value.Name)
				count++
				flag = true
				continue
			}
			flag = false
		}
	}
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
	class = RemoveStudent(4, class)
	PrintStudent(class)
}

func ErrorPrint() {
	fmt.Fprintf(os.Stderr, "print error msg\n")
	err := fmt.Errorf("print error msg")
	fmt.Println(err)
	log.Fatalf("print error msg\n")
}

func ArgFunc(args ...int) int{
	sum := 0
	for _,val := range args {
		sum += val	
	}
	return sum
}

func HandleCore(){
	type core struct{}

	defer func(){
		switch p:=recover(); p {
			case nil:
				break
			case core{}:
				fmt.Fprintf(os.Stderr, "core\n")
				func(){
					var buf [4096]byte
					n := runtime.Stack(buf[:], false)
					os.Stdout.Write(buf[:n])
				}()
				break
			default:
				panic(p)
				break
		}
	}()

	fmt.Println("not core")
	fmt.Println("set core")
	panic(core{})
	//panic("core exit")
}
