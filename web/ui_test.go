package web_test

import (
	"net/http"
	"testing"

	"github.com/steinfletcher/apitest"
	"github.com/steinfletcher/apitest-css-selector"
	"github.com/steinfletcher/templ-play/web"
)

func TestGetTodos(t *testing.T) {
	apiTest().
		Get("/todos").
		Expect(t).
		Status(http.StatusOK).
		Assert(
			selector.ContainsTextValue(
				"#post-1",
				"Buy milk",
			),
		).
		End()
}

func apiTest() *apitest.APITest {
	return apitest.New().
		Report(apitest.SequenceDiagram()).
		Handler(web.NewRouter())
}
