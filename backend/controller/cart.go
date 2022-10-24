package controller

import (
	"net/http"

	"github.com/Nanthat/sa-65-example/entity"
	"github.com/gin-gonic/gin"
)

// POST /carts
func CreateCart(c *gin.Context) {
	var cart entity.Cart
	var employee entity.Employee
	var member entity.Member

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร watchVideo
	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา video ด้วย id
	if tx := entity.DB().Where("id = ?", cart.EmployeeID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}
	// 10: ค้นหา resolution ด้วย id
	if tx := entity.DB().Where("id = ?", cart.MemberID).First(&member); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "member not found"})
		return
	}
	// 12: สร้าง WatchVideo
	ca := entity.Cart{
		Employee: employee, // โยงความสัมพันธ์กับ Entity Employee
		Member:   member,   // โยงความสัมพันธ์กับ Entity Member

	}

	// 13: บันทึก
	if err := entity.DB().Create(&ca).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": ca})
}

// GET /cart/:id
func GetCart(c *gin.Context) {
	var cart entity.Cart
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&cart); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cart not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cart})
}

// GET /carts
func ListCarts(c *gin.Context) {
	var carts []entity.Cart
	if err := entity.DB().Preload("Employee").Preload("Member").Raw("SELECT * FROM carts").Find(&carts).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": carts})
}

// DELETE /carts/:id
func DeleteCart(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM carts WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cart not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /carts
func UpdateCart(c *gin.Context) {
	var cart entity.Cart
	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", cart.ID).First(&cart); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cart not found"})
		return
	}

	if err := entity.DB().Save(&cart).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cart})
}
