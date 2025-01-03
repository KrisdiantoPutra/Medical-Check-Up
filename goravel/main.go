package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"goravel/bootstrap"

	"github.com/gin-gonic/gin"
	"github.com/goravel/framework/facades"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID              uint   `gorm:"primaryKey"`
	Email           string `gorm:"unique;not null"`
	Password        string `gorm:"not null"`
	Age             int    `gorm:"not null"`
	MedicalHistory  string
	Allergies       string
}



func connectDatabase() *gorm.DB {
	// Define your connection string
	dsn := "root:@tcp(localhost:3306)/medical?charset=utf8mb4&parseTime=True&loc=Local"

	// Connect to the database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {	
		log.Fatalf("Error connecting to database: %v", err)
	}

	log.Println("Connected to the database!")
	return db
}

func hashPassword(password string) (string, error) {
	// Menghasilkan salt secara otomatis dan meng-hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}


func main() {
	bootstrap.Boot()
	// Connect to the database
	db := connectDatabase()
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("Error getting database object: %v", err)
		}
		sqlDB.Close()
	}()

	// Create a channel to listen for OS signals
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	router := gin.Default()

	// Menyajikan file statis
	router.LoadHTMLGlob("resources/views/*.html")
	

	router.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(200, "login.html", nil) 
	})

	router.GET("/register", func(ctx *gin.Context) {
		ctx.HTML(200, "register.html", nil) 
	})

	router.GET("/verify", func(ctx *gin.Context) {
		ctx.HTML(200, "verify.html", nil) 
	})

	router.GET("/forgot-password", func(ctx *gin.Context) {
		ctx.HTML(200, "forgot.html", nil) 
	})

	router.GET("/bmi", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", nil) 
	})

	router.GET("/underweight-category", func(ctx *gin.Context) {
		ctx.HTML(200, "underweight.html", nil) 
	})

	router.GET("/normal-category", func(ctx *gin.Context) {
		ctx.HTML(200, "normal.html", nil) 
	})

	router.GET("/overweight-category", func(ctx *gin.Context) {
		ctx.HTML(200, "overweight.html", nil) 
	})

	router.GET("/obese-category", func(ctx *gin.Context) {
		ctx.HTML(200, "obese.html", nil) 
	})

	router.GET("/article1", func(ctx *gin.Context) {
		ctx.HTML(200, "article1.html", nil) 
	})

	router.GET("/article2", func(ctx *gin.Context) {
		ctx.HTML(200, "article2.html", nil) 
	})

	router.GET("/article3", func(ctx *gin.Context) {
		ctx.HTML(200, "article3.html", nil) 
	})

	router.GET("/article4.", func(ctx *gin.Context) {
		ctx.HTML(200, "article4.html", nil) 
	})

	router.GET("/article5", func(ctx *gin.Context) {
		ctx.HTML(200, "article5.html", nil) 
	})

	router.GET("/article6", func(ctx *gin.Context) {
		ctx.HTML(200, "article6.html", nil) 
	})

	router.GET("/article7", func(ctx *gin.Context) {
		ctx.HTML(200, "article7.html", nil) 
	})

	router.GET("/article8", func(ctx *gin.Context) {
		ctx.HTML(200, "article8.html", nil) 
	})

	router.GET("/article9", func(ctx *gin.Context) {
		ctx.HTML(200, "article9.html", nil) 
	})

	router.GET("/artikel1", func(ctx *gin.Context) {
		ctx.HTML(200, "artikel1.html", nil) 
	})

	router.GET("/artikel2.", func(ctx *gin.Context) {
		ctx.HTML(200, "artikel2.html", nil) 
	})

	router.GET("/artikel3", func(ctx *gin.Context) {
		ctx.HTML(200, "artikel3.html", nil) 
	})

	router.GET("/artikel4", func(ctx *gin.Context) {
		ctx.HTML(200, "artikel4.html", nil) 
	})

	router.GET("/artikel5", func(ctx *gin.Context) {
		ctx.HTML(200, "artikel5.html", nil) 
	})

	router.GET("/artikel6", func(ctx *gin.Context) {
		ctx.HTML(200, "artikel6.html", nil) 
	})

	router.GET("/artikel7", func(ctx *gin.Context) {
		ctx.HTML(200, "artikel7.html", nil) 
	})

	router.GET("/artikel8", func(ctx *gin.Context) {
		ctx.HTML(200, "artikel8.html", nil) 
	})

	router.GET("/artikel9", func(ctx *gin.Context) {
		ctx.HTML(200, "artikel9.html", nil) 
	})

	router.GET("/obtikel1", func(ctx *gin.Context) {
		ctx.HTML(200, "obtikel1.html", nil) 
	})

	router.GET("/obtikel2", func(ctx *gin.Context) {
		ctx.HTML(200, "obtikel2.html", nil) 
	})

	router.GET("/obtikel3", func(ctx *gin.Context) {
		ctx.HTML(200, "obtikel3.html", nil) 
	})

	router.GET("/obtikel4", func(ctx *gin.Context) {
		ctx.HTML(200, "obtikel4.html", nil) 
	})

	router.GET("/obtikel5", func(ctx *gin.Context) {
		ctx.HTML(200, "obtikel5.html", nil) 
	})

	router.GET("/obtikel6", func(ctx *gin.Context) {
		ctx.HTML(200, "obtikel6.html", nil) 
	})

	router.GET("/obtikel7", func(ctx *gin.Context) {
		ctx.HTML(200, "obtikel7.html", nil) 
	})

	router.GET("/obtikel8", func(ctx *gin.Context) {
		ctx.HTML(200, "obtikel8.html", nil) 
	})

	router.GET("/obtikel9", func(ctx *gin.Context) {
		ctx.HTML(200, "obtikel9.html", nil) 
	})
	
	router.GET("/dashboard", func(ctx *gin.Context) {
		ctx.HTML(200, "page2.html", nil) 
	})

	router.POST("/dashboard", func(ctx *gin.Context) {
		var user User
		if err := ctx.ShouldBind(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Enkripsi password sebelum menyimpannya
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	
	user.Password = hashedPassword

		if err := db.Create(&user).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create account"})
			return
		}

		// Redirect to page2.html after successful creation
		ctx.Redirect(http.StatusSeeOther, "/dashboard")
	})

	// Start the HTTP server
	go func() {
		if err := router.Run(":2024"); err != nil {
			facades.Log().Errorf("Server Run error: %v", err)
		}
	}()

	// Listen for the OS signal
	go func() {
		<-quit
		if err := facades.Route().Shutdown(); err != nil {
			facades.Log().Errorf("Route Shutdown error: %v", err)
		}
		os.Exit(0)
	}()

	// Block main goroutine
	select {}
}
