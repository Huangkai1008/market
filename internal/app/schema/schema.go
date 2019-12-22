package schema

import "market/pkg/utils"

type BaseSchema struct {
	ID        uint           `json:"id"`
	CreatedAt utils.JsonTime `json:"create_time"`
	UpdatedAt utils.JsonTime `json:"update_time"`
}
