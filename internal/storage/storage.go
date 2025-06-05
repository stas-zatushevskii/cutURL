package storage

type Storage interface {
	SetData(url string)
	GetData(int int) string
}

// поле id для подсчета используемых айдишников в бд
// example: 1: "{"oldURL: NewURL"}"
type UrlStorage struct {
	data map[int]string
	id   int
}

func NewStorage() *UrlStorage {
	return &UrlStorage{
		data: make(map[int]string),
		id:   0,
	}
}

func (s *UrlStorage) SetData(url string) {
	s.data[s.id] = url
	s.id++
}

func (s *UrlStorage) GetData(id int) string {
	return s.data[id]
}
