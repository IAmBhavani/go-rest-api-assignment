package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"go-rest-api-assignment/internal/student"
)

var (
	ErrNotImplemented = errors.New("not implemented")
)

// StudentRow - models how our Students look in the database
type StudentRow struct {
	ID        string
	Fname     sql.NullString
	Lname     sql.NullString
	DOB       sql.NullTime   `db:"date_of_birth"`
	Email     sql.NullString `json:"email"`
	Address   sql.NullString `json:"address"`
	Gender    sql.NullString `json:"gender"`
	CreatedBy sql.NullString `db:"createdBy"`
	CreatedOn sql.NullTime   `db:"createdOn"`
	UpdatedBy sql.NullString `db:"updatedBy"`
	UpdatedOn sql.NullTime   `db:"updatedOn"`
}

func convertStudentRowToStudent(c StudentRow) student.Student {
	return student.Student{
		ID:        c.ID,
		Fname:     c.Fname.String,
		Lname:     c.Lname.String,
		DOB:       convertNullTime(c.DOB),
		Email:     c.Email.String,
		Address:   c.Address.String,
		Gender:    c.Gender.String,
		CreatedBy: c.CreatedBy.String,
		CreatedOn: convertNullTime(c.CreatedOn),
		UpdatedBy: c.UpdatedBy.String,
		UpdatedOn: convertNullTime(c.UpdatedOn),
	}
}

func convertNullTime(nt sql.NullTime) string {
	if nt.Valid {
		return nt.Time.Format(time.RFC3339)
	}
	return "" // Returns the zero value of time.Time
}

// GetStudents - retrieves all students from the database
func (d *Database) GetStudents(ctx context.Context) ([]student.Student, error) {
	// Prepare to fetch all student records
	rows, err := d.Client.QueryContext(
		ctx,
		`SELECT id, fname, lname, date_of_birth, email, address, gender, createdBy, createdOn, updatedBy, updatedOn 
		FROM students`,
	)
	if err != nil {
		return nil, fmt.Errorf("an error occurred fetching students: %w", err)
	}
	defer rows.Close()

	// Slice to hold the retrieved students
	var students []student.Student

	// Iterate through the rows and scan the data
	for rows.Next() {
		var stdRow StudentRow
		err := rows.Scan(&stdRow.ID, &stdRow.Fname, &stdRow.Lname, &stdRow.DOB, &stdRow.Email, &stdRow.Address, &stdRow.Gender, &stdRow.CreatedBy, &stdRow.CreatedOn, &stdRow.UpdatedBy, &stdRow.UpdatedOn)
		if err != nil {
			return nil, fmt.Errorf("an error occurred scanning a student row: %w", err)
		}

		// Convert the row data into a student.Student and append it to the slice
		students = append(students, convertStudentRowToStudent(stdRow))
	}

	// Check for any error that may have occurred during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("an error occurred after scanning student rows: %w", err)
	}

	return students, nil
}

// GetStudent - retrieves a Student from the database by ID
func (d *Database) GetStudent(ctx context.Context, uuid string) (student.Student, error) {
	// fetch StudentRow from the database and then convert to student.Student
	var stdRow StudentRow
	row := d.Client.QueryRowContext(
		ctx,
		`SELECT id, fname, lname, date_of_birth, email, address, gender, createdBy, createdOn, updatedBy, updatedOn 
		FROM students 
		WHERE id = ?`,
		uuid,
	)
	err := row.Scan(&stdRow.ID, &stdRow.Fname, &stdRow.Lname, &stdRow.DOB, &stdRow.Email, &stdRow.Address, &stdRow.Gender, &stdRow.CreatedBy, &stdRow.CreatedOn, &stdRow.UpdatedBy, &stdRow.UpdatedOn)
	if err != nil {
		return student.Student{}, fmt.Errorf("an error occurred fetching a Student by uuid: %w", err)
	}
	// sqlx with context to ensure context cancelation is honoured
	return convertStudentRowToStudent(stdRow), nil
}

