package models

import "time"

type TestExecuteResultModel struct {
	Id               string                 `json:"_id" form:"_id"`
	TestRunNumber    int                    `json:"test_run_number" form:"test_run_number"`
	TestId           string                 `json:"test_id" form:"test_id"`
	StatusCode       int                    `json:"status_code" form:"status_code"`
	UpdatedAt        time.Time              `json:"updated_at" form:"updated_at"`
	CreatedAt        time.Time              `json:"created_at" form:"created_at"`
	TestRunInstances []TestRunInstanceModel `json:"test_run_instances" form:"test_run_instances"`
}
	
type TestRunInstanceModel struct {
	Id string `json:"_id" form:"_id"`
}
