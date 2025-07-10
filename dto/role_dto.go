package dto

import "myapp/models"

type CreateRoleRequest struct {
	Name          string `json:"name" validate:"required"`
	IsActive      bool   `json:"is_active"`
	PermissionIDs []uint `json:"permission_ids"`
}

type UpdateRoleRequest struct {
	Name          string `json:"name"`
	IsActive      bool   `json:"is_active"`
	PermissionIDs []uint `json:"permission_ids"`
}

type PermissionItem struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Key      string `json:"key"`
	IsActive bool   `json:"is_active"`
}

type RoleResponse struct {
	ID          uint             `json:"id"`
	Name        string           `json:"name"`
	IsActive    bool             `json:"is_active"`
	Permissions []PermissionItem `json:"permissions"`
}

func ToRoleResponse(role models.Role) RoleResponse {
	var permissions []PermissionItem
	for _, perm := range role.Permissions {
		permissions = append(permissions, PermissionItem{
			ID:       perm.ID,
			Name:     perm.Name,
			Key:      perm.Key,
			IsActive: perm.IsActive,
		})
	}

	return RoleResponse{
		ID:          role.ID,
		Name:        role.Name,
		IsActive:    role.IsActive,
		Permissions: permissions,
	}
}
