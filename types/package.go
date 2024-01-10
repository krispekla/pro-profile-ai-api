package types

import (
	"time"

	"github.com/krispekla/pro-profile-ai-api/.gen/postgres/public/model"
)

type PackageListingDTO struct {
	ID          int32                      `sql:"primary_key" alias:"package.id"`
	Name        string                     `alias:"package.name"`
	Description string                     `alias:"package.description"`
	CoverImgURL string                     `alias:"package.cover_img_url"`
	Created     time.Time                  `alias:"package.created"`
	Pricing     *[]model.PackagePrice      `alias:"package_price"`
	Imgs        *[]struct{ ImgURL string } `alias:"package_example_img"`
}

type PackageGeneratedDTO struct {
	ID          int32                        `alias:"generated_package.id" json:"id"`
	Status      model.GeneratedPackageStatus `alias:"generated_package.status" json:"status"`
	CoverImgURL *string                      `alias:"generated_package.cover_img_url" json:"cover_img_url"`
	Updated     time.Time                    `alias:"generated_package.updated" json:"updated"`
	PackageId   *int64                       `alias:"package_order_item.package_id" json:"package_id"`
}
