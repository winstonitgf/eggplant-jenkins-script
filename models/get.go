package models

import "time"

type TestModel struct {
	PageModel
	Tests []TestItemModel `json:"tests" form:"tests"`
}

type PageModel struct {
	Limit  int `json:"limit" form:"limit"`
	Offset int `json:"offset" form:"offset"`
	Count  int `json:"count" form:"count"`
}

type TestItemModel struct {
	Id                    string               `json:"_id" form:"_id"`
	ManagedApplicationId  string               `json:"managed_application_id" form:"managed_application_id"`
	SutId                 string               `json:"sut_id" form:"sut_id"`
	StopOnFailure         bool                 `json:"stop_on_failure" form:"stop_on_failure"`
	AbortTimeout          int                  `json:"abort_timeout" form:"abort_timeout"`
	NotificationEmail     string               `json:"notification_email" form:"notification_email"`
	TestLevelEpfArguments string               `json:"test_level_epf_arguments" form:"test_level_epf_arguments"`
	IsActive              bool                 `json:"is_active" form:"is_active"`
	Version               string               `json:"version" form:"version"`
	Description           string               `json:"description" form:"description"`
	Name                  string               `json:"name" form:"name"`
	UpdatedAt             time.Time            `json:"updated_at" form:"updated_at"`
	CreatedAt             time.Time            `json:"created_at" form:"created_at"`
	TestExecutions        []TestExecutionModel `json:"test_executions" form:"test_executions"`
}

type TestExecutionModel struct {
	Id string `json:"_id" form:"_id"`
}

