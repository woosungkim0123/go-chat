package domain

type User struct {
	ID           int    `json:"ID"`
	Name         string `json:"name"`
	LoginID      string `json:"loginID"`
	ProfileImage string `json:"profileImage"`
}
