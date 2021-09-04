package service_http

import (
	"reflect"

	"github.com/eolinker/eosc"
)

const (
	driverName = "http"
)

//driver 实现github.com/eolinker/eosc.eosc.IProfessionDriver接口
type driver struct {
	profession string
	name       string
	driver     string
	label      string
	desc       string
	configType reflect.Type
	params     map[string]string
}

//ConfigType 返回service_http驱动配置的反射类型
func (d *driver) ConfigType() reflect.Type {
	return d.configType
}

//Create 创建service_http驱动的实例
func (d *driver) Create(id, name string, v interface{}, workers map[eosc.RequireId]interface{}) (eosc.IWorker, error) {

	w := &serviceWorker{
		id:     id,
		name:   name,
		driver: driverName,
	}
	err := w.Reset(v, workers)
	if err != nil {
		return nil, err
	}

	return w, nil
}