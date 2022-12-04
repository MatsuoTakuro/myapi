package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/MatsuoTakuro/myapi-go-intermediate/apperrors"
	"github.com/MatsuoTakuro/myapi-go-intermediate/controllers/services"
	"github.com/MatsuoTakuro/myapi-go-intermediate/models"
	"github.com/gorilla/mux"
)

type ArticleController struct {
	service services.ArticleServicer
}

func NewArticleController(s services.ArticleServicer) *ArticleController {
	return &ArticleController{service: s}
}

// GET /hello のハンドラ
func (c *ArticleController) HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, `{"message": "Hello, world!"}`)
}

// POST /article のハンドラ
func (c *ArticleController) PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(w, req, err)
		return
	}

	article, err := c.service.PostArticleService(reqArticle)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	if err = json.NewEncoder(w).Encode(article); err != nil {
		err = apperrors.ResBodyEncodeFailed.Wrap(err, "fail to encode response body")
		apperrors.ErrorHandler(w, req, err)
		return
	}
}

// GET /article/list のハンドラ
func (c *ArticleController) ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	// クエリパラメータpageを取得
	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			err = apperrors.BadParam.Wrap(err, "queryparam must be number")
			apperrors.ErrorHandler(w, req, err)
			return
		}
	} else {
		page = 1
	}

	articleList, err := c.service.GetArticleListService(page)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	if err = json.NewEncoder(w).Encode(articleList); err != nil {
		err = apperrors.ResBodyEncodeFailed.Wrap(err, "fail to encode response body")
		apperrors.ErrorHandler(w, req, err)
		return
	}
}

// GET /article/{id} のハンドラ
func (c *ArticleController) ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		err = apperrors.BadParam.Wrap(err, "pathparam must be number")
		apperrors.ErrorHandler(w, req, err)
		return
	}

	article, err := c.service.GetArticleService(articleID)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	if err = json.NewEncoder(w).Encode(article); err != nil {
		err = apperrors.ResBodyEncodeFailed.Wrap(err, "fail to encode response body")
		apperrors.ErrorHandler(w, req, err)
		return
	}
}

// POST /article/nice のハンドラ
func (c *ArticleController) PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(w, req, err)
		return
	}

	article, err := c.service.PostNiceService(reqArticle)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	if err = json.NewEncoder(w).Encode(article); err != nil {
		err = apperrors.ResBodyEncodeFailed.Wrap(err, "fail to encode response body")
		apperrors.ErrorHandler(w, req, err)
		return
	}
}
