package main

// fmt is used to print output on the terminal
import "fmt"

/*
Rule of thumb
✅ Reading data? → Use a value receiver.
✅ Changing data? → Use a pointer receiver.
✅ Struct is large? → Prefer a pointer receiver to avoid copying the whole struct every time.
*/

// Struct is like a blueprint or template.
// Here we are saying that every User will have these 4 fields.
type User struct {
	ID    int
	Name  string
	Age   int
	Email string
}

// ---------------- VALUE RECEIVER ----------------

// This method belongs to the User struct.
//
// (u User) means Go makes a COPY of the original user.
// Any changes inside this function will NOT affect the original struct.
//
// Here we simply return the user's name.
func (u User) GetName() string {
	return u.Name
}

// ---------------- POINTER RECEIVER ----------------

// (u *User) means this function gets the ACTUAL user,
// not a copy.
//
// So if we change anything here,
// the original struct will also change.
func (u *User) SetName(name string) {
	u.Name = name
}

func main() {

	fmt.Println("STRUCTS")

	// Creating a new User object.
	// Think of this like filling out a registration form.
	u1 := User{
		ID:    1,
		Name:  "JOHN",
		Age:   30,
		Email: "Email@gmail.com",
	}

	// Printing every field of the user.
	fmt.Println(u1.ID, u1.Name, u1.Age, u1.Email)

	// Updating only the Age field.
	// Since u1 is our actual struct,
	// this permanently changes Age from 30 to 31.
	u1.Age = 31

	fmt.Println(u1.Age, u1.Name)

	// Calling the GetName() method.
	// It returns the Name field.
	fmt.Println(u1.GetName())

	// Calling SetName().
	// Even though SetName() uses a pointer receiver,
	// Go automatically takes the address of u1 (&u1).
	// So we don't need to write (&u1).SetName().
	u1.SetName("David")

	// Name has now changed permanently.
	fmt.Println(u1.Name)

	// Calling GetName() again.
	fmt.Println(u1.GetName())
}
