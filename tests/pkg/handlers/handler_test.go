package handlers_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"prw_server/app/pkg/api"
	"prw_server/app/pkg/handlers"
	"prw_server/tests/mocks"
	"testing"
)

func TestHandlerHello(t *testing.T) {
	tests := []struct {
		name     string
		joke     *api.JokeResponse
		err      error
		codeWant int
		bodyWant string
	}{
		{
			name:     "simple test",
			joke:     &api.JokeResponse{Joke: "test joke"},
			err:      nil,
			codeWant: 200,
			bodyWant: "test joke",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			apiMock := &mocks.Client{}
			apiMock.On("GetJoke").Return(test.joke, test.err)

			handler := handlers.NewHandler(apiMock)
			request, _ := http.NewRequest("GET", "/hello", nil)

			response := httptest.NewRecorder()

			handler.Hello(response, request)

			gotRaw, _ := ioutil.ReadAll(response.Body)

			if string(gotRaw) != test.bodyWant {
				t.Errorf("wrong response body %s want %s", gotRaw, test.bodyWant)
			}

			if status := response.Result().StatusCode; status != test.codeWant {
				t.Errorf("wrong response status %d want %d", status, test.codeWant)
			}
		})
	}
}
