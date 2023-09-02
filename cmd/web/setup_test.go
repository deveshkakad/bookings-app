package main

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/deveshkakad/bookings-app/internal/handlers"
	"github.com/deveshkakad/bookings-app/internal/models"
	"github.com/deveshkakad/bookings-app/internal/render"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

type myHandler struct{}

func (mh *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

//var app config.AppConfig
//var session *scs.SessionManager

