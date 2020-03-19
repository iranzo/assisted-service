package subsystem

import (
	"github.com/filanov/bm-inventory/models"
	"github.com/go-openapi/strfmt"
)

func clearDB() {
	db.Delete(&models.Image{})
	db.Delete(&models.Node{})
	db.Delete(&models.Cluster{})
}

func strToUUID(s string) *strfmt.UUID {
	u := strfmt.UUID(s)
	return &u
}