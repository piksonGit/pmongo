package main

import "fmt"

type Ranran struct {
	Mouse string
	Brain string
}
type Qijing struct {
	*Ranran
	Name    string
	Age     int
	Address string
}

func main() {
	var peter *Qijing = &Qijing{
		&Ranran{"嘴巴", "大脑"},
		"亓京", 12, "山东省莱芜市",
	}
	fmt.Println((*peter).Ranran.Brain)

}
