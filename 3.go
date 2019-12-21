package main

import (
	"fmt"
	"math/rand"
	"time"
)

type  qizi struct {
	HP 	int
	MP 	int
	MaxMP int
	DPS 	float32
	skill 	string
}

func main()  {
	Awin:=0.0
	Bwin:=0.0
	var A qizi
		A.HP = 100
		A.MP = 0
		A.MaxMP = 0
		A.DPS = 10
		A.skill = "剑士羁绊：每次攻击有50%概率触发连击（又一次攻击）"
	var B qizi
		B.HP = 300
		B.MP = 0
		B.MaxMP = 50
		B.DPS = 20
		B.skill = "蓝量达到50自动触发、降低对方10%当前攻击力"

		var i,p int
		fmt.Println("对战开始:\n")
	for p=1;p<=10 ;p++ {
		for i = 1; i <= 5; i++ {
			fmt.Println("进入A的回合")
			n := A.Agongji()
			B.HP -= 10 * n
			if B.Bdeath() {
				fmt.Println("B已死亡,A获胜")
				Awin++
				break
			}
			fmt.Println("进入B的回合")
			m := B.Bgongji()
			A.HP -= 20
			if m {
				A.DPS *= 0.9
			}
			if A.Adeath() {
				fmt.Println("A已死亡,B获胜")
				Bwin++
				break
			}
		}
		for i = 1; i <= 5; i++ {
			fmt.Println("进入B的回合")
			m := B.Bgongji()
			A.HP -= 20
			if m {
				A.DPS *= 0.9
			}
			if A.Adeath() {
				fmt.Println("A已死亡,B获胜")
				Bwin++
				break
			}
			fmt.Println("进入A的回合")
			n := A.Agongji()
			B.HP -= 10 * n
			if B.Bdeath() {
				fmt.Println("B已死亡,A获胜")
				Awin++
				break
			}
		}
	}
	fmt.Println("A的胜率为",Awin*5,"%")
	fmt.Println("B的胜率为",Bwin*5,"%")
}

func (A *qizi) Agongji() int{
	rand.Seed(time.Now().UnixNano())
	n:= rand.Intn(2)
	x := 0
	if n==1 {
		fmt.Println("A攻击了一次")
		x++
	}else {
		fmt.Println("A攻击了一次")
		x++
		A.Agongji()
	}
	return x
}
func (B *qizi) Bgongji() bool {
	B.MP+=10
	if B.MP==50{
		fmt.Println("B发动了技能")
		B.MP=0
		return true
	}else {
		return false
	}
}
func (A *qizi) Adeath() bool{
	if A.HP<=0{
		return true
	}else {
		return false
	}
}
func (B *qizi) Bdeath() bool{
	if B.HP<=0{
		return true
	}else {
		return false
	}
}