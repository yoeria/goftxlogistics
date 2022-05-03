package util

import "testing"

func Test_FindFile(t *testing.T) {
	type args struct {
		root     string
		filename string
	}
	tests := []struct {
		name      string
		args      args
		wantSPath string
		wantErr   bool
	}{
		{wantSPath: string("../creds.env"), args: args{root: "../", filename: "creds.env"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSPath, err := FindFile(tt.args.root, tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSPath != tt.wantSPath {
				t.Errorf("FindFile() = %v, want %v", gotSPath, tt.wantSPath)
			}
		})
	}
}
