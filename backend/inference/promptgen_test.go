package inference

import (
	"AELS/config"
	"AELS/persistence"
	"fmt"
	"testing"
)

func TestGeneratePrompt(t *testing.T) {

	config.Config.InitAndValidate("./testconfig.yml")

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
			want: fmt.Sprintf("%s%sHello, how's it going?%s%sI'm doing well, how can I help you?%s%sI need help writing an email%s%s",
				config.Config.Model.System,
				config.Config.Model.UserPrefix,
				config.Config.Model.UserPostfix,
				config.Config.Model.ModelPrefix,
				config.Config.Model.ModelPostfix,
				config.Config.Model.UserPrefix,
				config.Config.Model.UserPostfix,
				config.Config.Model.ModelPrefix,
			),
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
