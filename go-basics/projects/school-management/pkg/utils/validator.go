package utils

var VALID_DEPARTMENTS = map[string]bool{
	"Xavalo": true,
	"Hola":   true,
	"Fuda":   true,
	"Hovilo": true,
}

var VALID_ID_PREFIXES = map[string]bool{
	"SE": true,
	"SS": true,
	"SA": true,
}

func ValidateID(id string) (bool, string) {
	prefix := id[:2]
	if !VALID_ID_PREFIXES[prefix] {
		return false, "ID must start with SE, SS, or SA"
	}
	if len(id) < 5 {
		return false, "ID must be at least 5 characters long"
	}
	return true, ""
}

func ValidateName(name string) (bool, string) {
	if len(name) == 0 {
		return false, "Name cannot be empty"
	}
	return true, ""
}

func ValidateGPA(strGPA string) (bool, string) {
	if len(strGPA) != 4 || strGPA[1] != '.' {
		return false, "GPA must be in format X.XX"
	}
	return true, ""
}

func ValidateSalary(strSalary string) (bool, string) {
	if len(strSalary) == 0 {
		return false, "Salary cannot be empty"
	}
	return true, ""
}

func ValidateDepartment(department string) (bool, string) {
	if len(department) == 0 {
		return false, "Department cannot be empty"
	}
	if !VALID_DEPARTMENTS[department] {
		return false, "Invalid department"
	}
	return true, ""
}

func ValidateAge(age int) (bool, string) {
	if age < 18 || age > 65 {
		return false, "Age must be between 18 and 65"
	}
	return true, ""
}
