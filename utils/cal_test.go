package utils

import (
	"fmt"
	"reflect"
	"testing"
)

type School struct {
	Name    string
	address string
	phone   string
}

func TestSayHello(t *testing.T) {

	sayHello("Tiger")
}

func TestGetHello(t *testing.T) {

	sayHello("Ken")
}

func TestAddCycle(t *testing.T) {

	//AddCycle(10)
}

type inter interface {
}

func TestReflectDemo(t *testing.T) {

	school := School{
		"吴川市第一中学",
		"吴川市解放路11号",
		"0759-5605555",
	}

	schoolType := reflect.TypeOf(school)
	//schoolType.Implements(reflect.Type())
	//schType := reflect.TypeOf(&school)
	//_ = schoolType.(type)
	fmt.Println(schoolType.Name())
	//fmt.Println(schoolType.Len())
	//fmt.Println(schType.Bits())
	fmt.Println(schoolType.Kind().String())
	fmt.Println(schoolType.PkgPath())
	structField, _ := schoolType.FieldByName("name")
	fmt.Println(structField)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()

	value := reflect.ValueOf(school)

	t2 := value.Type()
	fmt.Println(t2 == schoolType)
	i := value.Interface()
	switch i.(type) {
	case School:
		fmt.Println("")
	default:
		fmt.Println("")
	}
	s := i.(School)
	fmt.Println(s.phone)
}

const (
	a = 0
	b
	c
	d
	e = 2
	f
	g
	h
	i = iota
)

func TestHi(t *testing.T) {

	var age = 10

	value := reflect.ValueOf(&age)

	value.Elem().SetInt(99)

	fmt.Println(age)

	sch := School{
		"一中",
		"解放路",
		"0759",
	}

	schVal := reflect.ValueOf(sch)

	typeVal := reflect.TypeOf(sch)

	numField := typeVal.NumField()

	for i := 0; i < numField; i++ {
		fieldName := typeVal.Field(i).Name

		fieldByName := schVal.FieldByName(fieldName)
		str := fieldByName.Kind().String()
		if str == "string" {
			//fieldByName.SetString("1212")
			fieldByName.Elem().SetString("132154")
		}
	}
	fmt.Println(sch)
}

func TestModifyFieldVal(t *testing.T) {

	sch := &School{
		"一中",
		"解放路",
		"0759",
	}

	reflectValue := reflect.ValueOf(sch)
	fmt.Println(reflectValue)

	elems := reflectValue.Elem()

	// 反射方式设置值,字段必须是大写开头
	elems.FieldByName("Name").SetString("first")

	reflectType := reflect.TypeOf(sch)
	fmt.Println(reflectType)

	fmt.Println(*sch)
}
