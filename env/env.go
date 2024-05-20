package env

import (
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

const EnvKey = "Env"
const ProdEnv = "prod"
const DevEnv = "dev"

var Environment *Env

type Env struct {
	Name     string // env Name, like `prod`, `dev` and etc.,
	Debugger bool   // whether to use debugger
}

func (env *Env) IsDevelopment() bool {
	return strings.EqualFold(DevEnv, env.Name)
}

func (env *Env) IsProduction() bool {
	return strings.EqualFold(ProdEnv, env.Name)
}

func (env *Env) GetEnvName() *string {
	return &env.Name
}
func Init() {
	envName := DevEnv
	if len(os.Getenv(EnvKey)) > 0 {
		envName = os.Getenv(EnvKey)
	}
	Environment = &Env{
		Name: envName,
		Debugger: func() bool {
			return envName != ProdEnv
		}(),
	}
	logrus.Infof("Config loaded successfully Env: %s", Environment.Name)

}
