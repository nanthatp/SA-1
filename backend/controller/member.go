package controller

import (
	"net/http"

	"github.com/Nanthat/sa-65-example/entity"
	"github.com/gin-gonic/gin"
)

// POST /member
func CreateMember(c *gin.Context) {

	var member entity.Member
	var gender entity.Gender
	var province entity.Province
	var employee entity.Employee

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร gender
	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา gender ด้วย id
	if tx := entity.DB().Where("id = ?", member.GenderID).First(&gender); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gender not found"})
		return
	}

	// 10: ค้นหา province ด้วย id
	if tx := entity.DB().Where("id = ?", member.ProvinceID).First(&province); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "province not found"})
		return
	}

	// 11: ค้นหา employee ด้วย id
	if tx := entity.DB().Where("id = ?", member.EmployeeID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}
	// 12: สร้าง member
	wv := entity.Member{
		FirstName: member.FirstName, // ตั้งค่าฟิลด์
		LastName:  member.LastName,  // ตั้งค่าฟิลด์
		Age:       member.Age,       // ตั้งค่าฟิลด์
		Gender:    gender,           // โยงความสัมพันธ์กับ Entity
		//Date_Of_Birth:    member.Date_Of_Birth,    // ตั้งค่าฟิลด์
		Province:  province,         // โยงความสัมพันธ์กับ Entity
		Telephone: member.Telephone, // ตั้งค่าฟิลด์
		Employee:  employee,         // โยงความสัมพันธ์กับ Entity
	}

	// 13: บันทึก
	if err := entity.DB().Create(&wv).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": wv})
}

// GET /member/:id
func GetMember(c *gin.Context) {
	var member entity.Member
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&member); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "member not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": member})
}

// GET /members
func ListMembers(c *gin.Context) {
	var members []entity.Member
	if err := entity.DB().Preload("Gender").Preload("Province").Preload("Employee").Raw("SELECT * FROM members").Find(&members).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": members})
}

// DELETE /members/:id
func DeleteMember(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM members WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "member not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /members
func UpdateMember(c *gin.Context) {
	var member entity.Member
	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", member.ID).First(&member); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "member not found"})
		return
	}

	if err := entity.DB().Save(&member).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": member})
}
