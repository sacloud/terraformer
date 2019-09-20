package naked

import (
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// PrivateHost 専有ホスト
type PrivateHost struct {
	ID               types.ID         `json:",omitempty" yaml:",omitempty" structs:",omitempty"`
	Name             string           `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description      string           `yaml:"description"`
	Tags             types.Tags       `yaml:"tags"`
	Icon             *Icon            `json:",omitempty" yaml:"icon,omitempty" structs:",omitempty"`
	CreatedAt        *time.Time       `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	Plan             *PrivateHostPlan `json:",omitempty" yaml:"plan,omitempty" structs:",omitempty"`
	Host             *Host            `json:",omitempty" yaml:"host,omitempty" structs:",omitempty"`
	AssignedCPU      int              `json:",omitempty" yaml:"assigned_cpu,omitempty" structs:",omitempty"`
	AssignedMemoryMB int              `json:",omitempty" yaml:"assigned_memory_mb,omitempty" structs:",omitempty"`
}

// PrivateHostPlan 専有ホストプラン
type PrivateHostPlan struct {
	ID           types.ID            `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name         string              `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Class        string              `json:",omitempty" yaml:"class,omitempty" structs:",omitempty"`
	CPU          int                 `json:",omitempty" yaml:"cpu,omitempty" structs:",omitempty"`
	MemoryMB     int                 `json:",omitempty" yaml:"memory_mb,omitempty" structs:",omitempty"`
	ServiceClass string              `json:",omitempty" yaml:"service_class,omitempty" structs:",omitempty"`
	Availability types.EAvailability `json:",omitempty" yaml:"availability,omitempty" structs:",omitempty"`
}
