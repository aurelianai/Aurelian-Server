<script lang="ts">
	import SvelteMarkdown from 'svelte-markdown';
	import Code from '$lib/components/markdown/Code.svelte';
	import Table from '$lib/components/markdown/Table.svelte';
	import List from '$lib/components/markdown/List.svelte';
	import ListItem from '$lib/components/markdown/ListItem.svelte';
	import { Avatar } from '@skeletonlabs/skeleton';

	import type { Message } from '$lib/types';

	export let msg: Message;
	let user = { ID: 0, FirstName: 'Ethan', LastName: 'Steere', email: 'ethansteere1@gmail.com' };
</script>

{#if msg.Role === 'USER'}
	<div class="flex items-start max-w-3xl p-6 mx-auto space-x-5 rounded-md shadow-xl card">
		<div class="w-10 h-full">
			<Avatar
				initials={user.FirstName.charAt(0) + user.LastName.charAt(0)}
				width="w-8"
				rounded="rounded-md"
				fill="fill-white"
				background="bg-primary-500"
			/>
		</div>

		<p class="break-all whitespace-pre-wrap">{msg.Content}</p>
	</div>
{:else}
	<div
		class="flex justify-start max-w-3xl p-6 mx-auto space-x-5 rounded-md shadow-xl card model-bg"
	>
		<Avatar src="/logo.svg" width="w-8" rounded="rounded-md" background="bg-inherit" />
		<SvelteMarkdown
			source={msg.Content}
			renderers={{ code: Code, table: Table, list: List, listitem: ListItem }}
		/>
	</div>
{/if}

<style lang="postcss">
	.model-bg {
		@apply bg-secondary-300;
	}
	:is(.dark .model-bg) {
		@apply bg-secondary-500;
	}
</style>
