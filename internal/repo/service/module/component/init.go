package service_module_component

import (
	"git.legchelife.ru/root/template/ent"
)

type Component struct {
	db *ent.Client
}

func New(db *ent.Client) *Component {
	return &Component{db}
}
