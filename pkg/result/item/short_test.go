package item

import "testing"

func TestShort_Text(t *testing.T) {
	type fields struct {
		Name        string
		Description string
		Result      string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "seccomp",
			fields: fields{
				Name:        "mode",
				Description: "filter means xxx",
				Result:      "filter",
			},
			want: "mode:\tfilter # filter means xxx",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Short{
				Name:        tt.fields.Name,
				Description: tt.fields.Description,
				Result:      tt.fields.Result,
			}
			got := s.Text()
			println(got)
			if got != tt.want {
				t.Errorf("Text() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShort_Colorful(t *testing.T) {
	type fields struct {
		Name        string
		Description string
		Result      string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "seccomp",
			fields: fields{
				Name:        "mode",
				Description: "filter means xxx",
				Result:      "filter",
			},
			want: "mode:\tfilter # filter means xxx",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Short{
				Name:        tt.fields.Name,
				Description: tt.fields.Description,
				Result:      tt.fields.Result,
			}
			got := s.Colorful()
			println(got)
		})
	}
}
