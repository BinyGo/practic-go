package main

type NotFoundError struct {
	Name string
}

func (e *NotFoundError) Error() string {
	return e.Name + ": not found"
}

// if e,ok=err.(*NotFoundError);ok{
// 	// e.Name wasn't found
// }

func main() {

}
