package accessor

import "github.com/sacloud/libsacloud/v2/sacloud/types"

// DiskPlan ディスクプランのアクセッサ
type DiskPlan interface {
	GetDiskPlanID() types.ID
	SetDiskPlanID(v types.ID)
	GetDiskPlanName() string
	SetDiskPlanName(v string)
	GetDiskPlanStorageClass() string
	SetDiskPlanStorageClass(v string)
}
