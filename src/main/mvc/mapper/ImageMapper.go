package mapper

import (
	"Opus/src/main/infrastructure/database"
	"Opus/src/main/mvc/model"
	"fmt"
)

type ImageMapper struct {
}

func (imageMapper *ImageMapper) FindById(id int) {
	image := new(model.Image)
	database.MysqlClientInstance.GormDBClient.First(&image, 54)
	fmt.Println(image.Url)
}
