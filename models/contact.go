package models

type Contact struct {
    Name    string `json:"name"`
    Company string `json:"company"`
    Email   string `json:"email"`
    Phone   string `json:"phone,omitempty"`
    Address string `json:"address,omitempty"`
}

