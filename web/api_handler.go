package web

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"service_catalog/models"
	"strconv"

	rqp "github.com/timsolov/rest-query-parser"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type ServiceCatalogSvc struct {
	client *gorm.DB
}

func NewServiceCatalogSvc(client *gorm.DB) *ServiceCatalogSvc {
	return &ServiceCatalogSvc{client: client}
}

func (svc *ServiceCatalogSvc) FindById(ctx context.Context) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		id, err := strconv.Atoi(ectx.Param("id"))
		if err != nil {
			log.Error(err)
			return Return(ectx, http.StatusBadRequest, fmt.Errorf("Invalid Id: %v", err), nil)
		}
		var service models.Service
		tx := svc.client.Find(&service, id)
		if err != nil && errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			log.Error(err)
			return Return(ectx, http.StatusNotFound, err, nil)
		}
		return Return(ectx, http.StatusOK, nil, service)

	}
}

func (svc *ServiceCatalogSvc) FindVersions(ctx context.Context) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		id, err := strconv.Atoi(ectx.Param("id"))
		if err != nil {
			log.Error(err)
			return Return(ectx, http.StatusBadRequest, fmt.Errorf("Invalid Id:%v", err), nil)
		}
		var versions []models.Version
		tx := svc.client.Where("service_id=?", id).Find(&versions)
		if err != nil && errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			log.Error(err)
			return Return(ectx, http.StatusNotFound, err, nil)
		}
		return Return(ectx, http.StatusOK, nil, versions)

	}
}

func (svc *ServiceCatalogSvc) List(ctx context.Context) echo.HandlerFunc {
	return func(ectx echo.Context) error {
		q, err := rqp.NewParse(ectx.Request().URL.Query(), rqp.Validations{
			"limit:required": rqp.MinMax(4, 20),
			"sort":           rqp.In("name", "id"),
			"name":           nil,
			"id:int":         nil,
		})

		if err != nil {
			log.Error(err)
			return Return(ectx, http.StatusBadRequest, err, nil)
		}

		var services []struct {
			ID          uint
			Name        string
			Description string
			Count       int
		}
		if len(q.Args()) == 0 {
			tx := svc.client.Raw(`select id,name,description, COALESCE(v.count, 0) as count from services order by id s left join (select service_id,
 count(service_id) as count from versions  group by service_id) v on s.id=v.service_id order by id`).Scan(&services)
			err := tx.Error
			if err != nil {
				log.Error(err)
				return Return(ectx, http.StatusInternalServerError, err, nil)
			}
			return Return(ectx, http.StatusOK, nil, services)
		}

		err = svc.client.Raw(fmt.Sprintf(`select id,name,description, COALESCE(v.count, 0) as count from services join (select service_id,
 count(service_id) as count from versions group by service_id) v on id=v.service_id and %s %s %s`, q.Where(), q.ORDER(), q.LIMIT()), q.Args()...).Scan(&services).Error
		if err != nil {
			log.Error(err)
			return Return(ectx, http.StatusInternalServerError, err, nil)
		}
		if len(services) > 0 {
			return ReturnWithLinks(ectx, http.StatusOK, nil, services, map[string]string{
				"next":     fmt.Sprintf("id[gt]=%d", services[len(services)-1].ID),
				"previous": fmt.Sprintf("id[lt]=%d", services[0].ID),
			})
		}
		return Return(ectx, http.StatusNoContent, nil, nil)

	}
}

type Response struct {
	Error string            `json:"error,omitempty"`
	Data  interface{}       `json:"data,omitempty"`
	Links map[string]string `json:"links,omitempty"`
}

func ReturnWithLinks(ectx echo.Context, status int, err error, data interface{}, links map[string]string) error {
	response := Response{
		Data:  data,
		Links: links,
	}

	return ectx.JSON(status, response)
}

func Return(ectx echo.Context, status int, err error, data interface{}) error {
	response := Response{
		Data: data,
	}

	if err != nil {
		response.Error = err.Error()
	}

	return ectx.JSON(status, response)
}
