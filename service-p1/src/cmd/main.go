package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/vazy1/reference-service/internal/parser"
)

func main() {
	csvFlag := flag.String("csv", "", "Путь к CSV файлу с данными воздушных судов")
	xmlFlag := flag.String("xml", "", "Путь к XML файлу с данными воздушных судов")
	urlFlag := flag.String("url", "", "URL веб-страницы с данными воздушных судов")
	outputFlag := flag.String("output", "", "Путь для сохранения результатов в JSON формате")
	flag.Parse()

	aircraftParser := parser.NewAircraftDataParser()
	var results []interface{}
	var err error

	if *csvFlag != "" {
		log.Printf("Парсинг CSV файла: %s", *csvFlag)
		results, err = aircraftParser.ParseCSVFile(*csvFlag)
		if err != nil {
			log.Fatalf("Ошибка парсинга CSV: %v", err)
		}
	} else if *xmlFlag != "" {
		log.Printf("Парсинг XML файла: %s", *xmlFlag)
		results, err = aircraftParser.ParseXMLFile(*xmlFlag)
		if err != nil {
			log.Fatalf("Ошибка парсинга XML: %v", err)
		}
	} else if *urlFlag != "" {
		log.Printf("Парсинг веб-страницы: %s", *urlFlag)
		results, err = aircraftParser.ParseWebsiteData(*urlFlag)
		if err != nil {
			log.Fatalf("Ошибка парсинга веб-страницы: %v", err)
		}
	} else {
		flag.Usage()
		os.Exit(1)
	}

	log.Printf("Найдено записей: %d", len(results))

	if *outputFlag != "" {
		jsonData, err := json.MarshalIndent(results, "", "  ")
		if err != nil {
			log.Fatalf("Ошибка сериализации JSON: %v", err)
		}

		if err := os.WriteFile(*outputFlag, jsonData, 0644); err != nil {
			log.Fatalf("Ошибка сохранения файла: %v", err)
		}

		log.Printf("Результаты сохранены в файл: %s", *outputFlag)
	} else {
		count := 5
		if len(results) < count {
			count = len(results)
		}
		
		fmt.Println("Примеры результатов:")
		for i := 0; i < count; i++ {
			jsonData, _ := json.MarshalIndent(results[i], "", "  ")
			fmt.Println(string(jsonData))
		}
	}
} 