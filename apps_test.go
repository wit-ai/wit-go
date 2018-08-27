package witai

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetApps(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`[{"name": "app1"}, {"name":"app2"}]`))
	}))
	defer func() { testServer.Close() }()

	c := NewClient(unitTestToken)
	c.APIBase = testServer.URL
	apps, _ := c.GetApps(1, 0)

	if len(apps) != 2 {
		t.Fatalf("expected 2 apps, got: %v", apps)
	}
}
