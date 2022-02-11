/*
 * @Auther: BinyGo
 * @Description:
 * @Date: 2022-02-11 19:19:29
 * @LastEditTime: 2022-02-11 22:04:51
 */
package singleton

import (
	"reflect"
	"testing"
)

func TestGetInstance(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *singleton
	}{
		{
			name: "new singleton",
			args: args{"BinyGo"},
			want: &singleton{name: "BinyGo"},
		},
		{
			name: "new singleton2?",
			args: args{"BinyGo2"},
			want: &singleton{name: "BinyGo"},
		},
		{
			name: "new singleton3?",
			args: args{"BinyGo3"},
			want: &singleton{name: "BinyGo"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetInstance(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInstance() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestGetInstance2(t *testing.T) {
	tests := []struct {
		name string
		want *singleton
	}{
		{
			name: "new singleton",
			want: &singleton{name: "BinyGo"},
		},
		{
			name: "new singleton2",
			want: &singleton{name: "BinyGo"},
		},
		{
			name: "new singleton3",
			want: &singleton{name: "BinyGo"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetInstance2(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInstance2() = %v, want %v", got, tt.want)
			}
		})
	}
}
