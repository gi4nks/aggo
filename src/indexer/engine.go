package indexer

import (
	//"fmt"
	. "indexes"
)

type Engine struct {

	Indexes map[string]Index

}

func NewEngine() *Engine {
	return &Engine{Indexes: make(map[string]Index)}
}

func Index(engine *Engine, file string) {

}
