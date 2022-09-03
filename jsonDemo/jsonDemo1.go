package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//声明技能结构
	type Skill struct {
		Name  string
		Level int
	}
	//声明角色结构
	type Actor struct {
		Name   string
		Age    int
		Skills []Skill
	}
	//填充基本角色数据
	a := Actor{
		Name: "cowboy",
		Age:  37,
		Skills: []Skill{
			{
				Name:  "Rollandroll",
				Level: 1,
			},
			{
				Name:  "Flashyourdogeye",
				Level: 2,
			},
			{
				Name:  "TimetohaveLunch",
				Level: 3,
			},
		},
	}
	result, err := json.Marshal(a)

	if err == nil {
		fmt.Println(string(result))
	} else {
		fmt.Println(err)
	}
	fmt.Printf("------------%s---------------\n", "Unmarshal")
	var actorData Actor
	unmarshalValErr := json.Unmarshal(result, &actorData)
	if unmarshalValErr == nil {
		fmt.Println(actorData)
	} else {
		fmt.Println(err)
	}
}
