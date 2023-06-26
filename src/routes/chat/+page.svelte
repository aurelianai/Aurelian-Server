<script lang="ts">
	import { ProgressRadial } from '@skeletonlabs/skeleton';
	import MessageBubble from './MessageBubble.svelte';
	import ChatInput from './ChatInput.svelte';
	import ChatHeader from './ChatHeader.svelte';
	import ChatSuggestions from './ChatSuggestions.svelte';
	import GeneratingSpinner from './GeneratingSpinner.svelte';

	import type { Message } from '$lib/types';

	import { onMount } from 'svelte';
	import { list_messages, new_message, complete } from './crud';
	import { selected_session_id } from '$lib/stores';

	let messages: Message[];
	let generating: boolean = false;
	let bottom: HTMLDivElement;

	onMount(async () => {
		messages = await list_messages();
	});

	selected_session_id.subscribe(async () => {
		messages = await list_messages();
	});

	const handle_message_send = async (event: any) => {
		messages = [...messages, await new_message('USER', event.detail.message_content)];
		generating = true;
		setTimeout(() => {
			bottom.scrollIntoView({
				behavior: 'smooth',
				block: 'end',
				inline: 'nearest'
			});
		}, 100);

		let model_response: Message;
		try {
			model_response = await complete($selected_session_id);
		} catch (err) {
			generating = false;
			return;
		}

		messages = [...messages, model_response];
		generating = false;
	};
</script>

<div class="fixed right-0 flex-grow left-64 top-4">
	<ChatHeader disabled={generating} />
</div>

<div class="flex flex-col h-full bg-[url('/login-background.svg')] bg-tertiary-50-200-token">
	{#if messages === undefined}
		<div class="flex flex-col items-center justify-center w-full space-y-5 h-96">
			<code>Fetching your messages ...</code>
			<ProgressRadial width="w-10" />
		</div>
	{:else if messages.length !== 0}
		<div class="w-full p-5 space-y-3">
			<div class="w-full h-24" />
			{#each messages as message}
				<MessageBubble msg={message} />
			{/each}
			{#if generating}
				<GeneratingSpinner />
			{/if}
			<div class="w-full h-36" />
			<div bind:this={bottom} />
		</div>
	{:else}
		<div class="flex flex-col justify-center h-full">
			<ChatSuggestions />
		</div>
	{/if}
</div>

<div class="fixed right-0 flex-grow px-10 left-64 bottom-8">
	<ChatInput on:send_message={handle_message_send} disabled={generating} />
</div>
