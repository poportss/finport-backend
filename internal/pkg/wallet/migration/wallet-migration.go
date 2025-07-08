package migration

import "github.com/poportss/finport-backend/internal/migrations"

func Versions() []migrations.Versions {

	return []migrations.Versions{&v1{}}
}
