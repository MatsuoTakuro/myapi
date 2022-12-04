package services

import (
	"database/sql"
	"errors"

	"github.com/MatsuoTakuro/myapi-go-intermediate/apperrors"
	"github.com/MatsuoTakuro/myapi-go-intermediate/models"
	"github.com/MatsuoTakuro/myapi-go-intermediate/repositories"
)

// PostArticleHandlerで使うことを想定したサービス
// 引数の情報をもとに新しい記事を作り、結果を返却
func (s *MyAppService) PostArticleService(article models.Article) (models.Article, error) {
	newArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Article{}, err
	}
	return newArticle, nil
}

// ArticleListHandlerで使うことを想定したサービス
// 指定pageの記事一覧を返却
func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {
	articleList, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return nil, err
	}

	if len(articleList) == 0 {
		err := apperrors.NAData.Wrap(ErrNoData, "no data")
		return nil, err
	}

	return articleList, nil
}

// ArticleDetailHandlerで使うことを想定したサービス
// 指定IDの記事情報を返却
func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {
	type articleResult struct {
		article models.Article
		err     error
	}
	ac := make(chan articleResult)
	defer close(ac)

	go func(ch chan<- articleResult, db *sql.DB, articleID int) {
		a, err := repositories.SelectArticleDetail(db, articleID)
		ch <- articleResult{
			article: a,
			err:     err,
		}
	}(ac, s.db, articleID)

	type commentListResult struct {
		commentList *[]models.Comment
		err         error
	}
	cc := make(chan commentListResult)
	defer close(cc)

	go func(ch chan<- commentListResult, db *sql.DB, articleID int) {
		cl, err := repositories.SelectCommentList(db, articleID)
		ch <- commentListResult{
			commentList: &cl,
			err:         err,
		}
	}(cc, s.db, articleID)

	var article models.Article
	var commentList []models.Comment
	var articleGetErr, commentGetErr error

	for i := 0; i < 2; i++ {
		select {
		case a := <-ac:
			article, articleGetErr = a.article, a.err
		case c := <-cc:
			commentList, commentGetErr = *c.commentList, c.err
		}
	}

	if articleGetErr != nil {
		if errors.Is(articleGetErr, sql.ErrNoRows) {
			err := apperrors.NAData.Wrap(articleGetErr, "no data")
			return models.Article{}, err
		}
		err := apperrors.GetDataFailed.Wrap(articleGetErr, "fail to get data")
		return models.Article{}, err
	}

	if commentGetErr != nil {
		err := apperrors.GetDataFailed.Wrap(commentGetErr, "fail to get data")
		return models.Article{}, err
	}

	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

// PostNiceHandlerで使うことを想定したサービス
// 指定IDの記事のいいね数を+1して、結果を返却
func (s *MyAppService) PostNiceService(article models.Article) (models.Article, error) {
	err := repositories.UpdateNiceNum(s.db, article.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NoTargetData.Wrap(err, "does not exist target article")
			return models.Article{}, err
		}
		err = apperrors.InsertDataFailed.Wrap(err, "fail to update nice count")
		return models.Article{}, err
	}

	// TODO: I think this response is different from what the user wants.
	// this is just article in the request which 1 nice is added to the nice of (if it exists in the request).
	return models.Article{
		ID:        article.ID,
		Title:     article.Title,
		Contents:  article.Contents,
		UserName:  article.UserName,
		NiceNum:   article.NiceNum + 1,
		CreatedAt: article.CreatedAt,
	}, nil
}
