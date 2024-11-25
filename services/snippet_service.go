package services

import (
	"encoding/json"
	"errors"
	"os"
)

type Snippet struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type SnippetService struct {
	filePath string
	snippets []Snippet
}

func NewSnippetService(filePath string) *SnippetService {
	return &SnippetService{
		filePath: filePath,
		snippets: []Snippet{},
	}
}

func (s *SnippetService) Load() error {
	file, err := os.Open(s.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			s.snippets = []Snippet{}
			return nil
		}
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&s.snippets); err != nil {
		return err
	}
	return nil
}

func (s *SnippetService) Save() error {
	file, err := os.Create(s.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(s.snippets)
}

func (s *SnippetService) GetAll() []Snippet {
	return s.snippets
}

func (s *SnippetService) Add(title, content string) error {
	if title == "" || content == "" {
		return errors.New("başlık ve içerik boş olamaz")
	}
	s.snippets = append(s.snippets, Snippet{Title: title, Content: content})
	return s.Save()
}

func (s *SnippetService) Update(index int, title, content string) error {
	if index < 0 || index >= len(s.snippets) {
		return errors.New("geçersiz indeks")
	}
	if title == "" || content == "" {
		return errors.New("başlık ve içerik boş olamaz")
	}
	s.snippets[index] = Snippet{Title: title, Content: content}
	return s.Save()
}

func (s *SnippetService) Delete(index int) error {
	if index < 0 || index >= len(s.snippets) {
		return errors.New("geçersiz indeks")
	}
	s.snippets = append(s.snippets[:index], s.snippets[index+1:]...)
	return s.Save()
}
