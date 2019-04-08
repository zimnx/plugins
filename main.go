package main

import (
	"github.com/zimnx/central/ext"
)

func Init(m ext.Map) {
	m.Add("hello", "world")
}