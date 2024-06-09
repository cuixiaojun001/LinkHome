package utils

import "testing"

func TestParseJWTToken(t *testing.T) {
	type args struct {
		tokenString string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "test case 1",
			args: args{
				tokenString: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VybmFtZSI6ImN1aXhpYW9qdW4iLCJyb2xlIjoibGFuZGxvcmQiLCJyZWZyZXNoIjpmYWxzZSwiZXhwIjoxNzE2MjgyOTY2fQ.jYTwaYuYew3mlOILGml0qy5uHKGXa4JuAhYuULaDCsg",
			},
			want:    2,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseJWTToken(tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseJWTToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseJWTToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}
