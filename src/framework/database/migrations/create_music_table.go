package migrations

import (
	music_models "github.com/mixnote/mixnote-api-go/src/music/model"
	"gorm.io/gorm"
)

type (
	createArtistTable struct {
		DB *gorm.DB
	}

	createAlbumTable struct {
		DB *gorm.DB
	}

	createSongTable struct {
		DB *gorm.DB
	}
)


func CreateArtistTable(db *gorm.DB) (c *createArtistTable) {
	c = &createArtistTable{
		DB: db,
	}
	return
}

func CreateAlbumTable(db *gorm.DB) (c *createAlbumTable) {
	c = &createAlbumTable{
		DB: db,
	}
	return
}

func CreateSongTable(db *gorm.DB) (c *createSongTable) {
	c = &createSongTable{
		DB: db,
	}
	return
}


func (c *createArtistTable) Up() error {
	return c.DB.Migrator().AutoMigrate(&music_models.Artist{})
}

func (c *createAlbumTable) Up() error {
	return c.DB.Migrator().AutoMigrate(&music_models.Album{})
}

func (c *createSongTable) Up() error {
	return c.DB.Migrator().AutoMigrate(&music_models.Song{})
}



func (c *createArtistTable) Down() error {
	return c.DB.Migrator().DropTable(&music_models.Artist{})
}

func (c *createAlbumTable) Down() error {
	return c.DB.Migrator().DropTable(&music_models.Album{})
}

func (c *createSongTable) Down() error {
	return c.DB.Migrator().DropTable(&music_models.Song{})
}
