// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Principal is navigation property
func (b *UnifiedRoleAssignmentRequestBuilder) Principal() *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/principal"
	return bb
}

// RoleDefinition is navigation property
func (b *UnifiedRoleAssignmentRequestBuilder) RoleDefinition() *UnifiedRoleDefinitionRequestBuilder {
	bb := &UnifiedRoleDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleDefinition"
	return bb
}
