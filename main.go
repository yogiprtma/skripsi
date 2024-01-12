package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"project-skripsi/document_legalization"
	"project-skripsi/employee"
	"project-skripsi/render"
	"project-skripsi/subject"
	"project-skripsi/user"
	webHandler "project-skripsi/web/handler"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("MYSQLUSER"), os.Getenv("MYSQLPASSWORD"), os.Getenv("MYSQLHOST"), os.Getenv("MYSQLPORT"), os.Getenv("MYSQLDATABASE"))
	// dsn := "root:@tcp(127.0.0.1:3306)/skripsi?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	subjectRepository := subject.NewRepository(db)
	subjectService := subject.NewService(subjectRepository)

	documentLegalizationRepository := document_legalization.NewRepository(db)
	documentLegalizationService := document_legalization.NewService(documentLegalizationRepository, userRepository)

	employeeRepository := employee.NewRepository(db)
	employeeService := employee.NewService(employeeRepository)

	adminDashboardHandler := webHandler.NewAdminDashboardHandler(documentLegalizationService, employeeService, userService, subjectService)
	adminPerihalHandler := webHandler.NewAdminPerihalHandler(subjectService)
	adminUserHandler := webHandler.NewAdminUserHandler(userService)
	adminPegawaiHandler := webHandler.NewAdminPegawaiHandler(employeeService)
	adminKunciHandler := webHandler.NewAdminKunciHandler(userService)
	adminDokumenHandler := webHandler.NewAdminDokumenHandler(documentLegalizationService)

	karyawanDashboardHandler := webHandler.NewKaryawanDashboardHandler(documentLegalizationService)

	kaprodiDashboardHandler := webHandler.NewKaprodiDashboardHandler(documentLegalizationService)

	wadekDashboardHandler := webHandler.NewWadekDashboardHandler(documentLegalizationService)

	sessionHandler := webHandler.NewSessionHandler(userService, employeeService)
	CheckDocumentHandler := webHandler.NewCheckDocumentHandler(userService, documentLegalizationService)

	userDashboardHandler := webHandler.NewUserDashboardHandler(documentLegalizationService)
	userlegalisasiHandler := webHandler.NewUserLegalisasiHandler(subjectService, documentLegalizationService, userService)

	router := gin.Default()

	cookieStore := cookie.NewStore([]byte("MY-SECRET-KEY"))
	cookieStore.Options(sessions.Options{MaxAge: 60 * 60 * 12})
	router.Use(sessions.Sessions("X-SKRIPSI-COOKIE", cookieStore))

	router.Static("/css", "./web/assets/css")
	router.Static("/img", "./web/assets/img")
	router.Static("/js", "./web/assets/js")
	router.Static("/scss", "./web/assets/scss")
	router.Static("/vendor", "./web/assets/vendor")
	router.Static("/file-document", "./file_document")
	router.Static("/file-signed", "./file_signed")

	router.HTMLRender = render.LoadTemplates("./web/templates")

	router.GET("/admin/dashboard", authAdminMiddleware(), adminDashboardHandler.Index)

	router.GET("/admin/kunci", authAdminMiddleware(), adminKunciHandler.Index)
	router.GET("/admin/kunci/:id", authAdminMiddleware(), adminKunciHandler.Detail)

	router.GET("/admin/dokumen", authAdminMiddleware(), adminDokumenHandler.Index)
	router.GET("/admin/dokumen/:id", authAdminMiddleware(), adminDokumenHandler.Detail)

	router.GET("/admin/akun", authAdminMiddleware(), adminUserHandler.Index)
	router.GET("/admin/akun/new", authAdminMiddleware(), adminUserHandler.New)
	router.POST("/admin/akun", authAdminMiddleware(), adminUserHandler.Create)
	router.GET("/admin/akun/edit/:id", authAdminMiddleware(), adminUserHandler.Edit)
	router.POST("/admin/akun/update/:id", authAdminMiddleware(), adminUserHandler.Update)
	router.POST("/admin/akun/delete/:id", authAdminMiddleware(), adminUserHandler.Delete)
	router.GET("/admin/akun/password/edit/:id", authAdminMiddleware(), adminUserHandler.EditPassword)
	router.POST("/admin/akun/password/update/:id", authAdminMiddleware(), adminUserHandler.UpdatePassword)

	router.GET("/admin/perihal", authAdminMiddleware(), adminPerihalHandler.Index)
	router.GET("/admin/perihal/new", authAdminMiddleware(), adminPerihalHandler.New)
	router.POST("/admin/perihal", authAdminMiddleware(), adminPerihalHandler.Create)
	router.GET("/admin/perihal/edit/:id", authAdminMiddleware(), adminPerihalHandler.Edit)
	router.POST("/admin/perihal/update/:id", authAdminMiddleware(), adminPerihalHandler.Update)
	router.POST("/admin/perihal/delete/:id", authAdminMiddleware(), adminPerihalHandler.Delete)

	router.POST("admin/pegawai", authAdminMiddleware(), adminPegawaiHandler.Create)
	router.GET("/admin/pegawai/new", adminPegawaiHandler.New)
	router.GET("/admin/pegawai", authAdminMiddleware(), adminPegawaiHandler.Index)
	router.GET("/admin/pegawai/edit/:id", authAdminMiddleware(), adminPegawaiHandler.Edit)
	router.POST("/admin/pegawai/update/:id", authAdminMiddleware(), adminPegawaiHandler.Update)
	router.GET("/admin/pegawai/password/edit/:id", authAdminMiddleware(), adminPegawaiHandler.EditPassword)
	router.POST("/admin/pegawai/password/update/:id", authAdminMiddleware(), adminPegawaiHandler.UpdatePassword)
	router.POST("/admin/pegawai/delete/:id", authAdminMiddleware(), adminPegawaiHandler.Delete)

	router.GET("/login", sessionHandler.NewSessionEmployee)
	router.POST("/employee-session", sessionHandler.CreateSessionEmployee)

	// user
	router.GET("/dashboard", authUserMiddleware(), userDashboardHandler.Index)
	router.GET("/dashboard/detail/:id", authUserMiddleware(), userDashboardHandler.Detail)

	router.GET("/pengajuan-legalisasi", authUserMiddleware(), userlegalisasiHandler.Index)
	router.POST("/pengajuan-legalisasi", authUserMiddleware(), userlegalisasiHandler.Create)

	router.GET("/", sessionHandler.New)
	router.POST("/session", sessionHandler.Create)
	router.GET("/logout", sessionHandler.Destroy)

	// Guest
	router.GET("/dokumen/cek-dokumen", CheckDocumentHandler.New)
	router.POST("/dokumen/cek-dokumen", CheckDocumentHandler.Find)
	router.GET("/dokumen/dokumen-tidak-valid", CheckDocumentHandler.Invalid)
	router.GET("/dokumen/:uuid", CheckDocumentHandler.Index)

	// Karyawan Akademik
	router.GET("/karyawan/dashboard", authKaryawanMiddleware(), karyawanDashboardHandler.Index)
	router.GET("/karyawan/dashboard/new/:id", authKaryawanMiddleware(), karyawanDashboardHandler.New)
	router.POST("/karyawan/dashboard/update/:id", authKaryawanMiddleware(), karyawanDashboardHandler.Update)
	router.POST("/karyawan/update", authKaryawanMiddleware(), karyawanDashboardHandler.Reject)

	// Kaprodi
	router.GET("/koordinator-prodi/dashboard", authKaprodiMiddleware(), kaprodiDashboardHandler.Index)
	router.GET("/koordinator-prodi/dashboard/detail/:id", authKaprodiMiddleware(), kaprodiDashboardHandler.Detail)
	router.POST("/koordinator-prodi/update", authKaprodiMiddleware(), kaprodiDashboardHandler.Reject)
	router.POST("/koordinator-prodi/dashboard/update/:id", authKaprodiMiddleware(), kaprodiDashboardHandler.Update)

	// wadek
	router.GET("/wadek/dashboard", authWadekMiddleware(), wadekDashboardHandler.Index)
	router.GET("/wadek/dashboard/detail/:id", authWadekMiddleware(), wadekDashboardHandler.Detail)
	router.POST("/wadek/update", authWadekMiddleware(), wadekDashboardHandler.Reject)
	router.POST("/wadek/dashboard/update/:id", wadekDashboardHandler.Update)

	router.Run()

}

func authAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		userSession := session.Get("authUser")
		adminSession := session.Get("authAdmin")
		wadekSession := session.Get("authWadek")
		kaprodiSession := session.Get("authKaprodi")
		karyawanSession := session.Get("authKaryawan")

		if adminSession == nil {
			c.Redirect(http.StatusFound, "/login")
			return
		}

		if userSession != nil {
			c.Redirect(http.StatusFound, "/dashboard")
			return
		}

		if wadekSession != nil {
			c.Redirect(http.StatusFound, "/wadek/dashboard")
			return
		}

		if kaprodiSession != nil {
			c.Redirect(http.StatusFound, "/koordinator-prodi/dashboard")
			return
		}

		if karyawanSession != nil {
			c.Redirect(http.StatusFound, "/karyawan/dashboard")
			return
		}
	}
}

func authUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		userSession := session.Get("authUser")
		adminSession := session.Get("authAdmin")
		wadekSession := session.Get("authWadek")
		kaprodiSession := session.Get("authKaprodi")
		karyawanSession := session.Get("authKaryawan")

		if adminSession != nil {
			c.Redirect(http.StatusFound, "/admin/dashboard")
			return
		}

		if userSession == nil {
			c.Redirect(http.StatusFound, "/")
			return
		}

		if wadekSession != nil {
			c.Redirect(http.StatusFound, "/wadek/dashboard")
			return
		}

		if kaprodiSession != nil {
			c.Redirect(http.StatusFound, "/koordinator-prodi/dashboard")
			return
		}

		if karyawanSession != nil {
			c.Redirect(http.StatusFound, "/karyawan/dashboard")
			return
		}
	}
}

func authWadekMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		userSession := session.Get("authUser")
		adminSession := session.Get("authAdmin")
		wadekSession := session.Get("authWadek")
		kaprodiSession := session.Get("authKaprodi")
		karyawanSession := session.Get("authKaryawan")

		if adminSession != nil {
			c.Redirect(http.StatusFound, "/admin/dashboard")
			return
		}

		if userSession != nil {
			c.Redirect(http.StatusFound, "/dashboard")
			return
		}

		if wadekSession == nil {
			c.Redirect(http.StatusFound, "/login")
			return
		}

		if kaprodiSession != nil {
			c.Redirect(http.StatusFound, "/koordinator-prodi/dashboard")
			return
		}

		if karyawanSession != nil {
			c.Redirect(http.StatusFound, "/karyawan/dashboard")
			return
		}
	}
}

func authKaprodiMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		userSession := session.Get("authUser")
		adminSession := session.Get("authAdmin")
		wadekSession := session.Get("authWadek")
		kaprodiSession := session.Get("authKaprodi")
		karyawanSession := session.Get("authKaryawan")

		if adminSession != nil {
			c.Redirect(http.StatusFound, "/admin/dashboard")
			return
		}

		if userSession != nil {
			c.Redirect(http.StatusFound, "/dashboard")
			return
		}

		if wadekSession != nil {
			c.Redirect(http.StatusFound, "/wadek/dashboard")
			return
		}

		if kaprodiSession == nil {
			c.Redirect(http.StatusFound, "/login")
			return
		}

		if karyawanSession != nil {
			c.Redirect(http.StatusFound, "/karyawan/dashboard")
			return
		}
	}
}

func authKaryawanMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		userSession := session.Get("authUser")
		adminSession := session.Get("authAdmin")
		wadekSession := session.Get("authWadek")
		kaprodiSession := session.Get("authKaprodi")
		karyawanSession := session.Get("authKaryawan")

		if adminSession != nil {
			c.Redirect(http.StatusFound, "/admin/dashboard")
			return
		}

		if userSession != nil {
			c.Redirect(http.StatusFound, "/dashboard")
			return
		}

		if wadekSession != nil {
			c.Redirect(http.StatusFound, "/wadek/dashboard")
			return
		}

		if kaprodiSession != nil {
			c.Redirect(http.StatusFound, "/koordinator-prodi/dashboard")
			return
		}

		if karyawanSession == nil {
			c.Redirect(http.StatusFound, "/")
			return
		}
	}
}
