package models

import "github.com/google/uuid"

type cv struct {
	Title string `json:"title"`
}

type cvInformation struct {
	FullName  string  `json:"full_name"`
	Photo     *[]byte `json:"photo"`
	Position  *string `json:"position"`
	Location  *string `json:"location"`
	Biography *string `json:"biography"`
}

type cvContact struct {
	Title string  `json:"title"`
	Link  *string `json:"link"`
}

type cvSkill struct {
	Name string `json:"name"`
}

type cvLanguage struct {
	Name  string `json:"name"`
	Level string `json:"level"`
}

type cvWorkExperience struct {
	Company     string  `json:"company"`
	Position    string  `json:"position"`
	StartDate   string  `json:"start_date"`
	EndDate     *string `json:"end_date"`
	Location    string  `json:"location"`
	Description string  `json:"description"`
}

type cvEducation struct {
	Institution string  `json:"institution"`
	StartDate   string  `json:"start_date"`
	EndDate     *string `json:"end_date"`
	Location    string  `json:"location"`
	Description *string `json:"description"`
	Faculty     string  `json:"faculty"`
	Degree      *string `json:"degree"`
}

type cvCertificate struct {
	Title       string  `json:"title"`
	Vendor      string  `json:"vendor"`
	StartDate   string  `json:"start_date"`
	EndDate     *string `json:"end_date"`
	Description *string `json:"description"`
}

type CvInfo struct {
	UserID          uuid.UUID           `json:"user_id"`
	CvID            uuid.UUID           `json:"cv_id"`
	CV              cv                  `json:"cv"`
	Information     *cvInformation      `json:"information"`
	Contacts        []*cvContact        `json:"contacts"`
	Skills          []*cvSkill          `json:"skills"`
	Languages       []*cvLanguage       `json:"languages"`
	WorkExperiences []*cvWorkExperience `json:"work_experiences"`
	Educations      []*cvEducation      `json:"educations"`
	Certificates    []*cvCertificate    `json:"certificates"`
}
