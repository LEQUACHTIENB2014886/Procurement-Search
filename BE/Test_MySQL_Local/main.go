package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Procurement struct {
	STT       int    `json:"stt"`
	BuyNo     string `json:"buyNo"`
	GSBH      string `json:"gsbh"`
	DDBH      string `json:"ddbh"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	DiffDay   int    `json:"diffDay"`
}

func main() {
	// Kết nối MySQL
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/procurement_search")
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	defer db.Close()

	// Kiểm tra kết nối
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database: ", err)
	}

	// Tạo router Gin
	r := gin.Default()

	// Cấu hình CORS
	r.Use(cors.Default()) // Để cho phép các request từ mọi nguồn

	// Endpoint để trả về dữ liệu Procurement với tính năng tìm kiếm và lọc
	r.GET("/api/procurement", func(c *gin.Context) {
		var procurements []Procurement
		buyNo := c.DefaultQuery("buyNo", "")
		gsbh := c.DefaultQuery("gsbh", "")

		// Câu lệnh SQL để tìm kiếm theo BUYNO và GSBH
		query := "SELECT STT, BUYNO, DDBH, GSBH, StartDate, EndDate, DiffDay FROM procurement_search WHERE 1=1"
		var args []interface{}

		if buyNo != "" {
			query += " AND BUYNO LIKE ?"
			args = append(args, "%"+buyNo+"%")
		}

		if gsbh != "" {
			query += " AND GSBH = ?"
			args = append(args, gsbh)
		}

		// Truy vấn dữ liệu từ cơ sở dữ liệu
		rows, err := db.Query(query, args...)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching data from database"})
			return
		}
		defer rows.Close()

		// Lấy dữ liệu từ bảng
		for rows.Next() {
			var p Procurement
			if err := rows.Scan(&p.STT, &p.BuyNo, &p.DDBH, &p.GSBH, &p.StartDate, &p.EndDate, &p.DiffDay); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning data"})
				return
			}
			procurements = append(procurements, p)
		}

		// Kiểm tra lỗi sau khi duyệt các dòng
		if err := rows.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating over rows"})
			return
		}

		// Trả về dữ liệu JSON
		c.JSON(http.StatusOK, procurements)
	})

	// Chạy server trên port 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
