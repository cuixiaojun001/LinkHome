package common

import (
	"context"
	"testing"
)

func TestGenerateContractContent(t *testing.T) {
	type args struct {
		in0        context.Context
		orderID    int
		templateID int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				in0:        context.Background(),
				orderID:    54,
				templateID: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateContractContent(tt.args.in0, tt.args.orderID, tt.args.templateID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateContractContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GenerateContractContent() got = %v, want %v", got, tt.want)
			}
		})
	}
}
