// Copyright 2018 cloudy itcloudy@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package services

import (
	"github.com/itcloudy/base-framework/pkg/conf"
	"github.com/itcloudy/base-framework/pkg/interfaces/repositories"
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"
)

type SystemAPIService struct {
	DB *gorm.DB
	repositories.ISystemAPIRepository
}

func (service *SystemAPIService) ServiceGetSystemAPIByID(id int) (result models.SystemApiDetail, err error) {
	return service.FindSystemAPIByID(service.DB, id)
}
func (service *SystemAPIService) ServiceSystemAPIDelete(ids []int) (err error) {
	return service.DeleteSystemAPI(service.DB, ids)
}
func (service *SystemAPIService) ServiceSystemAPIUpdate(model models.SystemApiUpdate) (result models.SystemApiDetail, err error) {
	return service.UpdateSystemAPI(service.DB, model)
}
func (service *SystemAPIService) ServiceSystemAPICreate(model models.SystemApiCreate) (result models.SystemApiDetail, err error) {
	var validate *validator.Validate
	validate = validator.New()
	if err = validate.Struct(model); err != nil {
		return
	}
	return service.InsertSystemAPI(service.DB, model)
}
func (service *SystemAPIService) ServiceActiveSystemAPI(ids []int, active bool) (err error) {
	return service.ActiveSystemAPI(service.DB, ids, active)

}
func (service *SystemAPIService) ServiceGetAllSystemAPI(page, size int, order string, query string, queryArgs ...interface{}) (results []*models.SystemApiList, pagination conf.Pagination, err error) {
	pagination.Current = page
	pagination.Size = size
	results, pagination.Total, err = service.FindAllSystemAPI(service.DB, page, size, order, query, queryArgs)
	return
}
