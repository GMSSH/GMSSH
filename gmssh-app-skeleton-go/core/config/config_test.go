package config

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestConfigFileAdapter(t *testing.T) {
	ctx := context.TODO()

	fileAdapter, err := NewAdapterFile()
	if err != nil {
		t.Error("Failed to load config file")
	}
	if !fileAdapter.Available(ctx) {
		t.Error("Failed to load config file")
	}

	data, err := fileAdapter.Data(ctx)
	if err != nil {
		t.Error("Config file parsing error")
	}
	t.Logf("data >>> %#v", data)
}

// Test data structure:
//
//	{
//	    "logger":{
//	        "level":"111",
//	        "stdout":true,
//	        "nihao":{
//	            "aaa":"000000000000"
//	        }
//	    },
//	    "zack":"22222"
//	}
func TestConfig(t *testing.T) {
	ctx := context.TODO()

	cfg, err := New()
	if err != nil {
		t.Error("Failed to load config file")
	}

	data := cfg.MustData(ctx)
	fmt.Println(">>>>> ", data)

	section := "logger.level"
	val, err := cfg.Cfg().GetValue(section).String()
	if err != nil {
		t.Error("Config file parsing error")
	}
	if !reflect.DeepEqual(val, "111") {
		fmt.Println("[+]>>>>> ", val)
		t.Error("Config file parsing error")
	}

	section1 := "logger.nihao.aaa"
	val1, err := cfg.Cfg().GetValue(section1).String()
	if err != nil {
		t.Error("Config file parsing error")
	}
	if !reflect.DeepEqual(val1, "000000000000") {
		fmt.Println("[+]>>>>> ", val1)
		t.Error("Config file parsing error")
	}

	section2 := "zack"
	val2, err := cfg.Cfg().GetValue(section2).String()
	if err != nil {
		t.Error("Config file parsing error")
	}
	if !reflect.DeepEqual(val2, "22222") {
		fmt.Println("[+]>>>>> ", val2)
		t.Error("Config file parsing error")
	}

	section3 := "logger.stdout"
	val3, err := cfg.Cfg().GetValue(section3).Bool()
	if err != nil {
		t.Error("Config file parsing error")
	}
	if !reflect.DeepEqual(val3, true) {
		fmt.Println("[+]>>>>> ", val3)
		t.Error("Config file parsing error")
	}
}
