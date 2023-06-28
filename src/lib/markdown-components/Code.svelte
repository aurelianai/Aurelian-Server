<script lang="ts">
	import { CodeBlock } from '@skeletonlabs/skeleton';

	export let text = '';
	export let lang = '';
</script>

<div class="p-5">
	<CodeBlock
		language={lang || 'plaintext'}
		code={text || 'Empty Code Block'}
		lineNumbers={false}
		rounded="rounded-md"
	/>
</div>
