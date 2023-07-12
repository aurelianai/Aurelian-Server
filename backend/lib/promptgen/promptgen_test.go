package promptgen

import (
	"AELS/persistence"
	"testing"
)

func TestGeneratePrompt(t *testing.T) {
	type args struct {
		messages []persistence.Message
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Three messages",
			args: args{
				messages: []persistence.Message{
					{
						Role:    "USER",
						Content: "Hello, how's it going?",
					},
					{
						Role:    "MODEL",
						Content: "I'm doing well, how can I help you?",
					},
					{
						Role:    "USER",
						Content: "I need help writing an email",
					},
				},
			},
			want:    "<|prompter|>Hello, how's it going?<|endoftext|><|assistant|>I'm doing well, how can I help you?<|endoftext|><|prompter|>I need help writing an email<|endoftext|><|assistant|>",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GeneratePrompt(tt.args.messages)
			if (err != nil) != tt.wantErr {
				t.Errorf("GeneratePrompt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GeneratePrompt() = %v, want %v", got, tt.want)
			}
		})
	}
}
