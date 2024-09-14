package config

import "testing"

func Test_Production(t *testing.T) {
	got := Production()
	if got.Name() != "production" {
		t.Errorf("Production() = %v, want %v", got.Name(), "production")
	}
}

func Test_Development(t *testing.T) {
	got := Development()
	if got.Name() != "development" {
		t.Errorf("Development() = %v, want %v", got.Name(), "development")
	}
}

func Test_path(t *testing.T) {
	type args struct {
		e Environment
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Production 向けのパスが生成される",
			args: args{e: Production()},
			want: "config/production.yaml",
		},
		{
			name: "Development 向けのパスが生成される",
			args: args{e: Development()},
			want: "config/development.yaml",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := path(tt.args.e); got != tt.want {
				t.Errorf("path() = %v, want %v", got, tt.want)
			}
		})
	}
}
