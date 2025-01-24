package shrink

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func Test_generateShortcode(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "id=0",
			args: args{id: 0},
			want: "",
		}, {
			name: "id != 0",
			args: args{id: 123456},
			want: "2n9c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, generateShortcode(tt.args.id))
		})
	}
}

func TestShrinkService_GetLongUrlByCode(t *testing.T) {
	type args struct {
		ctx       context.Context
		shortCode string
	}
	tests := []struct {
		name                  string
		args                  args
		getSessionManagerMock func(ctrl *gomock.Controller) *MockLinkRepository
		want                  string
		wantErr               bool
	}{
		{
			name: "success",
			args: args{
				ctx:       context.Background(),
				shortCode: "qwerty",
			},
			getSessionManagerMock: func(ctrl *gomock.Controller) *MockLinkRepository {
				mock := NewMockLinkRepository(ctrl)
				mock.EXPECT().
					GetLongUrlByCode(context.Background(), "qwerty").
					Return("https://google.com", nil).
					Times(1)

				return mock
			},
			want: "https://google.com",
		}, {
			name: "failed",
			args: args{
				ctx:       context.Background(),
				shortCode: "qwerty",
			},
			getSessionManagerMock: func(ctrl *gomock.Controller) *MockLinkRepository {
				mock := NewMockLinkRepository(ctrl)
				mock.EXPECT().
					GetLongUrlByCode(context.Background(), "qwerty").
					Return("", assert.AnError).
					Times(1)

				return mock
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			s := &ShrinkService{}
			if tt.getSessionManagerMock != nil {
				s.linkRepository = tt.getSessionManagerMock(ctrl)
			}

			got, err := s.GetLongUrlByCode(tt.args.ctx, tt.args.shortCode)
			if tt.wantErr {
				assert.Error(t, err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
