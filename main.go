// Todo: 1.สร้างไฟล์ main.go
//package
package main 

//import lib
import (
	"fmt"
    "log"
    "net/http"

    "github.com/SujinanFms/sujinan-backend/routes"
    "github.com/SujinanFms/sujinan-backend/database"
    "github.com/gorilla/handlers"
    "github.com/joho/godotenv"
)

//Func หลักการทำงานของ Go
func main() {
    err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

    	fmt.Println("Hello Suji")

    database.ConnectDB()
    // database.Migrate()


	//คำสั่งต่างๆ
    log.Println("Server running at :8080")
    // http.ListenAndServe(":8080", router)

      router := routes.SetupRouter()

    // ห่อ router ด้วย CORS middleware
  corsHandler := handlers.CORS(
        handlers.AllowedOrigins([]string{"http://localhost:3000"}), // frontend origin
        handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}), // OPTIONS สำคัญมาก
        handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
        handlers.AllowCredentials(),
    )(router)

    log.Println("Server running at :8080")
    log.Fatal(http.ListenAndServe(":8080", corsHandler))
}
