package utils

import "testing"

func TestIsValidBuiltinType(t *testing.T) {
	type args struct {
		valueType string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "bool_1",
			args: args{valueType: "bool"},
			want: true,
		},
		{
			name: "bool_2",
			args: args{valueType: "Bool"},
			want: true,
		},
		{
			name: "bool_3",
			args: args{valueType: "boolean"},
			want: false,
		},
		{
			name: "int_1",
			args: args{valueType: "INT"},
			want: true,
		},
		{
			name: "int_2",
			args: args{valueType: "int"},
			want: true,
		},
		{
			name: "float_1",
			args: args{valueType: "Float"},
			want: true,
		},
		{
			name: "string_1",
			args: args{valueType: "string"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidBuiltinType(tt.args.valueType); got != tt.want {
				t.Errorf("IsValidBuiltinType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidValue(t *testing.T) {
	type args struct {
		valueType string
		value     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "bool_normal",
			args: args{valueType: "bool", value: "TRUE"},
			want: true,
		},
		{
			name: "bool_abnormal",
			args: args{valueType: "bool", value: "11"},
			want: false,
		},
		{
			name: "int_1",
			args: args{valueType: "int", value: "-1"},
			want: true,
		},
		{
			name: "int_2",
			args: args{valueType: "int", value: "111111111111111111111111111111111111111111111"},
			want: false,
		},
		{
			name: "float_1",
			args: args{valueType: Float, value: "0"},
			want: true,
		},
		{
			name: "float_2",
			args: args{valueType: Float, value: "-0.1"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidValue(tt.args.valueType, tt.args.value); got != tt.want {
				t.Errorf("IsValidValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
