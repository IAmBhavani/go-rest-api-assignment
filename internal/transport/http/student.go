package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"go-rest-api-assignment/internal/student"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/go-playground/validator/v10"
)

type StudentService interface {
	GetStudents(ctx context.Context) ([]student.Student, error)
	GetStudent(ctx context.Context, ID string) (student.Student, error)
	PostStudent(ctx context.Context, cmt student.Student) (student.Student, error)
	UpdateStudent(ctx context.Context, ID string, newStd student.Student) (student.Student, error)
	DeleteStudent(ctx context.Context, ID string) error
	ReadyCheck(ctx context.Context) error
}

// GetStudents - retrieves all students
func (h *Handler) GetStudents(w http.ResponseWriter, r *http.Request) {
	students, err := h.Service.GetStudents(r.Context())
	if err != nil {
		log.Error("Failed to fetch students:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if len(students) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(students); err != nil {
		log.Error("Failed to encode response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// GetStudent - retrieve a student by ID
func (h *Handler) GetStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "Bad Request: missing student ID", http.StatusBadRequest)
		return
	}

	cmt, err := h.Service.GetStudent(r.Context(), id)
	if err != nil {
		if errors.Is(err, student.ErrFetchingStudent) {
			http.Error(w, "Student not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		log.Error("Failed to encode response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// PostStudentRequest
type PostStudentRequest struct {
	Fname     string `json:"fname" validate:"required"`
	Lname     string `json:"lname" validate:"required"`
	DOB       string `json:"date_of_birth" validate:"required"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Gender    string `json:"gender"`
	CreatedBy string `json:"createdBy"`
	CreatedOn string `json:"createdOn"`
	UpdatedBy string `json:"updatedBy"`
	UpdatedOn string `json:"updatedOn"`
}

func studentFromPostStudentRequest(u PostStudentRequest) student.Student {
	return student.Student{
		Fname:     u.Fname,
		Lname:     u.Lname,
		DOB:       u.DOB,
		Email:     u.Email,
		Address:   u.Address,
		Gender:    u.Gender,
		CreatedBy: u.CreatedBy,
		CreatedOn: u.CreatedOn,
		UpdatedBy: u.UpdatedBy,
		UpdatedOn: u.UpdatedOn,
	}
}

// PostStudent - adds a new student
func (h *Handler) PostStudent(w http.ResponseWriter, r *http.Request) {
	var postStdReq PostStudentRequest
	if err := json.NewDecoder(r.Body).Decode(&postStdReq); err != nil {
		http.Error(w, "Bad Request: failed to decode request body", http.StatusBadRequest)
		return
	}

	validate := validator.New()
	err := validate.Struct(postStdReq)
	if err != nil {
		log.Info("Validation error:", err)
		http.Error(w, "Bad Request: validation failed", http.StatusBadRequest)
		return
	}

	std := studentFromPostStudentRequest(postStdReq)
	std, err = h.Service.PostStudent(r.Context(), std)
	if err != nil {
		log.Error("Failed to create student:", err)
		http.Error(w, "Internal Server Error: failed to create student", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(std); err != nil {
		log.Error("Failed to encode response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// UpdateStudentRequest -
type UpdateStudentRequest struct {
	Fname     string `json:"fname" validate:"required"`
	Lname     string `json:"lname" validate:"required"`
	DOB       string `json:"date_of_birth" validate:"required"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Gender    string `json:"gender"`
	CreatedBy string `json:"createdBy"`
	CreatedOn string `json:"createdOn"`
	UpdatedBy string `json:"updatedBy"`
	UpdatedOn string `json:"updatedOn"`
}

// convert the validated struct into something that the service layer understands
func studentFromUpdateStudentRequest(u UpdateStudentRequest) student.Student {
	return student.Student{
		Fname:     u.Fname,
		Lname:     u.Lname,
		DOB:       u.DOB,
		Email:     u.Email,
		Address:   u.Address,
		Gender:    u.Gender,
		CreatedBy: u.CreatedBy,
		CreatedOn: u.CreatedOn,
		UpdatedBy: u.UpdatedBy,
		UpdatedOn: u.UpdatedOn,
	}
}

// UpdateStudent - updates a student by ID
func (h *Handler) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["id"]

	var updateStdRequest UpdateStudentRequest
	if err := json.NewDecoder(r.Body).Decode(&updateStdRequest); err != nil {
		http.Error(w, "Bad Request: failed to decode request body", http.StatusBadRequest)
		return
	}

	validate := validator.New()
	err := validate.Struct(updateStdRequest)
	if err != nil {
		log.Println("Validation error:", err)
		http.Error(w, "Bad Request: validation failed", http.StatusBadRequest)
		return
	}

	std := studentFromUpdateStudentRequest(updateStdRequest)
	std, err = h.Service.UpdateStudent(r.Context(), studentID, std)
	if err != nil {
		log.Error("Failed to update student:", err)
		http.Error(w, "Internal Server Error: failed to update student", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(std); err != nil {
		log.Error("Failed to encode response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// DeleteStudent - deletes a student by ID
func (h *Handler) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["id"]

	if studentID == "" {
		http.Error(w, "Bad Request: missing student ID", http.StatusBadRequest)
		return
	}

	err := h.Service.DeleteStudent(r.Context(), studentID)
	if err != nil {
		log.Error("Failed to delete student:", err)
		http.Error(w, "Internal Server Error: failed to delete student", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Response{Message: "Successfully Deleted"}); err != nil {
		log.Error("Failed to encode response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
