package db

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertUser(t *testing.T) {
	db := GetDB(t)
	db.InsertUser("李明", 18)
	db.InsertUser("张伟", 19)
	db.InsertUser("王芳", 20)
	db.InsertUser("刘强", 21)
	db.InsertUser("陈静", 22)
	db.InsertUser("杨波", 23)
	db.InsertUser("赵红", 24)
	db.InsertUser("周杰", 25)
	db.InsertUser("吴婷", 26)
	db.InsertUser("郑阳", 27)
	db.InsertUser("孙华", 28)
	db.InsertUser("马超", 29)
	db.InsertUser("朱梅", 30)
	db.InsertUser("胡军", 31)
	db.InsertUser("郭敏", 32)
	db.InsertUser("何平", 33)
	db.InsertUser("高峰", 34)
	db.InsertUser("林娟", 35)
	db.InsertUser("徐亮", 36)
	db.InsertUser("黄海", 37)
	db.InsertUser("魏莉", 38)
	db.InsertUser("王刚", 39)
	db.InsertUser("张丽", 40)
	db.InsertUser("李鑫", 41)
	db.InsertUser("刘洋", 42)
	db.InsertUser("陈阳", 43)
	db.InsertUser("杨帆", 44)

}

func TestData(t *testing.T) {
	names := []string{
		"李明", "张伟", "王芳", "刘强", "陈静", "杨波", "赵红", "周杰", "吴婷", "郑阳",
		"孙华", "马超", "朱梅", "胡军", "郭敏", "何平", "高峰", "林娟", "徐亮", "黄海",
		"魏莉", "王刚", "张丽", "李鑫", "刘洋", "陈阳", "杨帆", "赵亮", "周璐", "吴强",
	}
	for i, name := range names {
		age := 18 + (i % 43)
		fmt.Printf("db.InsertUser(\"%s\", %d)\n", name, age)
	}
}

func TestGetUser(t *testing.T) {
	db := GetDB(t)
	db.GetUser("张三")
}

func TestUpdateUser(t *testing.T) {
	db := GetDB(t)
	db.UpdateUser("张三", 21)
}

func TestDeleteUser(t *testing.T) {
	db := GetDB(t)
	db.DeleteUser("张三")
}

func TestGetAllUsers(t *testing.T) {
	db := GetDB(t)
	users, err := db.GetAllUsers()
	assert.NoError(t, err)
	for _, user := range users {
		fmt.Println(user.Name, user.Age)
	}
}
