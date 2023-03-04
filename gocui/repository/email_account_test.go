package repository

import (
	"learngocui/model"
	"learngocui/tui"
	"reflect"
	"testing"
)

func TestSeedAccountRepository_FindById(t *testing.T) {
	seed := tui.SeedData()
	tests := []struct {
		name   string
		fields []model.EmailAccount
		args   int
		want   *model.EmailAccount
	}{
		{
			name:   "return first element",
			fields: seed,
			args:   0,
			want:   &seed[0],
		},
		{
			name:   "return nil on out of bounds",
			fields: seed,
			args:   len(seed),
			want:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			self := &SeedAccountRepository{
				store: tt.fields,
			}
			if got := self.FindById(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindById() = %v, want %v", got, tt.want)
			}
		})
	}
}
