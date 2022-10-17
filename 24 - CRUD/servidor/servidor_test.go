package servidor

import "testing"

func Test_teste(t *testing.T) {
	type args struct {
		parm string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Teste Rodrigo",
			args: args{
				parm: "Leandro",
			}, 
			want: "sim",
		},
		{
			name: "Teste Rodrigo",
			args: args{
				parm: "rodrigo",
			}, 
			want: "sim",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := teste(tt.args.parm); got != tt.want {
				t.Errorf("teste() = %v, want %v", got, tt.want)
			}
		})
	}
}
