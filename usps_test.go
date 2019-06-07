package go_usps

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInitUSPS(t *testing.T) {
	username := "USPSUsername"
	password := "USPSPassword"
	appId := "USPSAppId"

	type args struct {
		username     string
		password     string
		appId        string
		isProduction bool
	}
	tests := []struct {
		name string
		args args
		want *USPS
	}{
		{name: "Success flow", args: args{username, password, appId, false}, want: InitUSPS(username, password, appId, false)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitUSPS(tt.args.username, tt.args.password, tt.args.appId, tt.args.isProduction); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitUSPS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleInitUSPS() {
	username := "USPSUsername"
	password := "USPSPassword"
	appId := "USPSAppId"

	usps := InitUSPS(username, password, appId, false)
	fmt.Println(usps.AppId)
}
