package repo

import (
	"git.legchelife.ru/root/template/internal/repo/company"
	"git.legchelife.ru/root/template/internal/repo/section"
	"git.legchelife.ru/root/template/internal/repo/user"
	"github.com/upper/db/v4"
)

type Repo struct {
	UserRepo    *user.Repo
	CompanyRepo *rp_company.Company
	SectionRepo *rp_section.Section
}

func New(db *db.Session) *Repo {
	return &Repo{
		user.New(db),
		rp_company.New(db),
		rp_section.New(db),
	}
}

func NewFake(
	UserRepo *user.Repo,
	CompanyRepo *rp_company.Company,
	SectionRepo *rp_section.Section,
) *Repo {
	return &Repo{
		UserRepo,
		CompanyRepo,
		SectionRepo,
	}
}
