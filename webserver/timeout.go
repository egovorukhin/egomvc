package webserver

type Timeout struct {
	Read  int `yaml:"read"`
	Write int `yaml:"write"`
	Idle  int `yaml:"idle"`
}

func (t Timeout) Get() (int, int, int) {

	/*
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

		//простой
	*/
	return t.Read, t.Write, t.Idle
}
