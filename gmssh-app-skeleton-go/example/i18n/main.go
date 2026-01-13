package main

import (
	"path/filepath"

	"github.com/DemonZack/simplejrpc-go/core/gi18n"
)

func main() {
	fPath, _ := filepath.Abs(filepath.Dir(""))
	path := filepath.Join(fPath, "testdata")
	gi18n.Instance().SetPath(path)
	gi18n.Instance().SetLanguage(gi18n.English.String())

	key := "Welcome"
	val := gi18n.Instance().T(key)
	println(val)

	gi18n.Instance().SetLanguage(gi18n.SimplifiedChinese.String())
	val = gi18n.Instance().T(key)
	println(val)
}
