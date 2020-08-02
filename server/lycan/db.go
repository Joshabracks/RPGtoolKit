package lycan


func main() {

}

// GetAll : Takes in an Args struct and returns db struct or an error
func GetAll() {

}

// GetOne : Takes in an Args struct and returns db struct or an error
func GetOne() {

}

// Update : Takes in a db struct and updates any fields that do not match the database
// returns true if update is successful.  Otherwise returns error.
func Update() {

}

// Delete : Takes in db struct
// if an exact match is found in the DB, it is deleted and true is returned
// otherwise returns an error
func Delete() {

}

// New takes in a DB struct and runs validations against it
// if all validations pass, it is saved to the database and true is returned
// otherwise returns an error
func New() {

}

// DB "Database Object"
type DB struct {
	group string
	id int
	str []str
	num []num
	manyToMany []manyToMany
	oneToMany []oneToMany
	manyToOne []manyToOne
	oneToOne []oneToOne
}

type str struct {
	name string
	id int
	value string
	required bool
	requiredErr string
	unique bool
	uniqueErr string
	max int
	maxErr string
	min int
	minErr string
	crypt bool
}

type num struct {
	name string
	id int
	value int
	required bool
	requiredErr string
	unique bool
	uniqueErr string
	max int
	maxErr string
	min int
	minErr string
}

type manyToMany struct {
	name string
	id int
	group1 string
	group2 string
	value []string
}

type oneToMany struct {
	name string
	id int
	group1 string
	group2 string
	value []string
}

type manyToOne struct {
	name string
	id int
	group1 string
	group2 string
	value string     
}

type oneToOne struct {
	name string
	id int
	group1 string
	group2 string
	value string
}