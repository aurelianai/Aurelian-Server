<script lang="ts">
	import { CodeBlock } from '@skeletonlabs/skeleton';

	export let text = '';
	export let lang = '';
</script>

<CodeBlock
	language={lang || 'plaintext'}
	code={text || 'Empty Code Block'}
	rounded="rounded-md"
/>
