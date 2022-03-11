package commentlikestorage

import (
	"context"
	"finnal-exam/common"
	"finnal-exam/modules/commentlike/commentlikemodel"
)

func (s *sqlStore) Create(ctx context.Context, data *commentlikemodel.CommentLike) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
