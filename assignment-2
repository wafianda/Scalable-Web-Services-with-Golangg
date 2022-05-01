package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/hacktiv?charset=utf8mb4&parseTime=True&loc=Local")
	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	} else {
		log.Println("Connection successful!")
	}
	defer db.Close()
	router := gin.Default()

	type Orders_by struct {
		Itemid       string     `json: "itemid"`
		Itemcode     string     `json: "itemcode"`
		Description  string     `json: "description"`
		Quantity     string     `json: "quantity"`
		Orderid      string     `json: "orderid"`
		Customername string     `json: "customername"`
		Orderedat    *time.Time `json: "orderedat"`
	}

	// Buat Panggil Data Spesifik
	router.GET("orders/:itemid", func(c *gin.Context) {
		var (
			orders_by Orders_by
			result    gin.H
		)
		itemid := c.Param("itemid")
		row := db.QueryRow("select itemid, itemcode, description, quantity, orderid, customername, orderedat from orders_by where itemid = ?;", itemid)
		err := row.Scan(&orders_by.Itemid, &orders_by.Itemcode, &orders_by.Description, &orders_by.Quantity, &orders_by.Orderid, &orders_by.Customername, &orders_by.Orderedat)
		if err != nil {
			result = gin.H{
				"Hasil": "Data Tidak Ditemukan!",
				"Total": 0,
			}
		} else {
			result = gin.H{
				"hasil": orders_by,
				"Total": 1,
			}
		}
		c.JSON(http.StatusOK, result)
	})

	// Buat Panggil Semua Data
	router.GET("orders/", func(c *gin.Context) {
		var (
			orders_by  Orders_by
			orders_bys []Orders_by
		)
		rows, err := db.Query("select itemid, itemcode, description, quantity, orderid, customername, orderedat from orders_by;")
		if err != nil {
			fmt.Print(err.Error())
		}
		for rows.Next() {
			err = rows.Scan(&orders_by.Itemid, &orders_by.Itemcode, &orders_by.Description, &orders_by.Quantity, &orders_by.Orderid, &orders_by.Customername, &orders_by.Orderedat)
			orders_bys = append(orders_bys, orders_by)
			if err != nil {
				fmt.Print(err.Error())
			}
		}
		defer rows.Close()
		c.JSON(http.StatusOK, gin.H{
			"hasil": orders_bys,
			"Total": len(orders_bys),
		})
	})

	// Buat Tambah Data
	router.POST("orders/", func(c *gin.Context) {
		var buffer bytes.Buffer
		itemid := c.PostForm("itemid")
		itemcode := c.PostForm("itemcode")
		description := c.PostForm("description")
		quantity := c.PostForm("quantity")
		orderid := c.PostForm("orderid")
		customername := c.PostForm("customername")
		orderedat := c.PostForm("orderedat")
		stmt, err := db.Prepare("insert into orders_by (itemid, itemcode, description, quantity, orderid, customername, orderedat) values(?,?,?,?,?,?,?);")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(itemid, itemcode, description, quantity, orderid, customername, orderedat)

		if err != nil {
			fmt.Print(err.Error())
		}

		buffer.WriteString(itemcode)
		buffer.WriteString(description)
		buffer.WriteString(quantity)
		buffer.WriteString(orderid)
		buffer.WriteString(customername)
		buffer.WriteString(orderedat)
		defer stmt.Close()
		data := buffer.String()
		c.JSON(http.StatusOK, gin.H{
			"Pesan": fmt.Sprintf("Data Berhasil Ditambahkan %s ", data),
		})
	})

	// Buat Update Data
	router.PUT("orders/", func(c *gin.Context) {
		var buffer bytes.Buffer
		itemid := c.PostForm("itemid")
		itemcode := c.PostForm("itemcode")
		description := c.PostForm("description")
		quantity := c.PostForm("quantity")
		orderid := c.PostForm("orderid")
		customername := c.PostForm("customername")
		orderedat := c.PostForm("orderedat")
		stmt, err := db.Prepare("update orders_by set itemcode= ?, description= ?, quantity= ?, orderid= ?, customername= ?, orderedat= ? where itemid= ?;")

		if err != nil {
			fmt.Println(err.Error())
		}
		_, err = stmt.Exec(itemcode, description, quantity, orderid, customername, orderedat, itemid)
		if err != nil {
			fmt.Println(err.Error())
		}

		buffer.WriteString(itemcode)
		buffer.WriteString(description)
		buffer.WriteString(quantity)
		buffer.WriteString(orderid)
		buffer.WriteString(customername)
		buffer.WriteString(orderedat)
		defer stmt.Close()
		data := buffer.String()
		c.JSON(http.StatusOK, gin.H{
			"Pesan": fmt.Sprintf("Data Berhasil Diubah Menjadi %s ", data),
		})
	})

	// Buat Delete Data
	router.DELETE("orders/", func(c *gin.Context) {
		itemid := c.PostForm("itemid")
		stmt, err := db.Prepare("delete from orders_by where itemid= ?;")
		if err != nil {
			fmt.Println(err.Error())
		}
		_, err = stmt.Exec(itemid)
		if err != nil {
			fmt.Print(err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"Pesan": fmt.Sprintf("Data Berhasil Terhapus! %s", itemid),
		})
	})

	router.Run(":8080")
}
