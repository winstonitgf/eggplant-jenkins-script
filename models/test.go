package models

type TestsModel struct {
	Limit  int       `json:"limit" form:"limit"`
	Offset int       `json:"offset" form:"offset"`
	Count  int       `json:"count" form:"count"`
	Tests  TestModel `json:"tests" form:"tests"`
}

type TestModel struct {
	Id string `json:"_id" form:"_id"`
}
