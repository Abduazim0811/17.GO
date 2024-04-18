// ### Beriglan `employees.json` faylidan quyidagi amaliyotlarni bajaring:
// * `Employee` struct slice ga ma'lumotlarni `Unmarshal` qilib oling
// * `id` = 3 berilgan employee ni `age` fieldini 50 ga o'zgartiring
// * Yangi `id = 6` bogan ihtiyoriy employee yarating
// * `Employee` slice ga yangi employee ni qoshing
// * Q'oshilgan va o'zgartirilgan employee la bilan Marshal qiling yangi `employees_new.json` file ga yozing
// * Code ni `github` ga joylang
package main

import (
	"encoding/json"
	"log"
	"os"
)

type Employee struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Position string `json:"position"`
}

func main() {
	file, err := os.Open("/home/abduazim/Projects/Golang/NT_Homeworks/17.GO/employees.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var employees []Employee

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&employees)
	if err != nil {
		log.Fatal(err)
	}

	for key, value := range employees {
		if value.ID == 3 {
			employees[key].Age = 50
			break
		}
	}

	newEmployee := Employee{ID: 6, Name: "Abduazim Yusufov", Age: 19, Position: "Software Engineer"}
	employees = append(employees, newEmployee)

	jsonBytes, err := json.MarshalIndent(employees, "", "   ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.Remove("employees_new.json")
	if err != nil && !os.IsNotExist(err) {
		log.Fatal(err)
	}

	newFile, err := os.Create("employees_new.json")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	_, err = newFile.Write(jsonBytes)
	if err != nil {
		log.Fatal(err)
	}
}

