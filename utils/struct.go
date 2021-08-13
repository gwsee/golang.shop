package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)


//根据 struct 将特点的tag 解析成一个map数据返回
func Resolve(ptr interface{},m map[string]interface{},f string,filter string) (err error)  {
	if reflect.TypeOf(ptr).Kind() != reflect.Ptr{
		return errors.New("i is not pointer")
	}
	val:=reflect.ValueOf(ptr)
	elem:=val.Elem()
	ty:=elem.Type()
	for k:=0;k<elem.NumField();k++{
		v:=elem.Field(k)
		t:=ty.Field(k)
		if t.Anonymous {
			structResolve(v,t,m,f)
		}else{
			key:=t.Name
			if f!=""{
				r,ok:=t.Tag.Lookup(f)
				if !ok{
					continue
				}
				key = r
			}
			m[key] = v.Interface() //一般是第一层优先原则
		}
	}
	if filter != ""{
		fArr := strings.Split(filter,",")
		for _,v:=range fArr{
			delete(m, v)
		}
	}
	return err
}
func isValid(t reflect.Type,v reflect.Value)( b bool) {
	n := t.Name()
	if n == "string" || n=="byte" {
		b=fmt.Sprintf("%v",v)!=""
	}else if n =="int" ||n=="float64"{
		b=fmt.Sprintf("%v",v)!="0"
	}
	return
}
func structResolve(val reflect.Value,field reflect.StructField,m map[string]interface{},f string)  {
	ty:=field.Type
	num:=val.NumField()
	for i:=0;i<num;i++{
		v:=val.Field(i)
		t:=ty.Field(i)
		if t.Anonymous{
			structResolve(v,t,m,f)
		}else{
			key:=t.Name
			if f!=""{
				r,ok:=t.Tag.Lookup(f)
				if !ok{
					continue
				}
				key = r
			}
			_,ok:=m[key]
			if !ok{
				m[key] = v.Interface()
			}
		}
	}
}

