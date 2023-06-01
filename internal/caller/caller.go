package caller

import "fmt"

func Init() {
	// 初始化数据库
	if err := InitDB(); err != nil {
		fmt.Println(fmt.Sprintf("database err, err=%v", err))
	}
}
