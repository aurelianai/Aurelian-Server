<script lang="ts">
	import type { Message } from '$lib/types';

	export let msg: Message;
</script>

{#if msg.role === 'USER'}
	<div class="max-w-5xl p-4 mx-auto space-y-2 rounded-tr-none card">
		<header class="flex items-center justify-between">
			<p class="font-bold">{msg.role}</p>
		</header>
		<p class="whitespace-pre-wrap">{msg.content}</p>
	</div>
{:else}
	<!-- svelte-ignore a11y-click-events-have-key-events -->
	<div class="max-w-5xl p-4 mx-auto space-y-2 rounded-tl-none card bg-secondary-300">
		<header class="flex items-center justify-between">
			<p class="font-bold">{msg.role}</p>
		</header>
		<p class="whitespace-pre-wrap">{msg.content}</p>
	</div>
{/if}
