package planparse

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	tfjson "github.com/hashicorp/terraform-json"
)

// Open a plan JSON file and return a plan struct
func (p PlanParse) Open() (tfjson.Plan, error) {

	// Attempt to open a file at the FilePath
	var plan tfjson.Plan

	// Check if the file exists
	_, err := filepath.Abs(p.FilePath)
	if err != nil {
		return plan, err
	}

	// Check if the file is json
	if filepath.Ext(p.FilePath) != ".json" {
		return plan, fmt.Errorf("file at %s is not a json file, ensure that the plan was converterd with terraform show -json", p.FilePath)
	}

	// Open the file at the FilePath as bytes
	bytesFromPlan, err := ioutil.ReadFile(p.FilePath)
	if err != nil {
		return plan, err
	}

	// Unmarshal the bytes into the plan struct
	err = plan.UnmarshalJSON(bytesFromPlan)
	if err != nil {
		return plan, err
	}

	if plan.Validate() != nil {
		return plan, fmt.Errorf("file at %s is not a valid plan file, ensure that the plan was converterd with terraform show -json", p.FilePath)
	}

	return plan, nil
}
