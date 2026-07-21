package main

import "fmt"

type Human interface {
	Say() string
}

type Person struct {
	Name string
}

// 値	を更新処理の場合の記載の仕方
func (p *Person) Say() string {
	p.Name = "Mr." + p.Name
	return p.Name
}

func DriveCar(human Human) {
	if human.Say() == "Mr.Mike" {
		fmt.Println("Run")
	} else {
		fmt.Println("Get out")
	}
}

// インターフェースに関しての講義
func main() {
	var Mike Human = &Person{"Mike"}
	// Mike.Say()
	DriveCar(Mike)
}
