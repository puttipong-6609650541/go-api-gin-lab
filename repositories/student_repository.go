package repositories

import (
	"database/sql"
	"fmt"

	"example.com/student-api/models"
)

type StudentRepository struct {
	DB *sql.DB
}

func (r *StudentRepository) GetAll() ([]models.Student, error) {
	rows, err := r.DB.Query("SELECT id, name, major, gpa FROM students")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []models.Student
	for rows.Next() {
		var s models.Student
		rows.Scan(&s.Id, &s.Name, &s.Major, &s.GPA)
		students = append(students, s)
	}
	return students, nil
}

func (r *StudentRepository) GetByID(id string) (*models.Student, error) {
	row := r.DB.QueryRow(
		"SELECT id, name, major, gpa FROM students WHERE id = ?",
		id,
	)

	var s models.Student
	err := row.Scan(&s.Id, &s.Name, &s.Major, &s.GPA)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *StudentRepository) Create(s models.Student) error {
	_, err := r.DB.Exec(
		"INSERT INTO students (id, name, major, gpa) VALUES (?, ?, ?, ?)",
		s.Id, s.Name, s.Major, s.GPA,
	)
	return err
}

// Update แก้ไขข้อมูลนักศึกษาตาม ID
// คืนค่า error = nil ถ้าสำเร็จ, error ถ้าไม่พบหรือเกิดปัญหา
func (r *StudentRepository) Update(id string, s models.Student) (*models.Student, error) {
	// ทำการ UPDATE และดูว่า rows ที่ถูก affect มีกี่แถว
	result, err := r.DB.Exec(
		"UPDATE students SET name = ?, major = ?, gpa = ? WHERE id = ?",
		s.Name, s.Major, s.GPA, id,
	)
	if err != nil {
		return nil, err
	}

	// RowsAffected บอกว่า UPDATE กระทบกี่แถว
	// ถ้าได้ 0 แปลว่าไม่มี student ID นี้ในฐานข้อมูล
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, fmt.Errorf("student not found")
	}

	// ดึงข้อมูลล่าสุดกลับมาส่งให้ client
	s.Id = id
	return &s, nil
}

// Delete ลบนักศึกษาตาม ID
// ใช้หลักการเดียวกับ Update คือเช็ค RowsAffected
func (r *StudentRepository) Delete(id string) error {
	result, err := r.DB.Exec(
		"DELETE FROM students WHERE id = ?",
		id,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("student not found")
	}

	return nil
}

