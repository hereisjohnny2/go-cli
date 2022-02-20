package messages

import "testing"

func TestMessage_Run(t *testing.T) {
	type fields struct {
		Text  string
		Helpf bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "should show long help if Helpf is true",
			fields: fields{
				Text:  "test",
				Helpf: true,
			},
		},
		{
			name: "should alert that text data is required",
			fields: fields{
				Text:  "",
				Helpf: false,
			},
		},
		{
			name: "should log a valid text",
			fields: fields{
				Text:  "test",
				Helpf: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &Message{
				Text:  tt.fields.Text,
				Helpf: tt.fields.Helpf,
			}
			cmd.Run()
		})
	}
}
