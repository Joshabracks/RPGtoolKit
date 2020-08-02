package lycan


// UserCheck - precheck variables for a new User object
var UserCheck = []Attr {
	Attr {
		name: "name",
		typ: "string",
		required: true,
		min: 6,
		minErr: "User name must be at least 6 characters long",
		max: 20,
		maxErr: "User name cannot exceed 20 characters",
		requiredErr: "User name is Required",
		unique: true,
		uniqueErr: "That user name is already in use",
		crypt: false,
	},
	Attr {
		name: "email",
		typ: "string",
		required: true,
		min: 0,
		minErr: "nil",
		max: 100,
		maxErr: "Email cannot exceed 100 characters",
		requiredErr: "Email is required",
		unique: true,
		uniqueErr: "That email is already in use",
		crypt: false,
	},
	Attr {
		name: "password",
		typ: "string",
		required: true,
		min: 6,
		minErr: "Password name must be at least 6 characters long",
		max: 64,
		maxErr: "Password cannot exceed 64 characters",
		requiredErr: "Password is required",
		unique: false,
		crypt: false,
	},
}

// User constructed from POST data
// cross-checked by the user object
type User struct {
	name string
	email string
	password string
}

// NewUser - cross-checks variables for and constructs a new User object.


