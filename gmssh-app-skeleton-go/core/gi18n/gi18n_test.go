package gi18n

import (
	"path/filepath"
	"reflect"
	"testing"

	"github.com/DemonZack/simplejrpc-go/os/gpath"
)

func TestI18nFileScan(t *testing.T) {
	gpath.GmCfgPath = filepath.Join(filepath.Dir(""), "testdata")
	i18nAdapter := NewI18nFileAdapter()
	t.Log("[*] search i18n path", i18nAdapter.searchPaths)
}

func TestI18nFileParser(t *testing.T) {
	gpath.GmCfgPath = filepath.Join(filepath.Dir(""), "testdata")
	i18nAdapter := NewI18nFileAdapter()
	localePath := i18nAdapter.localesPath[English]
	parser := CreateI18nParser(localePath.fileType, localePath.path)

	content, err := parser.GetContent()
	if err != nil {
		t.Error(err)
	}

	key := "Welcome"
	t.Log("[*] content >>> ", content, content.GetString(key))
}

func TestI18nTSetPath(t *testing.T) {
	fPath, _ := filepath.Abs(filepath.Dir(""))
	path := filepath.Join(fPath, "testdata")
	Instance().SetPath(path)
}

func TestI18nSetLanguage(t *testing.T) {
	fPath, _ := filepath.Abs(filepath.Dir(""))
	path := filepath.Join(fPath, "testdata")
	Instance().SetPath(path)
	Instance().SetLanguage(English.String())
}

func TestI18nTranslate(t *testing.T) {
	fPath, _ := filepath.Abs(filepath.Dir(""))
	path := filepath.Join(fPath, "testdata")
	Instance().SetPath(path)
	Instance().SetLanguage(English.String())

	key := "Welcome"
	val := Instance().T(key)
	t.Log(reflect.DeepEqual("Welcome to our application", val))

	Instance().SetLanguage(SimplifiedChinese.String())
	val = Instance().T(key)
	t.Log(reflect.DeepEqual("欢迎来到应用首页", val))

	// Test not exists key
	key = "Welcome1"
	val = Instance().T(key)
	t.Log(reflect.DeepEqual("Welcome1", val))
}

func TestI18nTranslateFormat(t *testing.T) {
	fPath, _ := filepath.Abs(filepath.Dir(""))
	path := filepath.Join(fPath, "testdata")
	Instance().SetPath(path)
	Instance().SetLanguage(English.String())

	key := "GMSSH"
	val := Instance().TranslateFormat(key, "World")
	t.Log(reflect.DeepEqual("Hello World", val))

	// Test not exists key
	key = "GMSSH1"
	val = Instance().TranslateFormat(key, "World")
	t.Log(reflect.DeepEqual("Hello World", val))
}
