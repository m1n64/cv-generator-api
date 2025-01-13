package entities

import "github.com/google/uuid"

type CV struct {
	Title string `json:"title"`
}

type CVInformation struct {
	FullName  string  `json:"full_name"`
	Photo     *[]byte `json:"photo"`
	Position  *string `json:"position"`
	Location  *string `json:"location"`
	Biography *string `json:"biography"`
}

type CVContact struct {
	Title string  `json:"title"`
	Link  *string `json:"link"`
}

type CVSkill struct {
	Name string `json:"name"`
}

type CVLanguage struct {
	Name  string `json:"name"`
	Level string `json:"level"`
}

type CVWorkExperience struct {
	Company     string  `json:"company"`
	Position    string  `json:"position"`
	StartDate   string  `json:"start_date"`
	EndDate     *string `json:"end_date"`
	Location    string  `json:"location"`
	Description string  `json:"description"`
}

type CVEducation struct {
	Institution string  `json:"institution"`
	StartDate   string  `json:"start_date"`
	EndDate     *string `json:"end_date"`
	Location    string  `json:"location"`
	Description *string `json:"description"`
	Faculty     string  `json:"faculty"`
	Degree      *string `json:"degree"`
}

type CVCertificate struct {
	Title       string  `json:"title"`
	Vendor      string  `json:"vendor"`
	StartDate   string  `json:"start_date"`
	EndDate     *string `json:"end_date"`
	Description *string `json:"description"`
}

type CvInfo struct {
	UserID          uuid.UUID           `json:"user_id"`
	CvID            uuid.UUID           `json:"cv_id"`
	CV              CV                  `json:"cv"`
	Information     *CVInformation      `json:"information"`
	Contacts        []*CVContact        `json:"contacts"`
	Skills          []*CVSkill          `json:"skills"`
	Languages       []*CVLanguage       `json:"languages"`
	WorkExperiences []*CVWorkExperience `json:"work_experiences"`
	Educations      []*CVEducation      `json:"educations"`
	Certificates    []*CVCertificate    `json:"certificates"`
}
