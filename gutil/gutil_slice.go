package gutil

import "reflect"

// SliceCopy 进行切片浅拷贝。
// []interface{}.
func SliceCopy(data []interface{}) []interface{} {
	newData := make([]interface{}, len(data))
	copy(newData, data)
	return newData
}

// SliceDelete 删除切片中指定索引对应的元素。
// 如果索引无效则不做任何操作。
func SliceDelete(data []interface{}, index int) (newSlice []interface{}) {
	if index < 0 || index >= len(data) {
		return data
	}
	// Determine array boundaries when deleting to improve deletion efficiency.
	if index == 0 {
		return data[1:]
	} else if index == len(data)-1 {
		return data[:index]
	}
	// If it is a non-boundary delete,
	// it will involve the creation of an array,
	// then the deletion is less efficient.
	return append(data[:index], data[index+1:]...)
}

// IndexOf 获取第一个参数在第二个参数中的位置
func IndexOf(params ...interface{}) int {
	v := reflect.ValueOf(params[0])
	arr := reflect.ValueOf(params[1])

	var t = reflect.TypeOf(params[1]).Kind()

	if t != reflect.Slice && t != reflect.Array {
		panic("Type Error! Second argument must be an array or a slice.")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == v.Interface() {
			return i
		}
	}
	return -1
}
