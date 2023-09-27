package creational

type singleton struct {
	count int
}

// define a pointer to a struct of type singleton as nil
var instance *singleton 

func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton) // new createsa pointer to an instance
	}
	return instance
}

func (s *singleton) AddOnce() int {
	s.count++
	return s.count
}