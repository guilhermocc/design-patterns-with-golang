package creational

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_getConveyance(t *testing.T) {
	type args struct {
		conveyanceType string
	}
	tests := []struct {
		name             string
		args             args
		want             Conveyor
		wantErr          bool
		errMessageWanted string
	}{
		{
			name: "test get conveyance truck",
			args: args{
				conveyanceType: "truck",
			},
			want: &truck{
				conveyance: conveyance{
					name:  "Truck conveyance",
					speed: 4,
				},
			},
			wantErr: false,
		},
		{
			name: "test get conveyance ship",
			args: args{
				conveyanceType: "ship",
			},
			want: &ship{
				conveyance: conveyance{
					name:  "Ship conveyance",
					speed: 1,
				},
			},
			wantErr: false,
		},
		{
			name: "test get conveyance airplane",
			args: args{
				conveyanceType: "airplane",
			},
			want: &airplane{
				conveyance: conveyance{
					name:  "Airplane conveyance",
					speed: 1,
				},
			},
			wantErr: false,
		},
		{
			name: "test get conveyance unknown",
			args: args{
				conveyanceType: "unknown",
			},
			want:             nil,
			wantErr:          true,
			errMessageWanted: "wrong conveyance type passe",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetConveyance(tt.args.conveyanceType)
			if tt.wantErr {
				require.Error(t, err, tt.errMessageWanted)
			} else if err != nil {
				t.Errorf("getGun() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.Equal(t, tt.want, got)
		})
	}
}
