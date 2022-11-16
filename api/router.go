package api

import (
	"database/sql"
	"net/http"

	"github.com/MatsuoTakuro/myapi-go-intermediate/controllers"
	"github.com/MatsuoTakuro/myapi-go-intermediate/services"
	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) *mux.Router {
	// Services.MyAppService implements all serivce methods
	// IOW, it implements both ArticleServicer and CommentServicer interface
	// So, interfaces sides are separated, but the implementation side is not
	ser := services.NewMyAppService(db)
	aCon := controllers.NewArticleController(ser)
	cCon := controllers.NewCommentController(ser)

	r := mux.NewRouter()

	r.HandleFunc("/hello", aCon.HelloHandler).Methods(http.MethodGet)

	r.HandleFunc("/article", aCon.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", aCon.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", aCon.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", aCon.PostNiceHandler).Methods(http.MethodPost)

	r.HandleFunc("/comment", cCon.PostCommentHandler).Methods(http.MethodPost)

	return r
}
