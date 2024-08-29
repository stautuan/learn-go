package main

// interface
type employee interface {
    getName()   string
    getSalary() int
}

// struct
type contractor struct {
    name         string
    hourlyPay    int
    hoursPerYear int
}
func (c contractor) getSalary() int {
    salary := c.hourlyPay * c.hoursPerYear
    return salary
}
func (c contractor) getName() string {
    return c.name
}


// struct
type fullTime struct {
    name   string
    salary int
}
func (ft fullTime) getSalary() int {
    return ft.salary
}
func (ft fullTime) getName() string {
    return ft.name
}