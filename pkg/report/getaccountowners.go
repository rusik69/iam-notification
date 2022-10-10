package report

import (
	"encoding/csv"
	"os"
	"strings"
)

func GetAccountOwners(csvFileName string) (map[string]string, error) {
	f, err := os.Open(csvFileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	accountOwners := make(map[string]string)
	skip := true
	for _, record := range records {
		if skip {
			skip = false
			continue
		}
		account := record[0]
		account = strings.Split(account, "(")[1]
		account = strings.Split(account, ")")[0]
		account = strings.TrimSpace(account)
		email := strings.TrimSpace(record[3])
		if account == "" || email == "" {
			continue
		}
		accountOwners[account] = email
	}
	return accountOwners, nil
}
