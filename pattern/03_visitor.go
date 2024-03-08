package pattern

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

// интерфейс Visitor имеет методы, принимающие уакзатели на структуры,
// с которыми мы хотим взаимодействовать.
// таким образом мы получаем доступ к полям структур,
// не добавляя никаких новых функций в сами структуры
type Visitor interface {
	VisitManager(m *Manager)
	VisitProgrammer(p *Programmer)
	VisitSalesman(s *Salesman)
}

// Интерфейс Employee имеет метод Accept, который принимает/внедряет визитор
type Employee interface {
	Accept(v Visitor)
}

type Manager struct {
	Name   string
	Salary float64
}

func (m *Manager) Accept(v Visitor) {
	v.VisitManager(m)
}

type Programmer struct {
	Name   string
	Salary float64
}

func (p *Programmer) Accept(v Visitor) {
	v.VisitProgrammer(p)
}

type Salesman struct {
	Name   string
	Salary float64
}

func (s *Salesman) Accept(v Visitor) {
	v.VisitSalesman(s)
}

// визитор, который будет считать TotalSalary
type SalaryCalculator struct {
	TotalSalary float64
}

func (sc *SalaryCalculator) VisitManager(m *Manager) {
	sc.TotalSalary += m.Salary
}

func (sc *SalaryCalculator) VisitProgrammer(p *Programmer) {
	sc.TotalSalary += p.Salary
}

func (sc *SalaryCalculator) VisitSalesman(s *Salesman) {
	sc.TotalSalary += s.Salary
}
