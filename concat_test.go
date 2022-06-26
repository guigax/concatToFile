package main

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"testing"
)

func Test_getAmountOfFiles(t *testing.T) {
	type args struct {
		lineCount int
		splitAt   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1000 lines", args{1000, 500}, 2},
		{"753567 lines, split at 5000", args{753567, 5000}, 151},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAmountOfFiles(tt.args.lineCount, tt.args.splitAt); got != tt.want {
				t.Errorf("getAmountOfFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_getAmountOfFiles(b *testing.B) {
	var tests = []struct {
		name  string
		lines int
		split int
	}{
		{"1000 lines", 1000, 500},
		{"753567 lines, split at 5000", 753567, 5000},
	}
	for _, tt := range tests {
		b.Run(fmt.Sprintf(tt.name, tt.lines, tt.split), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				getAmountOfFiles(tt.lines, tt.split)
			}
		})
	}
}

func Test_readFile(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString("line 753550")

	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"1 line wrote", args{&buffer}, []string{"line 753550"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readFile(tt.args.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("readFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_readFile(b *testing.B) {
	var buffer bytes.Buffer
	for i := 1; i <= 2500000; i++ {
		buffer.WriteString("line " + strconv.Itoa(i) + "\n")
	}

	for i := 0; i < b.N; i++ {
		readFile(&buffer)
	}
}
