package archive

import (
	"os"
	"testing"
)

// TestUnpackArchive tests the UnpackArchive function.
func TestUnpackArchive(t *testing.T) {
	type args struct {
		targetDir   string
		archiveFile string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "unpack-tar-gz",
			args: args{
				targetDir:   "./testdata/hello-tar-gz-unpack",
				archiveFile: "./testdata/hello.tar.gz",
			},
			wantErr: false,
		},
		{
			name: "unpack-tgz",
			args: args{
				targetDir:   "./testdata/hello-tgz-unpack",
				archiveFile: "./testdata/hello.tgz",
			},
			wantErr: false,
		},
		{
			name: "unpack-zip",
			args: args{
				targetDir:   "./testdata/hello-zip-unpack",
				archiveFile: "./testdata/hello.zip",
			},
			wantErr: false,
		},
		{
			name: "unsupported-archive-file",
			args: args{
				targetDir:   "./testdata/hello-unsupported",
				archiveFile: "./testdata/hello.txt",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer os.RemoveAll(tt.args.targetDir)
			if err := UnpackArchive(tt.args.targetDir, tt.args.archiveFile); (err != nil) != tt.wantErr {
				t.Errorf("UnpackArchive() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
