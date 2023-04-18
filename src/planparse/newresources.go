package planparse

import tfjson "github.com/hashicorp/terraform-json"

type NewResource struct {
	Type string // The type of resource
	Name string // The name of the resource
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
