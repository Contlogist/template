package uc_module_component

import (
	"git.legchelife.ru/gitlab-instance-7d441567/catalog_m/pkg/models/context"
)

func (uc *Repo) GetTest(ctx context.Base) (bool, error) {
	re, err := uc.repo.Module.Component.GetTest(ctx)
	if err != nil {
		return re, err
	} else {
		return re, nil
	}
}
