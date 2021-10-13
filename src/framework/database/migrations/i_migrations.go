package migrations

import "gorm.io/gorm"

type IMigrations interface {
	Up() error
	Down() error
}

func _addColumnsToTable(db *gorm.DB, dst interface{}, column string) error {
	if !db.Migrator().HasColumn(dst, column) {
		if err := db.Migrator().AddColumn(dst, column); err != nil {
			return err
		}
	}
	return nil
}