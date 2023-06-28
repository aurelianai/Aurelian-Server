<script lang="ts">
	import SvelteMarkdown from 'svelte-markdown';
	import Code from '$lib/markdown-components/Code.svelte';
	import Table from '$lib/markdown-components/Table.svelte';

	import type { Message } from '$lib/types';

	export let msg: Message;
</script>

{#if msg.role === 'USER'}
	<div class="max-w-5xl p-4 mx-auto space-y-2 rounded-tr-none shadow-xl card">
		<header class="flex items-center justify-between">
			<p class="font-bold">{msg.role}</p>
		</header>
		<p class="whitespace-pre-wrap">{msg.content}</p>
	</div>
{:else}
	<div
		class="max-w-5xl p-4 mx-auto space-y-2 rounded-tl-none shadow-xl card bg-secondary-300-600-token"
	>
		<header class="flex items-center justify-between">
			<p class="font-bold">{msg.role}</p>
		</header>
		<SvelteMarkdown source={msg.content} renderers={{ code: Code, table: Table }} />
	</div>
{/if}
