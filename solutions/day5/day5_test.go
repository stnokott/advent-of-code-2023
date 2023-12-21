// Package main runs the input for Day 5
package main

import (
	"testing"
)

func TestAlmanachLowestLocation(t *testing.T) {
	tests := []struct {
		name string
		a    *Almanach
		want int
	}{
		{
			"simple",
			&Almanach{
				Seeds: []int{4, 18},
				Maps: []Map{
					{
						Name: "soil-to-fertilizer",
						Ranges: []Range{
							{SrcStart: 2, SrcEnd: 10, DstOffset: -1},
							{SrcStart: 12, SrcEnd: 18, DstOffset: 90},
						},
					},
					{
						Name: "fertilizer-to-location",
						Ranges: []Range{
							{SrcStart: 0, SrcEnd: 10, DstOffset: 999},
							{SrcStart: 100, SrcEnd: 107, DstOffset: -50},
							{SrcStart: 108, SrcEnd: 109, DstOffset: 70},
						},
					},
				},
			},
			178,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.LowestLocation(); got != tt.want {
				t.Errorf("Almanach.LowestLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapGo(t *testing.T) {
	type args struct {
		src int
	}
	tests := []struct {
		name string
		m    Map
		args args
		want int
	}{
		{"one range", Map{Name: "simple", Ranges: []Range{{SrcStart: 5, SrcEnd: 10, DstOffset: 10}}}, args{6}, 16},
		{
			"multiple ranges match",
			Map{
				Name: "multiple-ranges-match",
				Ranges: []Range{
					{SrcStart: 87, SrcEnd: 100, DstOffset: -10},
					{SrcStart: 53, SrcEnd: 67, DstOffset: 50},
				},
			},
			args{60},
			110,
		},
		{
			"multiple ranges no match",
			Map{
				Name: "multiple-ranges-no-match",
				Ranges: []Range{
					{SrcStart: 87, SrcEnd: 100, DstOffset: -10},
					{SrcStart: 53, SrcEnd: 67, DstOffset: 50},
				},
			},
			args{75},
			75,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Go(tt.args.src); got != tt.want {
				t.Errorf("Map.Go() = %v, want %v", got, tt.want)
			}
		})
	}
}
