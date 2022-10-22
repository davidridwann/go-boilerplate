package loader

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func TestYAMLLoader_Load(t *testing.T) {
	type Config struct {
		Out string `yaml:"out"`
		Err string `yaml:"err"`
	}

	type Example struct {
		Log Config `yaml:"log"`
	}

	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    Example
		wantErr bool
	}{
		{
			name: "example",
			args: args{
				reader: bytes.NewReader([]byte(`log:
  out: /dev/stdout
  err: /dev/stderr
`)),
			},
			want: Example{
				Log: Config{
					Out: "/dev/stdout",
					Err: "/dev/stderr",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			yamlLoader := YAMLLoader[Example]{}
			got, err := yamlLoader.Load(tt.args.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Load() got = %v, want %v", got, tt.want)
			}
		})
	}
}
