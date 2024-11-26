package main

import (
	"net/http"
	"testing"

	"im3shn/snippetbox/internal/assert"
)

func TestPing(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/ping")

	assert.Equal(t, code, http.StatusOK)
	assert.Equal(t, body, "OK")
}

func TestSnippetView(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody string
	}{
        {
            name: "Valid ID",
            urlPath: "/snippet/view/1",
            wantCode: http.StatusOK,
            wantBody: "Test 1",
        }, {
            name: "Non-existent ID",
            urlPath: "/snippet/view/2",
            wantCode: http.StatusNotFound,
        }, {
            name: "Negative ID",
            urlPath: "/snippet/view/-1",
            wantCode: http.StatusNotFound,
        }, {
            name: "String ID",
            urlPath: "/snippet/view/one",
            wantCode: http.StatusNotFound,
        }, {
            name: "Empty ID",
            urlPath: "/snippet/view/",
            wantCode: http.StatusNotFound,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            code, _, body := ts.get(t, tt.urlPath)

            assert.Equal(t, code, tt.wantCode)
            
            if tt.wantBody != "" {
                assert.StringContains(t, body, tt.wantBody)
            }
        })
    }
}
