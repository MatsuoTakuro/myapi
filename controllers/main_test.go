package controllers_test

import (
	"testing"

	"github.com/MatsuoTakuro/myapi-go-intermediate/controllers"
	"github.com/MatsuoTakuro/myapi-go-intermediate/controllers/testdata"
	_ "github.com/go-sql-driver/mysql"
)

var aCon *controllers.ArticleController

func TestMain(m *testing.M) {
	ser := testdata.NewServiceMock()
	aCon = controllers.NewArticleController(ser)

	m.Run()
}
