package data

import (
	"database/sql"
	"fmt"
	"time"
)

func TestNew(dbPool *sql.DB) Models {
	db = dbPool

	return Models{
		User: &UserTest{},
		Plan: &PlanTest{},
	}
}

type UserTest struct {
	ID        int
	Email     string
	FirstName string
	LastName  string
	Password  string
	Active    int
	IsAdmin   int
	CreatedAt time.Time
	UpdatedAt time.Time
	Plan      *Plan
}

// GetAll returns a slice of all users, sorted by last name
func (u *UserTest) GetAll() ([]*User, error) {
	var users []*User
	// take password from DB, otherwise "TestConfig_PostLoginPage" test will fail
	// due to "err: crypto/bcrypt: hashedSecret too short to be a bcrypted password"
	user := User{
		ID:        1,
		Email:     "admin@example.com",
		FirstName: "Admin",
		LastName:  "Example",
		Password:  "abc",
		// Password:  "$2a$12$1zGLuYDDNvATh4RA4avbKuheAMpb1svexSzrQm7up.bnpwQHs0jNe",
		Active:    1,
		IsAdmin:   1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	users = append(users, &user)

	return users, nil
}

// GetByEmail returns one user by email
func (u *UserTest) GetByEmail(email string) (*User, error) {
	user := User{
		ID:        1,
		Email:     "admin@example.com",
		FirstName: "Admin",
		LastName:  "Example",
		Password:  "abc",
		// Password:  "$2a$12$1zGLuYDDNvATh4RA4avbKuheAMpb1svexSzrQm7up.bnpwQHs0jNe",
		Active:    1,
		IsAdmin:   1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &user, nil
}

// GetOne returns one user by id
func (u *UserTest) GetOne(id int) (*User, error) {
	return u.GetByEmail("")
}

func (u *UserTest) Update(user User) error {
	return nil
}

func (u *UserTest) Delete() error {
	return nil
}

func (u *UserTest) DeleteByID(id int) error {
	return nil
}

func (u *UserTest) Insert(user User) (int, error) {
	return 2, nil
}

func (u *UserTest) ResetPassword(password string) error {
	return nil
}

func (u *UserTest) PasswordMatches(hashedPassword, plainText string) (bool, error) {
	fmt.Println("Called func of test-models/userTest")
	return true, nil
}

type PlanTest struct {
	ID                  int
	PlanName            string
	PlanAmount          int
	PlanAmountFormatted string
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

func (p *PlanTest) GetAll() ([]*Plan, error) {
	var plans []*Plan

	plan := Plan{
		ID:         1,
		PlanName:   "Broze Plan",
		PlanAmount: 1000,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	plans = append(plans, &plan)
	return plans, nil
}

// GetOne returns one plan by id
func (p *PlanTest) GetOne(id int) (*Plan, error) {
	plan := Plan{
		ID:         1,
		PlanName:   "Broze Plan",
		PlanAmount: 1000,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
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
	return "$10.00"
}
