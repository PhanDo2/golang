package commentlikebiz

import (
	"context"
	"finnal-exam/component/asyncjob"
	"finnal-exam/modules/commentlike/commentlikemodel"
)

type UserLikeComment interface {
	Create(ctx context.Context, data *commentlikemodel.CommentLike) error
}

type IncreaseLikeCount interface {
	IncreaseLikeCount(ctx context.Context, id int) error
}

type userLikeCommentBiz struct {
	store    UserLikeComment
	incStore IncreaseLikeCount
}

func NewUserLikeCommentBiz(store UserLikeComment, incStore IncreaseLikeCount) *userLikeCommentBiz {
	return &userLikeCommentBiz{store: store, incStore: incStore}
}

func (biz *userLikeCommentBiz) LikeComment(
	ctx context.Context,
	data *commentlikemodel.CommentLike,
) error {
	err := biz.store.Create(ctx, data)

	if err != nil {
		return commentlikemodel.ErrCannotLikeComment(err)
	}

	// side effect
	job := asyncjob.NewJob(func(ctx context.Context) error {
		return biz.incStore.IncreaseLikeCount(ctx, data.Id)
	})

	_ = asyncjob.NewGroup(true, job).Run(ctx)

	return nil
}
