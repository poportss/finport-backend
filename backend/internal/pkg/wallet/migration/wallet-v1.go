package migration

import "gorm.io/gorm"

type v1 struct{}

func (v *v1) Migrate(db *gorm.DB) error {
	createWalletSQL := `
CREATE TABLE IF NOT EXISTS wallets (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    user_id BIGSERIAL REFERENCES USERS(id),
    balance INTEGER NOT NULL DEFAULT 0);`

	createUserWalletTypeSQL := `
CREATE TABLE IF NOT EXISTS wallet_user_types (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    wallet_type BIGSERIAL REFERENCES wallet_types(id),
    wallet_id BIGSERIAL REFERENCES wallets(id)
    )`

	createWalletTypeSQL := `
CREATE TABLE IF NOT EXISTS wallet_types (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    name varchar(255) REFERENCES,
    )`

	// Executando todas as migrations em ordem correta
	if err := db.Exec(createWalletSQL).Error; err != nil {
		return err
	}

	if err := db.Exec(createUserWalletTypeSQL).Error; err != nil {
		return err
	}

	if err := db.Exec(createWalletTypeSQL).Error; err != nil {
		return err
	}

	return nil
}
