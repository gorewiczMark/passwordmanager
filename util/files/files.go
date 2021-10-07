package files

import (
	"encoding/csv"
	"fmt"
	"os"
	"passwordmanager/util/validations"
)

// GetPasswordsFromFile reads the passwords for a csv file
func GetPasswordsFromFile(filename string) [][]string {
	file, err := os.Open(filename)
	if os.IsNotExist(err) {
		return nil
	}
	validations.Check(err)
	defer file.Close()

	r := csv.NewReader(file)
	r.Read()                  //read head of csv file and toss
	lines, err := r.ReadAll() //read rest of file
	validations.Check(err)
	return lines
}

// WritePasswordsToFile - writes new password to file
func WritePasswordsToFile(filename, newEntry string) {
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile(filename, options, os.FileMode(0600))
	validations.Check(err)
	_, err = fmt.Fprintln(file, newEntry)
	validations.Check(err)
	err = file.Close()
	validations.Check(err)
}
