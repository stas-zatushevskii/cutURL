package storage

type Storage interface {
	SetData(URL string)
	GetData(int int) string
}

// поле id для подсчета используемых айдишников в бд
// example: 1: "{"oldURL: NewURL"}"
type URLStorage struct {
	data map[int]string
	id   int
}

func NewStorage() *URLStorage {
	return &URLStorage{
		data: make(map[int]string),
		id:   0,
	}
}

func (s *URLStorage) SetData(URL string) {
	s.data[s.id] = URL
	s.id++
}

func (s *URLStorage) GetData(id int) string {
	return s.data[id]
}
