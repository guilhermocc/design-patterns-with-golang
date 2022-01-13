package creational

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_getSportsFactory(t *testing.T) {
	type args struct {
		brand string
	}
	type want struct {
		factory SportsFactory
		shoes   Shoe
		shirt   Clothing
	}
	tests := []struct {
		name      string
		args      args
		want      want
		wantErr   bool
		errWanted string
	}{
		{
			name: "test get nikeFactory concrete factory",
			args: args{
				brand: "nike",
			},
			want: want{
				factory: &nikeFactory{},
				shoes: &nikeShoe{
					sneaker: sneaker{
						logo: "nikeFactory",
						size: 14,
					},
				},
				shirt: &nikeShirt{
					shirt: shirt{
						logo: "nikeFactory",
						size: 14,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "test get adidasFactory concrete factory",
			args: args{
				brand: "adidas",
			},
			want: want{
				factory: &adidasFactory{},
				shoes: &adidasShoe{
					sneaker: sneaker{
						logo: "adidasFactory",
						size: 14,
					},
				},
				shirt: &adidasShirt{
					shirt: shirt{
						logo: "adidasFactory",
						size: 14,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "test get unknown concrete factory",
			args: args{
				brand: "none",
			},
			want: want{
				factory: nil,
				shoes:   nil,
				shirt:   nil,
			},
			wantErr:   true,
			errWanted: "unknown brand none",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getSportsFactory(tt.args.brand)
			if tt.wantErr {
				require.Error(t, err, tt.errWanted)
				return
			} else if err != nil {
				t.Errorf("getSportsFactory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.Equal(t, tt.want.factory, got)
			require.Equal(t, tt.want.shoes, got.makeShoe())
			require.Equal(t, tt.want.shirt, got.makeShirt())
		})
	}
}
