package user

import "gamersnote.com/v1/dtos"

// DeleteOne ユーザーを削除します。UserIDで検索します。
func (r repository) DeleteOne(userID string) error {
	q := &dtos.User{UserID: userID}
	result := r.db.Delete(q)
	return result.Error
}
