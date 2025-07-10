package seed

import (
	"fmt"
	"myapp/models"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SeedPermissions(db *gorm.DB, e *echo.Echo) {
	routes := e.Routes()
	// 1. Avval hamma tegishli jadvalni tozalaymiz

	// db.Exec("TRUNCATE TABLE role_permissions RESTART IDENTITY CASCADE")
	// db.Exec("TRUNCATE TABLE permissions RESTART IDENTITY CASCADE")
	// db.Exec("TRUNCATE TABLE permission_groups RESTART IDENTITY CASCADE")
	// db.Exec("TRUNCATE TABLE roles RESTART IDENTITY CASCADE")

	// 1. PermissionGroup larni aniqlash
	groupMap := map[string]models.PermissionGroup{}
	for _, route := range routes {
		segments := strings.Split(route.Path, "/")
		if len(segments) > 3 {
			groupName := strings.Title(segments[3]) // Masalan: category
			if _, ok := groupMap[groupName]; !ok {
				pg := models.PermissionGroup{Name: groupName, IsActive: true}
				db.Where("name = ?", groupName).FirstOrCreate(&pg)
				groupMap[groupName] = pg
			}
		}
	}

	// 2. Har bir route uchun permission yaratish
	for _, route := range routes {
		key := fmt.Sprintf("%s:%s", route.Method, route.Path)
		name := generatePermissionName(route.Method, route.Path)
		groupName := strings.Title(strings.Split(route.Path, "/")[3])
		group := groupMap[groupName]

		permission := models.Permission{
			Name:              name,
			Key:               key,
			IsActive:          true,
			PermissionGroupID: group.ID,
		}

		db.Where("key = ?", key).FirstOrCreate(&permission)
	}

	// 3. Role lar
	admin := models.Role{Name: "admin", IsActive: true}
	moderator := models.Role{Name: "moderator", IsActive: true}
	user := models.Role{Name: "user", IsActive: true}
	db.Where("name = ?", "admin").FirstOrCreate(&admin)
	db.Where("name = ?", "moderator").FirstOrCreate(&moderator)
	db.Where("name = ?", "user").FirstOrCreate(&user)

	// 4. RolePermission biriktirish
	var allPermissions []models.Permission
	db.Find(&allPermissions)

	for _, p := range allPermissions {

		// admin – barcha ruxsatlarga ega
		db.FirstOrCreate(&models.RolePermission{}, models.RolePermission{
			RoleID:       admin.ID,
			PermissionID: p.ID,
		})

		// moderator – faqat category va product
		if strings.Contains(p.Key, "/category") || strings.Contains(p.Key, "/product") {
			db.FirstOrCreate(&models.RolePermission{}, models.RolePermission{
				RoleID:       moderator.ID,
				PermissionID: p.ID,
			})
		}

		// user – faqat category GET POST
		if strings.Contains(p.Key, "/category") &&
			(strings.HasPrefix(p.Key, "GET") || strings.HasPrefix(p.Key, "POST")) {
			db.FirstOrCreate(&models.RolePermission{}, models.RolePermission{
				RoleID:       user.ID,
				PermissionID: p.ID,
			})
		}
	}

}

func generatePermissionName(method, path string) string {
	entity := strings.Split(path, "/")
	if len(entity) > 3 {
		entityName := strings.Title(entity[3])
		switch method {
		case "GET":
			if strings.Contains(path, ":id") {
				return entityName + " View"
			}
			return entityName + " List"
		case "POST":
			return entityName + " Create"
		case "PUT", "PATCH":
			return entityName + " Update"
		case "DELETE":
			return entityName + " Delete"
		}
	}
	return method + " " + path
}
