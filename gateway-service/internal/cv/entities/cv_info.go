package entities

import "github.com/google/uuid"

type CV struct {
	Title string `msgpack:"title"`
}

type CVInformation struct {
	FullName  string  `msgpack:"full_name"`
	Photo     *[]byte `msgpack:"photo"`
	Position  *string `msgpack:"position"`
	Location  *string `msgpack:"location"`
	Biography *string `msgpack:"biography"`
}

type CVContact struct {
	Title string  `msgpack:"title"`
	Link  *string `msgpack:"link"`
}

type CVSkill struct {
	Name string `msgpack:"name"`
}

type CVLanguage struct {
	Name  string `msgpack:"name"`
	Level string `msgpack:"level"`
}

type CVWorkExperience struct {
	Company     string  `msgpack:"company"`
	Position    string  `msgpack:"position"`
	StartDate   string  `msgpack:"start_date"`
	EndDate     *string `msgpack:"end_date"`
	Location    string  `msgpack:"location"`
	Description string  `msgpack:"description"`
}

type CVEducation struct {
	Institution string  `msgpack:"institution"`
	StartDate   string  `msgpack:"start_date"`
	EndDate     *string `msgpack:"end_date"`
	Location    string  `msgpack:"location"`
	Description *string `msgpack:"description"`
	Faculty     string  `msgpack:"faculty"`
	Degree      *string `msgpack:"degree"`
}

type CVCertificate struct {
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
	CV              CV                  `msgpack:"cv"`
	Information     *CVInformation      `msgpack:"information"`
	Contacts        []*CVContact        `msgpack:"contacts"`
	Skills          []*CVSkill          `msgpack:"skills"`
	Languages       []*CVLanguage       `msgpack:"languages"`
	WorkExperiences []*CVWorkExperience `msgpack:"work_experiences"`
	Educations      []*CVEducation      `msgpack:"educations"`
	Certificates    []*CVCertificate    `msgpack:"certificates"`
}
