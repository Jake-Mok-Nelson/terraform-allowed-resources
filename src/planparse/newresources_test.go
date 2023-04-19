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

// Test IsAllowed() with a resource that is allowed to be created
func TestIsAllowed(t *testing.T) {
	// Create a new NewResource struct
	newResource := NewResource{
		Type: "null_resource",
		Name: "resource_one",
	}

	allowList := []string{"null_resource"}

	// Assert that the new resource is allowed to be created
	assert.True(t, newResource.IsAllowed(allowList))
}

// Test IsAllowed() with a resource that is not allowed to be created
func TestIsAllowed_NotAllowed(t *testing.T) {
	// Create a new NewResource struct
	newResource := NewResource{
		Type: "aws_instance",
		Name: "test-instance",
	}

	allowList := []string{"null_resource"}

	// Assert that the new resource is not allowed to be created
	assert.False(t, newResource.IsAllowed(allowList))
}

// Test IsAllowed() where the allow list is empty
func TestIsAllowed_EmptyAllowList(t *testing.T) {
	// Create a new NewResource struct
	newResource := NewResource{
		Type: "aws_instance",
		Name: "test-instance",
	}

	allowList := []string{}

	// Assert that the new resource is not allowed to be created
	assert.False(t, newResource.IsAllowed(allowList))
}

// Test IsAllowed() where the allow list is nil
func TestIsAllowed_NilAllowList(t *testing.T) {
	// Create a new NewResource struct
	newResource := NewResource{
		Type: "aws_instance",
		Name: "test-instance",
	}

	// Assert that the new resource is not allowed to be created
	assert.False(t, newResource.IsAllowed(nil))
}

// Test IsAllowed() where the allow list contains a wildcard
// don't allow wildcared rules to match multiple resources (yet)
func TestIsAllowed_Wildcard(t *testing.T) {
	// Create a new NewResource struct
	newResource := NewResource{
		Type: "aws_instance",
		Name: "test-instance",
	}

	allowList := []string{"*"}

	// Assert that the new resource is allowed to be created
	assert.False(t, newResource.IsAllowed(allowList))
}
