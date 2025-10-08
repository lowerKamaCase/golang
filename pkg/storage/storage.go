package storage

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
)

type EmailHash struct {
	Email string `json:"email"`
	Hash  string `json:"hash"`
}

type EmailStorage struct {
	filename string
	mu       *sync.RWMutex
}

func NewEmailStorage(filename string, mu *sync.RWMutex) *EmailStorage {
	return &EmailStorage{
		filename: filename,
		mu:       mu,
	}
}

// Add добавляет новую запись
func (s *EmailStorage) Add(email, hash string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	records, err := s.readAllRecords()
	if err != nil {
		return err
	}
	fmt.Println(records)

	var updatedRecords []EmailHash
	updatedRecords = append(updatedRecords, records...)
	updatedRecords = append(updatedRecords, EmailHash{Email: email, Hash: hash})

	s.writeAllRecords(updatedRecords)
	return err
}

// DeleteByEmail удаляет запись по email
func (s *EmailStorage) DeleteByEmail(email string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	records, err := s.readAllRecords()
	if err != nil {
		return err
	}

	var found bool
	var updatedRecords []EmailHash

	for _, record := range records {
		if record.Email == email {
			found = true
			fmt.Printf("🗑️  Deleting email: %s\n", email)
			continue
		}
		updatedRecords = append(updatedRecords, record)
	}

	if !found {
		return fmt.Errorf("email not found: %s", email)
	}

	return s.writeAllRecords(updatedRecords)
}

func (s *EmailStorage) DeleteByHash(hash string) error {
	fmt.Println("Got hash: ", hash)
	s.mu.Lock()
	defer s.mu.Unlock()

	records, err := s.readAllRecords()
	if err != nil {
		return err
	}

	var found bool
	var updatedRecords []EmailHash

	for _, record := range records {
		if record.Hash == hash {
			found = true
			fmt.Printf("🗑️  Deleting email by hash: %s\n", hash)
			continue
		}
		updatedRecords = append(updatedRecords, record)
	}

	if !found {
		return fmt.Errorf("hash not found: %s", hash)
	}

	for _, rec := range updatedRecords {
		fmt.Println("rec: ", rec)
	}

	return s.writeAllRecords(updatedRecords)
}

// Update обновляет хеш для email
func (s *EmailStorage) Update(email, newHash string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	records, err := s.readAllRecords()
	if err != nil {
		return err
	}

	var found bool
	for i, record := range records {
		if record.Email == email {
			records[i].Hash = newHash
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("email not found: %s", email)
	}

	return s.writeAllRecords(records)
}

// GetHashByEmail возвращает хеш по email
func (s *EmailStorage) GetHashByEmail(email string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	records, err := s.readAllRecords()
	if err != nil {
		return "", err
	}

	for _, record := range records {
		if record.Email == email {
			return record.Hash, nil
		}
	}

	return "", fmt.Errorf("email not found: %s", email)
}

// GetAll возвращает все записи в виде map
func (s *EmailStorage) GetAll() (map[string]string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	records, err := s.readAllRecords()
	if err != nil {
		return nil, err
	}

	result := make(map[string]string)
	for _, record := range records {
		result[record.Email] = record.Hash
	}

	return result, nil
}

// Вспомогательные методы
func (s *EmailStorage) readAllRecords() ([]EmailHash, error) {
	file, err := os.Open(s.filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []EmailHash{}, nil
		}
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return []EmailHash{}, nil
	}

	var recordsMap map[string]string
	if err := json.Unmarshal(data, &recordsMap); err != nil {
		return nil, err
	}

	var records []EmailHash
	for email, hash := range recordsMap {
		records = append(records, EmailHash{
			Email: email,
			Hash:  hash,
		})
	}

	return records, nil
}

func (s *EmailStorage) writeAllRecords(records []EmailHash) error {
	file, err := os.Create(s.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	dataMap := make(map[string]string)
	for _, record := range records {
		dataMap[record.Email] = record.Hash
	}

	jsonData, err := json.MarshalIndent(dataMap, "", "  ")
	if err != nil {
		return err
	}

	if _, err := file.Write(jsonData); err != nil {
		return err
	}

	return nil
}


func (s *EmailStorage) Exists(email string) (bool, error) {
	users, err := s.GetAll()
	if err != nil {
		return false, err
	}
	_, exists := users[email]
	return exists, nil
}
