package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	fmt.Printf("所有美食：\n- 小面 6元\n- 饭团 7元\n- 香菇滑鸡 12元\n- 小炒肉 15元\n- 黄焖鸡 16元\n- 冒菜 18元\n")
	m:= map [int]string{
		1: "小面  6元",
		2: "饭团  7元",
		3: "香菇滑鸡  12元",
		4: "小炒肉  15元",
		5: "黄焖鸡  16元",
		6: "冒菜  18元",
	}
	fmt.Println("今天要恰的:\n",m[choose()])
}

func choose()int {
	rand.Seed(time.Now().UnixNano())
	ret:= rand.Intn(6)
	return ret

}