package main

func getExpenseReport(e expense) (string, float64) {
	// e is an istance of expense
	// check if expense is an email and sms
	em, ok := e.(email)
	if ok {
		return em.toAddress, em.cost()
	}

	sm, ok := e.(email)
	if ok {
		return sm.toPhoneNumber, sm.cost()
	}

	return "", 0.0
}

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

// this is an empty struct, it occupies zero bytes
type invalid struct{}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}