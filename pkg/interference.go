package pkg

import "fmt"

// 导入payload
func ManagePayload() {
	fmt.Println("收集部分指定payload")

}

//设定攻击频率
func SetFrequency() {
	fmt.Println("设定攻击频率")
}

//并发攻击
func ConcurrentAttack(){
	fmt.Println("并发攻击所有ip")
}