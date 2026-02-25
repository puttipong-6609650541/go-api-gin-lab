package services

import (
	"fmt"

	"example.com/student-api/models"
	"example.com/student-api/repositories"
)

type StudentService struct {
	Repo *repositories.StudentRepository
}

// validateStudent ตรวจสอบข้อมูลก่อน Create หรือ Update
// คืนค่า error ถ้าข้อมูลไม่ถูกต้อง
func validateStudent(student models.Student) error {
	if student.Name == "" {
		return fmt.Errorf("name must not be empty")
	}
	if student.GPA < 0.0 || student.GPA > 4.0 {
		return fmt.Errorf("gpa must be between 0.00 and 4.00")
	}
	return nil
}

func (s *StudentService) GetStudents() ([]models.Student, error) {
	return s.Repo.GetAll()
}

func (s *StudentService) GetStudentByID(id string) (*models.Student, error) {
	return s.Repo.GetByID(id)
}

func (s *StudentService) CreateStudent(student models.Student) error {
	// ตรวจสอบ id ด้วยเพราะเป็น Create ต้องมี id ใหม่ เผื่อลืม
	if student.Id == "" {
		return fmt.Errorf("id must not be empty")
	}
	// เรียก validate ตามที่ได้บอก
	if err := validateStudent(student); err != nil {
		return err
	}
	return s.Repo.Create(student)
}

func (s *StudentService) UpdateStudent(id string, student models.Student) (*models.Student, error) {
	// Update ไม่ต้องตรวจ id ใน body เพราะใช้ id จาก URL param :D
	if err := validateStudent(student); err != nil {
		return nil, err
	}
	return s.Repo.Update(id, student)
}
func (s *StudentService) DeleteStudent(id string) error {
	return s.Repo.Delete(id)
}