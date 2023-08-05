package envs

import (
	"sync"
)

type EnvStore struct {
	mutex sync.Mutex
	env   Env
}

var EnvStoreInstance *EnvStore

func (e *EnvStore) SetEnv(env Env) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	e.env = env
}

func (e *EnvStore) GetEnv() Env {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	return e.env
}

func IntiEnvStore(env Env) {
	EnvStoreInstance = &EnvStore{env: env}
}
