package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type User struct {
	Name   string
	Age    int
	Gender string
}

// 给 User 实现 String 方法
func (user *User) String() string {
	buffer := bytes.NewBufferString("This is ")
	buffer.WriteString(user.Name + ", ")
	if user.Gender == "male" {
		buffer.WriteString("He ")
	} else {
		buffer.WriteString("She ")
	}

	buffer.WriteString("is ")
	buffer.WriteString(strconv.Itoa(user.Age))
	buffer.WriteString(" years old.")

	return buffer.String()
}

func main() {
	// Fscan, Sscan
	// Fprint, Sprint, Errorf
	u := User{
		Name:   "",
		Age:    0,
		Gender: "",
	}

	count, _ := fmt.Sscan("zenowang 18 male", &u.Name, &u.Age, &u.Gender)
	fmt.Println("args num:", count, "userinfo:", u.Name, u.Age, u.Gender)

	// Stringer接口实现 ToString 功能
	uPrt := &u
	fmt.Println(uPrt)
}
