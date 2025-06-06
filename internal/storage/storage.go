package storage

type Storage interface {
	SetData(URL, newURL string)
	GetData(URL string) string
}

// example ->  abcd: https://google.com"
type URLStorage struct {
	data map[string]string
}

func NewStorage() *URLStorage {
	return &URLStorage{
		data: make(map[string]string),
	}
}

func (s *URLStorage) SetData(URL, newURL string) {
	s.data[newURL] = URL
}

func (s *URLStorage) GetData(newURL string) string {
	return s.data[newURL]
}
