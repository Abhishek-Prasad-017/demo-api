package service

import (

	"github.com/magiconair/properties/assert"
	"log"
	"testing"
)

func TestDatabase(t *testing.T) {
	var results Softwares
	var res Softwares
	res = GetSoftwareDetails("", nil)
	log.Println(res)
	assert.Equal(t, res,results)
}
