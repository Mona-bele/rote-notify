package postgres

import (
	"github.com/Mona-bele/rote-notify/infra/database/drive"
	util "github.com/Mona-bele/rote-notify/utils"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type IPostgres interface {
	Connect() []string
}

type Postgres struct {
	drive.IDrive
	Env *util.Env
}

func NewPostgres(env *util.Env) *Postgres {
	return &Postgres{
		Env: env,
	}
}

func (p *Postgres) Connect() string {
	return p.Env.DATABASE_URL_PRIMARY
}
