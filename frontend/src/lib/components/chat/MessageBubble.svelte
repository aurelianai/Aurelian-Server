<script lang="ts">
	import SvelteMarkdown from 'svelte-markdown';
	import Code from '$lib/components/markdown/Code.svelte';
	import CodeSpan from '$lib/components/markdown/CodeSpan.svelte';
	import Table from '$lib/components/markdown/Table.svelte';
	import List from '$lib/components/markdown/List.svelte';
	import ListItem from '$lib/components/markdown/ListItem.svelte';
	import { Avatar } from '@skeletonlabs/skeleton';
	import type { Message } from '$lib/types';
	import { UserStore } from '$lib/ts/chat/util';

	export let msg: Message;
</script>

{#if msg.Role === 'USER'}
	<div class="flex justify-center max-w-2xl p-6 mx-auto space-x-5 rounded-md shadow-xl card">
		<div class="w-8 h-8">
			<Avatar
				initials={$UserStore.FirstName.charAt(0) + $UserStore.LastName.charAt(0)}
				width="w-8 h-8"
				rounded="rounded-md"
				fill="fill-white"
				background="bg-primary-500"
			/>
		</div>

		<div class="flex flex-col w-full">
			<p class="break-all">{msg.Content}</p>
		</div>

		<div class="w-8" />
	</div>
{:else}
	<div
		class="flex justify-center max-w-2xl p-6 mx-auto space-x-5 rounded-md shadow-xl model-bg card"
	>
		<div class="w-8 h-8">
			<Avatar src="/logo.svg" width="w-8 h-8" rounded="rounded-md" background="bg-inherit" />
		</div>

		<div class="flex flex-col w-full space-y-5">
			<SvelteMarkdown
				source={msg.Content}
				renderers={{
					code: Code,
					codespan: CodeSpan,
					table: Table,
					list: List,
					listitem: ListItem
				}}
			/>
		</div>

		<div class="w-8" />
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
