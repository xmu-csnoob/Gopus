package mapper

import (
	"Opus/src/main/infrastructure/database"
	"Opus/src/main/mvc/model"
)

type ImageMapper struct {
}

func (imageMapper *ImageMapper) FindById(id int) model.Image {
	image := new(model.Image)
	database.MysqlClientInstance.GormDBClient.First(&image, id)
	return *image
}
