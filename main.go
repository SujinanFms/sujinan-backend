//Todo: 1.สร้างไฟล์ main.go
//package
package main 

//import lib
import (
	  "fmt"
    "log"
    "net/http"
    "github.com/SujinanFms/sujinan-backend/routes"
)

//Func หลักการทำงานของ Go
func main() {
	fmt.Println("Hello Suji")
	//คำสั่งต่างๆ
    router := routes.SetupRouter()
    log.Println("Server running at :8080")
    http.ListenAndServe(":8080", router)
}
