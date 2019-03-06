package grpcAPI

import (
	"context"
	"reflect"
	"server/adapters/grpcAPI/entity"
	"server/usecase"
	"testing"
)

func Test_server_Eval(t *testing.T) {
	uc := usecase.New()
	type fields struct {
		usecase usecase.Eval
	}
	type args struct {
		ctx context.Context
		t   *entity.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Result
		wantErr bool
	}{
		{
			name: "correct input",
			fields: fields{
				usecase: uc,
			},
			args: args{
				ctx: context.Background(),
				t:   &entity.Request{Value: "10*3/5"},
			},
			want:    &entity.Result{Value: "6"},
			wantErr: false,
		},
		{
			name: "wrong input",
			fields: fields{
				usecase: uc,
			},
			args: args{
				ctx: context.Background(),
				t:   &entity.Request{Value: "1_2"},
			},
			want:    &entity.Result{},
			wantErr: true,
		},
		{
			name: "empty input",
			fields: fields{
				usecase: uc,
			},
			args: args{
				ctx: context.Background(),
				t:   &entity.Request{Value: ""},
			},
			want:    &entity.Result{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := server{
				usecase: tt.fields.usecase,
			}
			got, err := s.Eval(tt.args.ctx, tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.Eval() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.Eval() = %v, want %v", got, tt.want)
			}
		})
	}
}
