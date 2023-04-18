package planparse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test GetNewResources() with a valid plan that contains new resources
func TestGetNewResources(t *testing.T) {
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

	// Get the new resources
	newResources := GetNewResources(plan)

	// Assert that the new resources slice is not nil
	assert.NotNil(t, newResources)

	// Assert that the new resources slice has a length of 1
	assert.Equal(t, len(newResources), 2)

	// Assert that the new resource is of type aws_instance
	assert.Equal(t, newResources[0].Type, "null_resource")

	// Assert that the new resource is named test-instance
	assert.Equal(t, newResources[0].Name, "resource_one")
}
