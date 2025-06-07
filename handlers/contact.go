// /sujinan-backend/handlers/contact.go
package handlers

import (
	"fmt"
	"encoding/json"
	"net/http"
	"net/smtp" //send to gmail
	"os"
	"github.com/SujinanFms/sujinan-backend/models"
	"github.com/SujinanFms/sujinan-backend/database"
)


func HandleContact(w http.ResponseWriter, r *http.Request) {
	var contact models.Contact

	// Decode JSON จาก body
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Debug log: ดูข้อมูล contact ที่รับมา
	fmt.Printf("Received contact: %+v\n", contact)

	// บันทึกลงฐานข้อมูล
	result, err := database.DB.Exec(
		"INSERT INTO contacts (name, company, email, phone, address) VALUES ($1, $2, $3, $4, $5)",
		contact.Name, contact.Company, contact.Email, contact.Phone, contact.Address,
	)

	if err != nil {
		http.Error(w, "Failed to save contact", http.StatusInternalServerError)
		return
	}

	// 🟩 เรียกฟังก์ชันส่งอีเมลหลังจาก INSERT สำเร็จ
	if err := sendEmail(contact); err != nil {
		fmt.Println("❌ Error sending email:", err)
	} else {
		fmt.Println("✅ Email sent successfully")
	}

	// Debug log
	fmt.Printf("Received contact: %+v\n", contact)

	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("Rows affected: %d\n", rowsAffected)

	// ส่ง response กลับ
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Contact received",
	})
}


//send to gmail
func sendEmail(contact models.Contact) error {
from := os.Getenv("EMAIL_FROM")
password := os.Getenv("EMAIL_PASSWORD")


	to := []string{from}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte(fmt.Sprintf(
		"Subject: New Contact Submission\r\n" +
		"Reply-To: %s\r\n" +
		"Content-Type: text/plain; charset=\"UTF-8\"\r\n" +
		"\r\n" +
		"Name: %s\nEmail: %s\nCompany: %s\nPhone: %s\nAddress: %s",
		contact.Email, contact.Name, contact.Email, contact.Company, contact.Phone, contact.Address,
	))



	// message := []byte(fmt.Sprintf(
	// 	"Subject: New Contact Submission\r\n\r\nName: %s\nEmail: %s\nCompany: %s\nPhone: %s\nAddress: %s",
	// 	contact.Name, contact.Email, contact.Company, contact.Phone, contact.Address,
	// ))

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		return err
	}

	fmt.Println("✅ Email sent successfully")
	return nil
}




// package handlers

// import (
// 	"encoding/json"
// 	"net/http"
// )

// type ContactRequest struct {
// 	Name  string `json:"name"`
// 	Email string `json:"email"`
// }

// func HandleContact(w http.ResponseWriter, r *http.Request) {
// 	var contact ContactRequest
// 	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
// 		http.Error(w, "Invalid request", http.StatusBadRequest)
// 		return
// 	}
// 	json.NewEncoder(w).Encode(map[string]string{
// 		"message": "Contact received",
// 	})
// }
