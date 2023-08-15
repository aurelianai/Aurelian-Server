package config

import "testing"

func TestInitAndValidate(t *testing.T) {
	type args struct {
		configFilePath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Complete and Correct Config File",
			args:    args{configFilePath: "./testfiles/correct.yml"},
			wantErr: false,
		}, {
			name:    "Correct but missing endpoint",
			args:    args{configFilePath: "./testfiles/missing_endpoint.yml"},
			wantErr: true,
		}, {
			name:    "0 Context Size",
			args:    args{configFilePath: "./testfiles/bad_context_size.yml"},
			wantErr: true,
		}, {
			name:    "Chat CTX Larger than Model CTX",
			args:    args{configFilePath: "./testfiles/chat_ctx_absurd.yml"},
			wantErr: true,
		}, {
			name:    "Set Chat CTX",
			args:    args{configFilePath: "./testfiles/set_chat_ctx.yml"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := &AelsConfig{}
			if err := config.InitAndValidate(tt.args.configFilePath); (err != nil) != tt.wantErr {
				t.Errorf("InitAndValidate() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.name == "Set Chat CTX" {
				if config.Model.ChatContextSize != config.Model.ContextSize {
					t.Errorf("InitAndValidate err, ChatContextSize of %d tokens should be have been inferred from context entry", config.Model.ContextSize)
				}
			}
		})
	}
}
