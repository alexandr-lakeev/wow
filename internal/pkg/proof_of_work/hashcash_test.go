package proof_of_work

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alexandr-lakeev/wow.git/internal/pkg/proof_of_work/dto"
	"github.com/alexandr-lakeev/wow.git/internal/pkg/proof_of_work/mocks"
)

func TestHashcash_Solve(t *testing.T) {
	testChallenge := dto.NewChallenge("ver1", 2, 1, "rand", "test", 1)

	type fields struct {
		hasher     Hasher
		version    string
		complexity int
		prefix     string
		max        int
	}

	type args struct {
		challenge *dto.Challenge
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "successful solving of the challenge",
			fields: fields{
				hasher: func() Hasher {
					mock := mocks.Hasher{}

					mock.EXPECT().Hash([]byte("ver1:2:1:rand:test:1")).
						Return([]byte{255, 255, 255, 255})

					mock.EXPECT().Hash([]byte("ver1:2:1:rand:test:2")).
						Return([]byte{0, 255, 255, 255})

					return &mock
				}(),
				version:    "ver1",
				complexity: 2,
				prefix:     "0",
				max:        3,
			},
			args: args{
				challenge: testChallenge,
			},
			want: 2,
		},
		{
			name: "error no solution found",
			fields: fields{
				hasher: func() Hasher {
					mock := mocks.Hasher{}

					mock.EXPECT().Hash([]byte("ver1:2:1:rand:test:1")).
						Return([]byte{255, 255, 255, 255})

					return &mock
				}(),
				version:    "ver1",
				complexity: 2,
				prefix:     "0",
				max:        2,
			},
			args: args{
				challenge: func() *dto.Challenge {
					return testChallenge
				}(),
			},
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(
				tt.fields.hasher,
				tt.fields.version,
				tt.fields.complexity,
				tt.fields.prefix,
				tt.fields.max,
			)

			got, err := h.Solve(tt.args.challenge)

			if (err != nil) != tt.wantErr {
				t.Errorf("Solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, got, tt.want)
		})
	}
}
