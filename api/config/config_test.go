package config

import (
	"reflect"
	"testing"
)

func TestConfig(t *testing.T) {
	expectConfig := NewConfig()
	expectConfig.AppName = "XXXXX"

	config = NewConfig()

	if !reflect.DeepEqual(config, expectConfig) {
		t.Errorf("is not same config, actual=[%+v] want=[%+v]", config, expectConfig)
	}

}
