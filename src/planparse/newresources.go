package planparse

import tfjson "github.com/hashicorp/terraform-json"

type NewResource struct {
	Type string // The type of resource
	Name string // The name of the resource
}

// IsAllowed checks if the resource is allowed to be created
func (r NewResource) IsAllowed(allowList []string) bool {
	// If the resource is not in the allow list return false
	for _, allowedResource := range allowList {
		if r.Type == allowedResource {
			return true
		}
	}

	return false
}

// GetNewResources returns a slice of resources that are new in the plan
func GetNewResources(plan tfjson.Plan) []NewResource {
	var newResources []NewResource
	if plan.ResourceChanges != nil {
		for _, resource := range plan.ResourceChanges {
			if resource.Change.Actions.Create() {
				newResources = append(newResources, NewResource{
					Type: resource.Type,
					Name: resource.Name,
				})
			}
		}
	}
	return newResources
}
