package actioninfo

import (
	"fmt"
	"log"
)

// DataParser описывает разбор данных активности и формирование отчёта.
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Info выводит отчёты об активностях из dataset.
// Ошибки отдельных записей логируются и не прерывают обработку остальных.
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Printf("не удалось разобрать запись %q: %v", data, err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("не удалось сформировать отчёт для записи %q: %v", data, err)
			continue
		}

		fmt.Println(info)
	}
}