// PostStudent - adds a new Student to the database
func (d *Database) PostStudent(ctx context.Context, std student.Student) (student.Student, error) {
	fmt.Printf("Received student: %+v\n", std)
	//std.ID = uuid.NewV4().String()

	dobTime, err := time.Parse(time.RFC3339, std.DOB)
	if err != nil {
		return student.Student{}, fmt.Errorf("invalid date format for DOB: %w", err)
	}

	postRow := StudentRow{
		//ID:      std.ID,
		Fname:   sql.NullString{String: std.Fname, Valid: true},
		Lname:   sql.NullString{String: std.Lname, Valid: true},
		DOB:     sql.NullTime{Time: dobTime, Valid: true},
		Email:   sql.NullString{String: std.Email, Valid: true},
		Address: sql.NullString{String: std.Address, Valid: true},
		Gender:  sql.NullString{String: std.Gender, Valid: true},

		CreatedBy: sql.NullString{String: ctx.Value("userID").(string), Valid: true},
		CreatedOn: sql.NullTime{Time: time.Now(), Valid: true},
	}

	fmt.Printf("Inserting row: %+v\n", postRow)

	rows, err := d.Client.NamedQueryContext(
		ctx,
		`INSERT INTO students 
		(fname, lname, date_of_birth, email, address, gender, createdBy, createdOn) VALUES
		(:fname, :lname, :date_of_birth, :email, :address, :gender, :createdBy, :createdOn)`,
		postRow,
	)
	if err != nil {
		return student.Student{}, fmt.Errorf("failed to insert Student: %w", err)
	}
	if err := rows.Close(); err != nil {
		return student.Student{}, fmt.Errorf("failed to close rows: %w", err)
	}

	return std, nil
}

// UpdateStudent - updates a Student in the database
func (d *Database) UpdateStudent(ctx context.Context, id string, std student.Student) (student.Student, error) {

	dobTime, err := time.Parse(time.RFC3339, std.DOB)
	if err != nil {
		return student.Student{}, fmt.Errorf("invalid date format for DOB: %w", err)
	}

	// UpdatedTime, err := time.Parse(time.RFC3339, std.UpdatedOn)
	// if err != nil {
	// 	return student.Student{}, fmt.Errorf("invalid date format for DOB: %w", err)
	// }
	stdRow := StudentRow{
		ID:      id,
		Fname:   sql.NullString{String: std.Fname, Valid: true},
		Lname:   sql.NullString{String: std.Lname, Valid: true},
		DOB:     sql.NullTime{Time: dobTime, Valid: true},
		Email:   sql.NullString{String: std.Email, Valid: true},
		Address: sql.NullString{String: std.Address, Valid: true},
		Gender:  sql.NullString{String: std.Gender, Valid: true},

		UpdatedBy: sql.NullString{String: ctx.Value("userID").(string), Valid: true},
		UpdatedOn: sql.NullTime{Time: time.Now(), Valid: true},
	}

	fmt.Printf("updating row: %+v\n", stdRow)

	rows, err := d.Client.NamedQueryContext(
		ctx,
		`UPDATE students SET
		fname = :fname,
		lname = :lname,
		date_of_birth = :date_of_birth,
		email = :email,
		address = :address,
		gender = :gender,
		updatedBy = :updatedBy,
		updatedOn = :updatedOn
		WHERE id = :id`,
		stdRow,
	)
	if err != nil {
		return student.Student{}, fmt.Errorf("failed to insert Student: %w", err)
	}
	if err := rows.Close(); err != nil {
		return student.Student{}, fmt.Errorf("failed to close rows: %w", err)
	}

	return convertStudentRowToStudent(stdRow), nil
}

// DeleteStudent - deletes a Student from the database
func (d *Database) DeleteStudent(ctx context.Context, id string) error {
	fmt.Printf("Deleting row with id: %+v\n", id)

	_, err := d.Client.ExecContext(
		ctx,
		`DELETE FROM students where id = ?`,
		id,
	)
	if err != nil {
		return fmt.Errorf("failed to delete Student from the database: %w", err)
	}
	return nil
}
