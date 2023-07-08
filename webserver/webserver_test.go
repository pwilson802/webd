package webserver_test

import (
	"testing"

	"github.com/pwilson802/webd/webserver"
)

func TestCreateSimpleWebServerConfig(t *testing.T) {
	config := webserver.ServerConfig{
		Port: 80,
	}
	want := 80
	got := config.Port
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}

}
