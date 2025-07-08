package migration

import "gorm.io/gorm"

type v1 struct{}

func (v *v1) Migrate(db *gorm.DB) error {

	createUsersSQL := `
	CREATE TABLE IF NOT EXISTS users (
    	id BIGSERIAL PRIMARY KEY,
    	created_at TIMESTAMP WITH TIME ZONE,
    	updated_at TIMESTAMP WITH TIME ZONE,
    	deleted_at TIMESTAMP WITH TIME ZONE,
    	name VARCHAR(255) NOT NULL,
    	document VARCHAR(255) NOT NULL,
    	password VARCHAR(255) NOT NULL
	)`

	if err := db.Exec(createUsersSQL).Error; err != nil {
		return err
	}

	return nil
}
