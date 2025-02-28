package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/uuid"
)

// Структура для операции
type Operation struct {
	OperationType string  `json:"operationType"`
	Amount        float64 `json:"amount"`
}
type Transaction struct {
	WalletId      uuid.UUID `json:"walletId"`
	OperationType string    `json:"operationType"`
	Amount        float64   `json:"amount"`
}

func TestWalletOperations(t *testing.T) {
	// Устанавливаем базовый URL сервера
	url := "http://localhost:8080/api/v1/wallet"

	// Генерируем один UUID для всех операций
	walletID := uuid.New()
	Transaction := Transaction{WalletId: walletID}
	// Создаем список операций
	operations := []Operation{
		{"DEPOSIT", 1000},
		{"WITHDRAW", 500},
		{"DEPOSIT", 2000},
		{"WITHDRAW", 1000},
		{"DEPOSIT", 500},
		{"WITHDRAW", 200},
		{"DEPOSIT", 1000},
		{"WITHDRAW", 100},
		{"DEPOSIT", 300},
		{"WITHDRAW", 400},
		{"DEPOSIT", 100000}, // несуразная операция
		{"WITHDRAW", 50000}, // несуразная операция
	}

	// Выполнение запросов
	for _, op := range operations {
		// Формируем тело запроса
		Transaction.OperationType = op.OperationType
		Transaction.Amount = op.Amount
		// Преобразуем тело в JSON
		body, err := json.Marshal(Transaction)
		if err != nil {
			t.Fatalf("Error marshalling request body: %v", err)
		}

		// Создаем POST-запрос
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
		if err != nil {
			t.Fatalf("Error creating request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		// Отправляем запрос
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("Error sending request: %v", err)
		}
		defer resp.Body.Close()

		// Проверяем статус ответа
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code 200, got %v", resp.StatusCode)
		}

		// Выводим ответ для проверки
		fmt.Printf("Response: %v\n", resp.Status)
	}

	// Можно добавить ожидания между запросами, чтобы убедиться, что сервер обрабатывает их параллельно.
}
