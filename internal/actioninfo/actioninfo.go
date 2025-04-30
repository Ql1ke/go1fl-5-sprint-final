package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(data string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Printf("Ошибка парсинга: %v", err)
			continue
		}
		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("Ошибка формирования информации: %v", err)
			continue
		}
		fmt.Println(info)
	}
}
