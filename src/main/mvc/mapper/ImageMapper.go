package mapper

import (
	"Opus/src/main/mvc/model"
	"Opus/src/main/mvc/util"
	"fmt"
)

type ImageMapper struct {
}

func (imageMapper *ImageMapper) FindById(id int) {
	image := new(model.Image)
	util.GormDBClient.First(&image, 54)
	fmt.Println(image.Url)
}
