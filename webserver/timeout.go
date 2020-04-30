package webserver

type Timeout struct {
	Read  int `json:"read" xml:"read"`
	Write int `json:"write" xml:"write"`
}

func (t Timeout) Get() (int, int) {

	//чтение
	read := 30
	if t.Read > 0 {
		read = t.Read
	}

	//запись
	write := 30
	if t.Write > 0 {
		read = t.Write
	}

	return read, write
}
