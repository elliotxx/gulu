package json

import (
	"reflect"
	"testing"
)

type TestStuct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestMustMarshal(t *testing.T) {
	type args struct {
		v interface{}
	}

	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "t1",
			args: args{
				v: &TestStuct{
					Name: "Tony",
					Age:  20,
				},
			},
			want: []byte(`{"name":"Tony","age":20}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MustMarshal(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MustMarshal() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

func TestMustPrettyMarshal(t *testing.T) {
	type args struct {
		v interface{}
	}

	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "t1",
			args: args{
				v: &TestStuct{
					Name: "Tony",
					Age:  20,
				},
			},
			want: []byte(`{
  "name": "Tony",
  "age": 20
}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MustPrettyMarshal(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MustPrettyMarshal() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

func TestMustMarshalString(t *testing.T) {
	type args struct {
		v interface{}
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "t1",
			args: args{
				v: &TestStuct{
					Name: "Tony",
					Age:  20,
				},
			},
			want: `{"name":"Tony","age":20}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MustMarshalString(tt.args.v); got != tt.want {
				t.Errorf("MustMarshalString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustPrettyMarshalString(t *testing.T) {
	type args struct {
		v interface{}
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "t1",
			args: args{
				v: &TestStuct{
					Name: "Tony",
					Age:  20,
				},
			},
			want: `{
  "name": "Tony",
  "age": 20
}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MustPrettyMarshalString(tt.args.v); got != tt.want {
				t.Errorf("MustPrettyMarshalString() = %v, want %v", got, tt.want)
			}
		})
	}
}
