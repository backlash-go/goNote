package singleton

import (
	"reflect"
	"testing"
)

func TestGetInstance(t *testing.T) {
	tests := []struct {
		name string
		want *singleton
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetInstance(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInstance() = %v, want %v", got, tt.want)
			}
		})
	}


}