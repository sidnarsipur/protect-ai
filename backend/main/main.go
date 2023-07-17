package main

import (
	"fmt"

	"github.com/sidnarsipur/protect-ai/models"
)

func main() {
	permissions := []models.Permission{
		{
			Agent: "Googlebot",
			AllowedPermission: models.PermissionBody{
				Paths: []string{
					"/",
				},
				FileTypes: models.FileTypes{
					TextFileTypes: []models.TextFileType{
						models.TextFileTypeTxt,
					},
					ImageFileTypes: []models.ImageFileType{
						models.ImageFileTypeJpg,
					},
				},
			},
			DisallowedPermission: models.PermissionBody{
				Paths: []string{
					"/admin",
				},
				FileTypes: models.FileTypes{
					TextFileTypes: []models.TextFileType{
						models.TextFileTypePdf,
					},
					ImageFileTypes: []models.ImageFileType{
						models.ImageFileTypePng,
					},
				},
			},
		},
	}

	permissionString := SetPermissions(permissions)
	fmt.Println(permissionString)

	GetSitemap("https://sidnarsipur.github.io/sitemap.xml")

}
