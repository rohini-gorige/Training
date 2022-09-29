package logger

import (
	"errors"
	"testing"
)

type mockfeilds struct {
}

func (*mockfeilds) Close() error {
	return nil
}

func (*mockfeilds) Rotate() error {
	return nil
}

func (*mockfeilds) Write(p []byte) (n int, err error) {
	if string(p) == "test" {
		return 4, nil
	}
	return 0, errors.New("faild")
}

func TestLogger(t *testing.T) {
	z, err := NewLogService("./", "demo.txt")
	if err != nil {
		t.Errorf("should be nil")
	}

	z = &mockfeilds{}
	err = z.Close()
	if err != nil {
		t.Errorf("should be nil")
	}

	err = z.Rotate()
	if err != nil {
		t.Errorf("should be nil")
	}
	_, err = z.Write([]byte("test"))
	if err != nil {
		t.Errorf("should be nil")
	}
	_, err = z.Write([]byte("error"))
	if err == nil {
		t.Errorf("should not be nil")
	}
}

// func TestLoadConfiguration(t *testing.T) {
// 	type args struct {
// 		Filename string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    Config
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name: "Logger",
// 			args: args{
// 				Filename: "demo.txt",
// 			},
// 			want: Config{
// 				Filename:   "demo.txt",
// 				MaxAge:     0,
// 				MaxBackups: 1,
// 				MaxSize:    1,
// 				LocalTime:  true,
// 				Compress:   true,
// 			},
// 			wantErr: false,
// 		}, {
// 			want: Config{
// 				Filename:   "demo.txt",
// 				MaxAge:     0,
// 				MaxBackups: 1,
// 				MaxSize:    1,
// 				LocalTime:  true,
// 				Compress:   true,
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := LoadConfiguration(tt.args.Filename)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("LoadConfiguration() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("LoadConfiguration() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_compress(t *testing.T) {
// 	type args struct {
// 		name string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    bool
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 		{"test", args{"demo.txt"}, true, false},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := compress(tt.args.name)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("compress() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if got != tt.want {
// 				t.Errorf("compress() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
