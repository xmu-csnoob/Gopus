package dao

import (
	"Opus/src/main/mvc/mapper"
	"Opus/src/main/mvc/model"
)

type ImageDao struct {
	imageMapper mapper.ImageMapper
}

func (imageDao ImageDao) FindImageInDB(id int) model.Image {
	return imageDao.imageMapper.FindById(id)
}
