package mapper

import (
	"Opus/src/main/application"
	"Opus/src/main/mvc/model"
	"fmt"
)

type ImageMapper struct {
}

func (imageMapper *ImageMapper) FindById(id int) {
	image := new(model.Image)
	application.MysqlClientInstance.GormDBClient.First(&image, 54)
	fmt.Println(image.Url)
}
