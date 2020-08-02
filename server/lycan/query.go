package lycan

// Attr objects are used to construct the models used to verify db objects
type Attr struct {
	name, minErr, maxErr, typ, requiredErr, uniqueErr string
	required, unique, crypt                           bool
	min, max                                          int
}

type query interface {
	getOne()
	getAll()
	update()
	delete()
	new()
	hash()
}

// Args : Database argument requires a group and either value or id
// value takes in a 2 string ["key" "value"] array
// id takes in a single string
type Args struct {
	group string
	value [2]string
	id    int
}

func (args Args) getOne() {

}

func (args Args) getAll() {

}

func (o DB) update() {

}

func (o DB) delete() {

}

func (o DB) new() {

}

func (s str) hash() {

}
