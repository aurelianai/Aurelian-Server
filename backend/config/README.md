## Model Configuration

This controls how prompts are formatted to ensure the desired output.

#### Example: [Together API](together.ai)

```yaml
model:
  system: "You are a helpful assistant"
  userPrefix: "human: "
  userPostfix: "\n"
  modelPrefix: "bot: "
  modelPostfix: "\n"
```

`system`

- The system prompt is inserted at the beginning of all prompts

`userPrefix`

- Prepended to all user inputs

`userPostfix`

- Appended to all user inputs

`modelPrefix`

- Prepended to all model inputs

`modelPostfix`

- Appended to all model inputs

`contextSize`

- Context size of the model

`chatContextSize`

- **[OPTIONAL]** Token limit for filling in with chats. Defaults to `contextSize`

## Inference Configuration

These settings are not related to the model, but instead about inference configuration

#### Example: TGI Dev Settings

```yaml
inference:
  backend: text-generation-inference
  endpoint: https://inf/generate_stream
  maxNewTokens: 50
```

`backend`

- This sets what inference software you are using and dictates the functions that are called to deal with it

`endpoint`

- Where the mfer is

`maxNewTokens`

- This cuts off generation in case it rambles endlessly
