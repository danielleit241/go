package utils

import (
	"bufio"
	"fmt"
	"net/mail"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

func UUIDGenerator() string {
	return uuid.New().String()
}

func ReadInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func ValidateRequired(fieldName, value string) (bool, string) {
	if strings.TrimSpace(value) == "" {
		return false, fmt.Sprintf("%s cannot be empty", fieldName)
	}
	return true, ""
}

func ValidateUUID(id string) (bool, string) {
	if ok, errMsg := ValidateRequired("ID", id); !ok {
		return false, errMsg
	}
	if _, err := uuid.Parse(id); err != nil {
		return false, "ID must be a valid UUID"
	}
	return true, ""
}

func ValidateEmail(email string) (bool, string) {
	if ok, errMsg := ValidateRequired("Email", email); !ok {
		return false, errMsg
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return false, "Email is invalid"
	}
	return true, ""
}

func ValidateDate(date string) (bool, string) {
	if ok, errMsg := ValidateRequired("Date", date); !ok {
		return false, errMsg
	}
	if _, err := time.Parse("2006-01-02", date); err != nil {
		return false, "Date must be in format YYYY-MM-DD"
	}
	return true, ""
}

func ValidateBookTitle(title string) (bool, string) {
	return ValidateRequired("Book title", title)
}

func ValidateBookAuthor(author string) (bool, string) {
	return ValidateRequired("Book author", author)
}

func ValidateUserName(name string) (bool, string) {
	return ValidateRequired("User name", name)
}

func ValidateUserEmail(email string) (bool, string) {
	return ValidateEmail(email)
}

func ValidateTransactionBookID(bookID string) (bool, string) {
	return ValidateUUID(bookID)
}

func ValidateTransactionUserID(userID string) (bool, string) {
	return ValidateUUID(userID)
}

func ValidateBorrowDate(borrowDate string) (bool, string) {
	return ValidateDate(borrowDate)
}

func ValidateReturnDate(returnDate string) (bool, string) {
	if strings.TrimSpace(returnDate) == "" {
		return true, ""
	}
	return ValidateDate(returnDate)
}

func GetBookTitle(prompt string) string {
	for {
		title := ReadInput(prompt)
		if ok, errMsg := ValidateBookTitle(title); !ok {
			fmt.Println("Error:", errMsg)
			continue
		}
		return title
	}
}

func GetBookAuthor(prompt string) string {
	for {
		author := ReadInput(prompt)
		if ok, errMsg := ValidateBookAuthor(author); !ok {
			fmt.Println("Error:", errMsg)
			continue
		}
		return author
	}
}

func GetUserName(prompt string) string {
	for {
		name := ReadInput(prompt)
		if ok, errMsg := ValidateUserName(name); !ok {
			fmt.Println("Error:", errMsg)
			continue
		}
		return name
	}
}

func GetUserEmail(prompt string) string {
	for {
		email := ReadInput(prompt)
		if ok, errMsg := ValidateUserEmail(email); !ok {
			fmt.Println("Error:", errMsg)
			continue
		}
		return email
	}
}

func GetID(prompt string) string {
	for {
		id := ReadInput(prompt)
		if ok, errMsg := ValidateUUID(id); !ok {
			fmt.Println("Error:", errMsg)
			continue
		}
		return id
	}
}

func GetBorrowDate(prompt string) string {
	for {
		borrowDate := ReadInput(prompt)
		if ok, errMsg := ValidateBorrowDate(borrowDate); !ok {
			fmt.Println("Error:", errMsg)
			continue
		}
		return borrowDate
	}
}

func GetReturnDate(prompt string) string {
	for {
		returnDate := ReadInput(prompt)
		if ok, errMsg := ValidateReturnDate(returnDate); !ok {
			fmt.Println("Error:", errMsg)
			continue
		}
		return returnDate
	}
}
