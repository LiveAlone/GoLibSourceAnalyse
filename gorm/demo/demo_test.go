package demo

import (
	"fmt"
	"testing"
)

func TestCreate(t *testing.T) {

	db, err := buildDb(hxxMis)
	if err != nil {
		fmt.Println(err)
		return
	}

	model := &DBMonitorTask{
		DBId:     1,
		DBName:   "test",
		Schedule: 20,
		Tables:   "*",
		Action:   1,
		Status:   1,
	}

	// 查询返回结果
	result := db.Table(DBMonitorTaskName).Create(model)
	fmt.Println(model.Id)            // 插入id
	fmt.Println(result.Error)        // error
	fmt.Println(result.RowsAffected) // effect rows

	// db.Select("Name", "Age", "CreatedAt").Create(&user)
	// INSERT INTO `users` (`name`,`age`,`created_at`) VALUES ("jinzhu", 18, "2020-07-04 11:05:21.775")
	//db.Omit("Name", "Age", "CreatedAt").Create(&user)
	// INSERT INTO `users` (`birthday`,`updated_at`) VALUES ("2020-01-01 00:00:00.000", "2020-07-04 11:05:21.775")

}

func TestDbAction(t *testing.T) {
	db, err := buildDb(hxxMis)
	if err != nil {
		fmt.Println(err)
		return
	}

	// create table
	//db.Table(DBMonitorTaskName).Create(&DBMonitorTask{
	//	DBId:     1,
	//	DBName:   "test",
	//	Schedule: 20,
	//	Tables:   "*",
	//	Action:   1,
	//	Status:   1,
	//})

	// read
	var task DBMonitorTask
	//db.Table(DBMonitorTaskName).First(&task)
	db.Table(DBMonitorTaskName).First(&task, 5)
	fmt.Println(task.CreatedAt)

	//update
	//db.Table(DBMonitorTaskName).Model(&task).Update("db_name", "hello")
	//db.Table(DBMonitorTaskName).Model(&task).Updates(&DBMonitorTask{DBName: "hello3", Status: 4})
	//db.Table(DBMonitorTaskName).Model(&task).Updates(map[string]interface{}{"DBName": "world3", "status": 6})

	//delete
	//db.Table(DBMonitorTaskName).Delete(&task)
}
