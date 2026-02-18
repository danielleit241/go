package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/danielleit241/pkg/utils"
)

func ReadInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func GetName(prompt string) string {
	for {
		name := ReadInput(prompt)
		if valid, errMsg := utils.ValidateName(name); !valid {
			fmt.Println("Error:", errMsg)
			continue
		}
		return name
	}
}

func GetID(prompt string) string {
	for {
		id := ReadInput(prompt)
		if valid, errMsg := utils.ValidateID(id); !valid {
			fmt.Println("Error:", errMsg)
			continue
		}
		return id
	}
}

func GetGPA(prompt string) string {
	for {
		gpa := ReadInput(prompt)
		if valid, errMsg := utils.ValidateGPA(gpa); !valid {
			fmt.Println("Error:", errMsg)
			continue
		}
		return gpa
	}
}

func GetSalary(prompt string) float64 {
	for {
		salaryStr := ReadInput(prompt)
		if valid, errMsg := utils.ValidateSalary(salaryStr); !valid {
			fmt.Println("Error:", errMsg)
			continue
		}
		salary, _ := ParseFloat(salaryStr)
		return salary
	}
}

func GetDepartment(prompt string) string {
	for {
		department := ReadInput(prompt)
		if valid, errMsg := utils.ValidateDepartment(department); !valid {
			fmt.Println("Error:", errMsg)
			continue
		}
		return department
	}
}

func GetAge(prompt string) int {
	for {
		ageStr := ReadInput(prompt)
		age, err := ParseInt(ageStr)
		if err != nil {
			fmt.Println("Error: Age must be a valid number")
			continue
		}
		if valid, errMsg := utils.ValidateAge(age); !valid {
			fmt.Println("Error:", errMsg)
			continue
		}
		return age
	}
}

func ParseFloat(input string) (float64, error) {
	var value float64
	_, err := fmt.Sscanf(input, "%f", &value)
	return value, err
}

func ParseInt(input string) (int, error) {
	var value int
	_, err := fmt.Sscanf(input, "%d", &value)
	return value, err
}
