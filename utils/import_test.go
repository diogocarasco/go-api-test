package utils

import (
	"reflect"
	"testing"

	"github.com/diogocarasco/go-api-test/models"
)

func TestImportFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    []models.Feiras
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ImportFile(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ImportFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ImportFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
