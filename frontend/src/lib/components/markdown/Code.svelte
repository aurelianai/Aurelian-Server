<script lang="ts">
	import { CodeBlock } from '@skeletonlabs/skeleton';

	export let text = '';
	export let lang = '';
</script>

<div class="flex justify-center w-full px-4">
	<CodeBlock
		language={lang || 'plaintext'}
		code={text || 'Empty Code Block'}
		lineNumbers={true}
		rounded="rounded-md"
	/>
</div>
