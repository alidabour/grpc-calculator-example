package usecase

import "testing"

func Test_Eval(t *testing.T) {
	evalUC := eval{}
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		e       eval
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "correct input",
			e:    evalUC,
			args: args{
				in: "100/10*5+10-5",
			},
			want:    55,
			wantErr: false,
		},
		{
			name: "wrong input",
			e:    evalUC,
			args: args{
				in: "1_1",
			},
			// want:    2,
			wantErr: true,
		},
		{
			name: "empty",
			e:    evalUC,
			args: args{
				in: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := eval{}
			got, err := e.Eval(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("eval.Eval() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("eval.Eval() = %v, want %v", got, tt.want)
			}
		})
	}
}
