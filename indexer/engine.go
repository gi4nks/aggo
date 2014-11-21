package indexer

import (
	//"fmt"
	"github.com/gi4nks/aggo/indexes"
)

type Engine struct {

	Indexes map[string]Index

}

func NewEngine() *Engine {
	return &Engine{Indexes: make(map[string]Index)}
}

func Index(engine *Engine, file string) {

}
