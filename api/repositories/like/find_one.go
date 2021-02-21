package like

import (
	"errors"

	"gamersnote.com/v1/dtos"
	"gorm.io/gorm"
)

// PutOne いいねを追加します。
func (r repository) FindOne(like *dtos.Like) (bool, error) {
	// いいね取得
	l := &dtos.Like{}
	result := r.db.Where(like).First(l)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}
