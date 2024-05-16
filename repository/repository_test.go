package repository

import (
	"testing"
	"time"
)

func TestRepository(t *testing.T) {
	timeNow := time.Now()
	t.Log(timeNow)
	t.Log(timeNow.Unix())
}
