package data

import (
	"database/sql"
	"fmt"
	"time"
)

func New_Test(dbPool *sql.DB) Models {
	db = dbPool
	return Models{
		User: &UserTest{},
		Plan: &PlanTest{},
	}
}

// User is the structure which holds one user from the database.
type UserTest struct{}

// GetAll returns a slice of all users, sorted by last name
func (u *UserTest) GetAll() ([]*User, error) {

	users := []*User{
		{
			ID:        2,
			Email:     "admin@example.com",
			FirstName: "Admin",
			LastName:  "Admin",
			Password:  "abc",
			Active:    1,
			IsAdmin:   1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Plan:      nil,
		},
	}

	return users, nil
}

// GetByEmail returns one user by email
func (u *UserTest) GetByEmail(email string) (*User, error) {
	user := User{
		ID:        2,
		Email:     "admin@example.com",
		FirstName: "Admin",
		LastName:  "Admin",
		Password:  "abc",
		Active:    1,
		IsAdmin:   1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Plan:      nil,
	}

	return &user, nil
}

// GetOne returns one user by id
func (u *UserTest) GetOne(id int) (*User, error) {
	return u.GetByEmail("")
}

// Update updates one user in the database, using the information
// stored in the receiver u
func (u *UserTest) Update(user *User) error {
	return nil
}

// Delete deletes one user from the database, by User.ID
func (u *UserTest) Delete() error {
	return nil
}

// DeleteByID deletes one user from the database, by ID
func (u *UserTest) DeleteByID(id int) error {
	return nil
}

// Insert inserts a new user into the database, and returns the ID of the newly inserted row
func (u *UserTest) Insert(user User) (int, error) {
	return 2, nil
}

// ResetPassword is the method we will use to change a user's password.
func (u *UserTest) ResetPassword(password string) error {
	return nil
}

// PasswordMatches uses Go's bcrypt package to compare a user supplied password
// with the hash we have stored for a given user in the database. If the password
// and hash match, we return true; otherwise, we return false.
func (u *UserTest) PasswordMatches(plainText string) (bool, error) {
	return true, nil
}

// Plan is the type for subscription plans
type PlanTest struct{}

func (p *PlanTest) GetAll() ([]*Plan, error) {
	plans := []*Plan{
		{
			ID:                  4,
			PlanName:            "Bronze",
			PlanAmount:          1000,
			PlanAmountFormatted: "1000.00",
			CreatedAt:           time.Now(),
			UpdatedAt:           time.Now(),
		},
	}
	return plans, nil
}

// GetOne returns one plan by id
func (p *PlanTest) GetOne(id int) (*Plan, error) {
	plan := Plan{
		ID:                  4,
		PlanName:            "Bronze",
		PlanAmount:          1000,
		PlanAmountFormatted: "1000.00",
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}

	return &plan, nil
}

// SubscribeUserToPlan subscribes a user to one plan by insert
// values into user_plans table
func (p *PlanTest) SubscribeUserToPlan(user User, plan Plan) error {
	return nil
}

// AmountForDisplay formats the price we have in the DB as a currency string
func (p *PlanTest) AmountForDisplay() string {
	return fmt.Sprintf("$%.2f", 1000.0)
}
