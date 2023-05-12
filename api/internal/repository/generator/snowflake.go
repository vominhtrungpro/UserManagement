package generator

import (
	"vominhtrungpro/usermanagement/internal/pkg/snowflake"
)

var (
	ProductSNF snowflake.SnowflakeGenerator
)

func InitSnowflakeGenerators() error {
	ProductSNF = snowflake.New()

	return nil
}
