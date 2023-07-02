package service_module_component

import (
	"git.legchelife.ru/gitlab-instance-7d441567/catalog_m/ent"
)

type Component struct {
	db *ent.Client
}

func New(db *ent.Client) *Component {
	return &Component{db}
}
