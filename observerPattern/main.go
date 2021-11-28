package main

import "fmt"

func main() {
	subject := Subject{}
	oa := Observable{Name: "Anna"}
	ob := Observable{Name: "Brad"}
	subject.registerObserver(&Observer{})
	subject.notifyObservers(oa, ob)
	fmt.Println(" -> observers now.")
	oc := Observable{Name: "Claudia"}
	subject.notifyObservers(oa, ob, oc)
	fmt.Println(" -> observers now.")
}

/* observer - Notify method(call back function)*/
type Observable struct {
	Name string
}

type Observer struct{}

func (ob *Observer) Notify(o *Observable) {
	print(o.Name + " ")
}

type Callback interface {
	Notify(o *Observable)
}

//subject - three method - registerObserver, removeObserver, NotifyObserver
type Subject struct {
	callbacks []Callback
}

//append from callback list when observer registered
func (o *Subject) registerObserver(c Callback) {
	o.callbacks = append(o.callbacks, c)
}

//append from callback list when observer removed
func (o *Subject) removeObserver(c Callback) {
	o.callbacks = append(o.callbacks, c)

	newCallbacks := []Callback{}
	for _, cb := range o.callbacks {
		if cb != c {
			newCallbacks = append(newCallbacks, cb)
		}
	}
	o.callbacks = newCallbacks
}

func (o *Subject) notifyObservers(oes ...Observable) {
	for _, oe := range oes {
		for _, c := range o.callbacks {
			c.Notify(&oe)
		}
	}
}
