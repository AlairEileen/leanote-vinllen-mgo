package tests

import (
	"github.com/leanote/leanote/app/db"
	"testing"
	//	. "github.com/leanote/leanote/app/lea"
	//	"github.com/leanote/leanote/app/service"
	//	"github.com/vinllen/mgo"
	//	"fmt"
)

func TestDBConnect(t *testing.T) {
	db.Init("mongodb://localhost:27017/leanote", "leanote")
}
