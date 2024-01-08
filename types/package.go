package types

import (
	"github.com/krispekla/pro-profile-ai-api/.gen/postgres/public/model"
)

type PackageListingDTO struct {
	*model.Package
	Pricing *model.PackagePrice `alias:"package_price"`
}
