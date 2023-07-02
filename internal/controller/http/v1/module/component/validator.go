package rt_module_component

import (
	"errors"
	"mime/multipart"
)

func (r *Router) ImageValidator(fileData *multipart.FileHeader) error {
	//проверяем размер (не более 2 мб)
	if fileData.Size > 2*1024*1024 {
		err := errors.New("размер файла не должен превышать 2 Мб")
		return err
	}

	//проверяем формат, должен быть jpg, png, gif
	switch fileData.Header["Content-Type"][0] {
	case "image/jpeg":
		return nil
	case "image/png":
		return nil
	case "image/gif":
		return nil
	default:
		err := errors.New("неверный формат файла. Допустимые форматы: jpg, png, gif")
		return err
	}
}
