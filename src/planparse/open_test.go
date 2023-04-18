package planparse

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test PlanParse.Open() with a valid file path
func TestPlanParse_Open(t *testing.T) {
	// Create a new PlanParse struct with a valid file path
	planParse := PlanParse{
		FilePath: "../../testdata/1.3.1/valid-tfplan.json",
	}

	// Open the file at the FilePath
	plan, err := planParse.Open()

	// Assert that the plan struct is not nil
	assert.NotNil(t, plan)

	// Assert that the error is nil
	assert.Nil(t, err)
}

// Test PlanParse.Open() with an invalid file path
func TestPlanParse_Open_InvalidFilePath(t *testing.T) {
	// Create a new PlanParse struct with an invalid file path
	planParse := PlanParse{
		FilePath: "../../testdata/1.3.1/missing-tfplan.json",
	}

	// Open the file at the FilePath
	_, err := planParse.Open()

	// Custom error for open ../../testdata/1.3.1/invalid-tfplan.json: no such file or directory

	assert.ErrorContains(t, err, "no such file or directory")
}

// Test PlanParse.Open() with a file path that is not a plan file
func TestPlanParse_Open_NotAPlanFile(t *testing.T) {
	// Create a new PlanParse struct with a file path that is not a plan file
	planParse := PlanParse{
		FilePath: "../../testdata/1.3.1/plan.tfplan",
	}

	// Open the file at the FilePath
	_, err := planParse.Open()

	// It's not a JSON file, so it should return an error
	assert.ErrorContains(t, err, "file at ../../testdata/1.3.1/plan.tfplan is not a json file, ensure that the plan was converterd with terraform show -json")
}

// Test PlanParse.Open() with a file path that is a plan file but is not a valid plan file
func TestPlanParse_Open_InvalidPlanFile(t *testing.T) {
	// Create a new PlanParse struct with a file path that is a plan file but is not a valid plan file
	planParse := PlanParse{
		FilePath: "../../testdata/1.3.1/invalid-tfplan.json",
	}

	// Open the file at the FilePath
	_, err := planParse.Open()
	fmt.Printf("err: %v\n", err)
	assert.Error(t, err)
	assert.ErrorContains(t, err, `invalid character '"' after object key:value pair`)
}
