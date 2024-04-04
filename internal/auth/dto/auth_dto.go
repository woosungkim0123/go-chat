package dto

type LoginDto struct {
	ID   int    `json:"ID"`
	Name string `json:"name"`
}

type ProfileDto struct {
	ID           int    `json:"ID"`
	Name         string `json:"name"`
	ProfileImage string `json:"profileImage"`
}
