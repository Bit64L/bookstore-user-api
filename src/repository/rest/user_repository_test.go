package rest

import (
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserInvalidJsonResponse(t *testing.T) {
	rest.FlushMockups()
	addMockErr := rest.AddMockups(&rest.Mock{
		URL:          "https://api.bookstore.com:5001/internal/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"test@test.com","password":"111111"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": "1", "first_name": "wx", "last_name": "l"}`,
	})
	if addMockErr != nil{
		panic(addMockErr)
	}


	repository := &userRepository{}
	user, err := repository.LoginUser("test@test.com", "111111")
	assert.Nil(t, user)
	assert.NotNil(t, err)

}
