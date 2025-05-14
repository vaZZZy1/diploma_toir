package parser

import (
	"encoding/csv"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/vazy1/reference-service/internal/model"
)

// AircraftDataParser представляет парсер данных о воздушных судах
type AircraftDataParser struct {
	httpClient *http.Client
}

// NewAircraftDataParser создает новый экземпляр парсера данных о воздушных судах
func NewAircraftDataParser() *AircraftDataParser {
	return &AircraftDataParser{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// ParseCSVFile парсит CSV файл с данными воздушных судов
func (p *AircraftDataParser) ParseCSVFile(filePath string) ([]model.Aircraft, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия файла: %w", err)
	}
	defer file.Close()

	return p.parseCSV(file)
}

// ParseCSVData парсит данные из CSV строки
func (p *AircraftDataParser) ParseCSVData(data string) ([]model.Aircraft, error) {
	reader := strings.NewReader(data)
	return p.parseCSV(reader)
}

// ParseXMLFile парсит XML файл с данными воздушных судов
func (p *AircraftDataParser) ParseXMLFile(filePath string) ([]model.Aircraft, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия файла: %w", err)
	}
	defer file.Close()

	return p.parseXML(file)
}

// ParseWebsiteData парсит данные с веб-сайта
func (p *AircraftDataParser) ParseWebsiteData(url string) ([]model.Aircraft, error) {
	resp, err := p.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("ошибка HTTP запроса: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("неуспешный HTTP ответ: %s", resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка парсинга HTML: %w", err)
	}

	var aircrafts []model.Aircraft

	// Пример: парсим таблицу с данными воздушных судов
	doc.Find("table.aircraft-data tbody tr").Each(func(i int, s *goquery.Selection) {
		var aircraft model.Aircraft
		
		// Пример получения данных из ячеек таблицы (адаптировать под реальную структуру)
		aircraft.RegistrationNumber = strings.TrimSpace(s.Find("td:nth-child(1)").Text())
		aircraft.SerialNumber = strings.TrimSpace(s.Find("td:nth-child(2)").Text())
		
		// Добавляем только если есть хотя бы регистрационный номер
		if aircraft.RegistrationNumber != "" {
			aircrafts = append(aircrafts, aircraft)
		}
	})

	return aircrafts, nil
}

// Вспомогательная функция для парсинга CSV данных
func (p *AircraftDataParser) parseCSV(r io.Reader) ([]model.Aircraft, error) {
	reader := csv.NewReader(r)
	
	// Читаем заголовок для определения индексов колонок
	header, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения заголовка CSV: %w", err)
	}

	// Определяем индексы колонок
	regNumIdx := -1
	typeIdx := -1
	serialNumIdx := -1
	statusIdx := -1

	for i, col := range header {
		switch strings.ToLower(strings.TrimSpace(col)) {
		case "registration", "reg", "регистрационный номер":
			regNumIdx = i
		case "type", "aircraft type", "тип вс":
			typeIdx = i
		case "serial number", "серийный номер":
			serialNumIdx = i
		case "status", "статус":
			statusIdx = i
		}
	}

	if regNumIdx == -1 {
		return nil, errors.New("колонка с регистрационным номером не найдена")
	}

	var aircrafts []model.Aircraft

	// Читаем записи
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("ошибка чтения CSV: %w", err)
		}

		var aircraft model.Aircraft
		aircraft.RegistrationNumber = strings.TrimSpace(record[regNumIdx])
		
		if typeIdx >= 0 && typeIdx < len(record) {
			aircraft.Type = strings.TrimSpace(record[typeIdx])
		}
		
		if serialNumIdx >= 0 && serialNumIdx < len(record) {
			aircraft.SerialNumber = strings.TrimSpace(record[serialNumIdx])
		}
		
		if statusIdx >= 0 && statusIdx < len(record) {
			aircraft.Status = strings.TrimSpace(record[statusIdx])
		}

		// Добавляем только если есть регистрационный номер
		if aircraft.RegistrationNumber != "" {
			aircrafts = append(aircrafts, aircraft)
		}
	}

	return aircrafts, nil
}

// XML структуры для парсинга
type xmlAircraftList struct {
	XMLName   xml.Name      `xml:"AircraftList"`
	Aircrafts []xmlAircraft `xml:"Aircraft"`
}

type xmlAircraft struct {
	RegistrationNumber string `xml:"RegistrationNumber"`
	Type               string `xml:"Type"`
	SerialNumber       string `xml:"SerialNumber"`
	Status             string `xml:"Status"`
}

// Парсинг XML данных
func (p *AircraftDataParser) parseXML(r io.Reader) ([]model.Aircraft, error) {
	var xmlData xmlAircraftList
	decoder := xml.NewDecoder(r)
	
	if err := decoder.Decode(&xmlData); err != nil {
		return nil, fmt.Errorf("ошибка декодирования XML: %w", err)
	}

	var aircrafts []model.Aircraft
	for _, xmlAc := range xmlData.Aircrafts {
		aircraft := model.Aircraft{
			RegistrationNumber: xmlAc.RegistrationNumber,
			Type:               xmlAc.Type,
			SerialNumber:       xmlAc.SerialNumber,
			Status:             xmlAc.Status,
		}
		
		if aircraft.RegistrationNumber != "" {
			aircrafts = append(aircrafts, aircraft)
		}
	}

	return aircrafts, nil
} 