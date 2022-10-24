package controller

import (
	"net/http"

	"github.com/Nanthat/sa-65-example/entity"
	"github.com/gin-gonic/gin"
)

// POST /watch_videos
func CreateOrder(c *gin.Context) {

	var order entity.Order
	var product entity.Product
	var cart entity.Cart

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร watchVideo
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา video ด้วย id
	if tx := entity.DB().Where("id = ?", order.ProductID).First(&order); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product not found"})
		return
	}

	// 10: ค้นหา resolution ด้วย id
	if tx := entity.DB().Where("id = ?", order.CartID).First(&cart); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cart not found"})
		return
	}

	// 12: สร้าง WatchVideo
	od := entity.Order{
		Product:  product, // โยงความสัมพันธ์กับ Entity Resolution
		Quantity: order.Quantity,
		Cart:     cart, // โยงความสัมพันธ์กับ Entity Video // ตั้งค่าฟิลด์ watchedTime
	}

	// 13: บันทึก
	if err := entity.DB().Create(&od).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": od})
}

// GET /watchvideo/:id
func GetOrder(c *gin.Context) {
	var order entity.Order
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&order); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "order not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": order})
}

// GET /watch_videos
func ListOrders(c *gin.Context) {
	//var orders []entity.Order
	result := []map[string]interface{}{}
	entity.DB().Table("orders").
		Select("order.id, order.cart_id").
		Joins("left join cart on cart.id = orders.cart_id").
		Find(&result)
	c.JSON(http.StatusOK, gin.H{"data": result})

}

// DELETE /watch_videos/:id
func DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM orders WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "order not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /watch_videos
func UpdateOrder(c *gin.Context) {
	var order entity.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", order.ID).First(&order); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "order not found"})
		return
	}

	if err := entity.DB().Save(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}
