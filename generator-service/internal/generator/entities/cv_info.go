package entities

import "github.com/google/uuid"

type cv struct {
	Title string `msgpack:"title"`
}

type cvInformation struct {
	FullName  string  `msgpack:"full_name"`
	Photo     *[]byte `msgpack:"photo"`
	Position  *string `msgpack:"position"`
	Location  *string `msgpack:"location"`
	Biography *string `msgpack:"biography"`
}

type cvContact struct {
	Title string  `msgpack:"title"`
	Link  *string `msgpack:"link"`
}

type cvSkill struct {
	Name string `msgpack:"name"`
}

type cvLanguage struct {
	Name  string `msgpack:"name"`
	Level string `msgpack:"level"`
}

type cvWorkExperience struct {
	Company     string  `msgpack:"company"`
	Position    string  `msgpack:"position"`
	StartDate   string  `msgpack:"start_date"`
	EndDate     *string `msgpack:"end_date"`
	Location    string  `msgpack:"location"`
	Description string  `msgpack:"description"`
}

type cvEducation struct {
	Institution string  `msgpack:"institution"`
	StartDate   string  `msgpack:"start_date"`
	EndDate     *string `msgpack:"end_date"`
	Location    string  `msgpack:"location"`
	Description *string `msgpack:"description"`
	Faculty     string  `msgpack:"faculty"`
	Degree      *string `msgpack:"degree"`
}

type cvCertificate struct {
	Title       string  `msgpack:"title"`
	Vendor      string  `msgpack:"vendor"`
	StartDate   string  `msgpack:"start_date"`
	EndDate     *string `msgpack:"end_date"`
	Description *string `msgpack:"description"`
}

type CvInfo struct {
	UserID          uuid.UUID           `msgpack:"user_id"`
	CvID            uuid.UUID           `msgpack:"cv_id"`
	Template        string              `msgpack:"template"`
	Color           *string             `msgpack:"color"`
	CV              cv                  `msgpack:"cv"`
	Information     *cvInformation      `msgpack:"information"`
	Contacts        []*cvContact        `msgpack:"contacts"`
	Skills          []*cvSkill          `msgpack:"skills"`
	Languages       []*cvLanguage       `msgpack:"languages"`
	WorkExperiences []*cvWorkExperience `msgpack:"work_experiences"`
	Educations      []*cvEducation      `msgpack:"educations"`
	Certificates    []*cvCertificate    `msgpack:"certificates"`
}
