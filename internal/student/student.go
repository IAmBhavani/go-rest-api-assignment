package student

import (
	"context"
	"errors"

	log "github.com/sirupsen/logrus"
)

var (
	ErrFetchingStudent  = errors.New("could not fetch student by ID")
	ErrFetchingStudents = errors.New("could not fetch students")
	ErrUpdatingStudent  = errors.New("could not update student")
	ErrNoStudentFound   = errors.New("no student found")
	ErrDeletingStudent  = errors.New("could not delete student")
	ErrNotImplemented   = errors.New("not implemented")
)

// Student - defines our Student structure
type Student struct {
	ID      string `json:"id"`
	Fname   string `json:"fname"`
	Lname   string `json:"lname"`
	DOB     string `json:"date_of_birth"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Gender  string `json:"gender"`

	CreatedBy string `json:"createdBy"`
	CreatedOn string `json:"createdOn"`
	UpdatedBy string `json:"updatedBy"`
	UpdatedOn string `json:"updatedOn"`
}

// StudentStore - defines the interface we need our Student storage
// layer to implement
type StudentStore interface {
	GetStudents(context.Context) ([]Student, error)
	GetStudent(context.Context, string) (Student, error)
	PostStudent(context.Context, Student) (Student, error)
	UpdateStudent(context.Context, string, Student) (Student, error)
	DeleteStudent(context.Context, string) error
	Ping(context.Context) error
}

// Service - the struct for our Student service
type Service struct {
	Store StudentStore
}

// NewService - returns a new Student service
func NewService(store StudentStore) *Service {
	return &Service{
		Store: store,
	}
}

// GetStudents - retrieves all Students from the database
func (s *Service) GetStudents(ctx context.Context) ([]Student, error) {
	// Call the store to get all students
	students, err := s.Store.GetStudents(ctx)
	if err != nil {
		log.Errorf("an error occurred fetching the Students: %s", err.Error())
		return nil, ErrFetchingStudents // Assuming ErrFetchingStudents is defined similarly to ErrFetchingStudent
	}
	return students, nil
}

// GetStudent - retrieves Students by their ID from the database
func (s *Service) GetStudent(ctx context.Context, ID string) (Student, error) {
	// calls store passing in the context
	std, err := s.Store.GetStudent(ctx, ID)
	if err != nil {
		log.Errorf("an error occured fetching the Student: %s", err.Error())
		return Student{}, ErrFetchingStudent
	}
	return std, nil
}

// PostStudent - adds a new Student to the database
func (s *Service) PostStudent(ctx context.Context, std Student) (Student, error) {
	std, err := s.Store.PostStudent(ctx, std)
	if err != nil {
		log.Errorf("an error occurred adding the Student: %s", err.Error())
	}
	return std, nil
}

// UpdateStudent - updates a Student by ID with new Student info
func (s *Service) UpdateStudent(
	ctx context.Context, ID string, newStudent Student,
) (Student, error) {
	std, err := s.Store.UpdateStudent(ctx, ID, newStudent)
	if err != nil {
		log.Errorf("an error occurred updating the Student: %s", err.Error())
	}
	return std, nil
}

// DeleteStudent - deletes a Student from the database by ID
func (s *Service) DeleteStudent(ctx context.Context, ID string) error {
	return s.Store.DeleteStudent(ctx, ID)
}

// ReadyCheck - a function that tests we are functionally ready to serve requests
func (s *Service) ReadyCheck(ctx context.Context) error {
	log.Info("Checking readiness")
	return s.Store.Ping(ctx)
}
