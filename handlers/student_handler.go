package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"example.com/student-api/models"
	"example.com/student-api/services"
)

type StudentHandler struct {
	Service *services.StudentService
}

func (h *StudentHandler) GetStudents(c *gin.Context) {
	students, err := h.Service.GetStudents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch students"})
		return
	}
	c.JSON(http.StatusOK, students)
}

func (h *StudentHandler) GetStudentByID(c *gin.Context) {
	id := c.Param("id")
	student, err := h.Service.GetStudentByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	c.JSON(http.StatusOK, student)
}

func (h *StudentHandler) CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := h.Service.CreateStudent(student); err != nil {
		// Validation error จาก service ส่งกลับเป็น 400 <-- 
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, student)
}

// UpdateStudent รับ PUT /students/:id
// อัปเดต name, major, gpa ของนักศึกษาตาม id
func (h *StudentHandler) UpdateStudent(c *gin.Context) {
	id := c.Param("id") // ดึง id จาก URL เช่น /students/66090003

	var student models.Student
	// ShouldBindJSON อ่าน JSON body และแปลงเป็น struct
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	updated, err := h.Service.UpdateStudent(id, student)
	if err != nil {
		// ตรวจสอบว่า error มาจาก "not found" หรือ validation
		if strings.Contains(err.Error(), "not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		} else {
			// เป็น validation error เช่น gpa ไม่ถูกต้อง
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, updated)
}

// DeleteStudent รับ DELETE /students/:id
// ลบนักศึกษาตาม id และส่ง 204 No Content กลับ
func (h *StudentHandler) DeleteStudent(c *gin.Context) {
	id := c.Param("id")

	if err := h.Service.DeleteStudent(id); err != nil {
		if strings.Contains(err.Error(), "not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete student"})
		}
		return
	}

	// 204 No Content = สำเร็จ แต่ไม่มี body ส่งกลับ
	c.Status(http.StatusNoContent)
}