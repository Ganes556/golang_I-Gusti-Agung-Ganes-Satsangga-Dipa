package main

import (
	"fmt"
	"log"
)

type Student struct {
	name  []string
	score []float32
}

func (s *Student) AddStudent(name string, score float32) {
	s.name = append(s.name, name)
	s.score = append(s.score, score)
}

func (s *Student) GetMinMax() (map[string]interface{}) {
	var minMax = make(map[string]interface{})

	minMax["max"] = []interface{}{s.name[0], s.score[0]}
	minMax["min"] = []interface{}{s.name[0], s.score[0]}
	for i, v := range s.score {
		if minMax["max"].([]interface{})[1].(float32) < v {
			minMax["max"] = []interface{}{s.name[i], v}
		}
		if minMax["min"].([]interface{})[1].(float32) > v {
			minMax["min"] = []interface{}{s.name[i], v}
		}
	}
	return minMax
}

func (s *Student) GetAvg() float32{
	var sum float32
	for _, v := range s.score{
		sum +=v
	}
	return sum/float32(len(s.score))
}

func (s *Student) showOuput() {
	var (
		avg = s.GetAvg()
		minMax = s.GetMinMax()
		min = minMax["min"].([]interface{})
		max = minMax["max"].([]interface{})

	)
	fmt.Println("\nOutput:")
	fmt.Printf("Average Score : %v\n", avg)
	fmt.Printf("Min Score of Students : %s (%v)\n", min[0].(string), min[1].(float32))
	fmt.Printf("Max Score of Students : %s (%v)\n", max[0].(string), max[1].(float32))
}


func main() {
	students := Student{}
	fmt.Println("Input:")
	for i:= 1; i <= 5; i++ {
		var (
			name string
			score float32
			index = i
		)
		fmt.Printf("Input %d Student's Name ",index)
		fmt.Scan(&name)
		fmt.Printf("Input %d Student's Score ",index)
		fmt.Scan(&score)
		if score < 0 {
			log.Fatal("Invalid Score!")
		}
		students.AddStudent(name, score)
	}
	
	students.showOuput()

}