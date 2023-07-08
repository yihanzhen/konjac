package kana

import (
	"fmt"
	"testing"
)

func TestIsHiragana(t *testing.T) {
	cases := []struct {
		r    rune
		want bool
	}{
		{
			r:    'る',
			want: true,
		},
		{
			r:    'ぶ',
			want: true,
		},
		{
			r: 'a',
		},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v", tc.r), func(t *testing.T) {
			if got := IsHiragana(tc.r); got != tc.want {
				t.Errorf("IsHiragana(%v), got %v, want %v", tc.r, got, tc.want)
			}
		})
	}
}

func TestIsCol(t *testing.T) {
	cases := []struct {
		r    rune
		col  int
		want bool
	}{
		{
			r:    'る',
			col:  2,
			want: true,
		},
		{
			r:   'ぼ',
			col: 3,
		},
		{
			r:   'a',
			col: 1,
		},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v", tc.r), func(t *testing.T) {
			if got := IsCol(tc.r, tc.col); got != tc.want {
				t.Errorf("IsCol(%v), got %v, want %v", tc.r, got, tc.want)
			}
		})
	}
}

func TestToCol(t *testing.T) {
	cases := []struct {
		r       rune
		col     int
		want    rune
		wantErr bool
	}{
		{
			r:    'る',
			col:  0,
			want: 'ら',
		},
		{
			r:       'ぼ',
			col:     6,
			wantErr: true,
		},
		{
			r:       'ぼ',
			col:     -1,
			wantErr: true,
		},
		{
			r:       'ゆ',
			col:     1,
			wantErr: true,
		},
		{
			r:       'a',
			col:     1,
			wantErr: true,
		},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v", tc.r), func(t *testing.T) {
			got, err := ToCol(tc.r, tc.col)
			if (err != nil) != tc.wantErr {
				t.Errorf("Col(%v, %v), got error %v, want error: %v", tc.r, tc.col, err != nil, tc.wantErr)
			}
			if err == nil {
				if got != tc.want {
					t.Errorf("Col(%v, %v), got %c, want %c", tc.r, tc.col, got, tc.want)
				}
			}
		})
	}
}
