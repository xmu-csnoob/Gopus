package service

import (
	"Opus/src/main/mvc/dao"
	"Opus/src/main/mvc/model"
)

type ImageService struct {
	imageDao dao.ImageDao
}

func (imageService ImageService) GetImageById(id int) model.Image {
	return imageService.imageDao.FindImageInDB(id)
}
