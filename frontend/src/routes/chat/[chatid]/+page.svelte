<script lang="ts">
	import MessageBubble from '$lib/components/chat/MessageBubble.svelte';
	import ChatInput from '$lib/components/chat/ChatInput.svelte';
	import ChatSuggestions from '$lib/components/chat/ChatSuggestions.svelte';
	import GeneratingSpinner from '$lib/components/chat/GeneratingSpinner.svelte';
	import type { Message } from '$lib/types';
	import type { PageData } from './$types';
	import { ChatStore, new_message, complete } from '$lib/ts/chat/util';

	export let data: PageData;
	let generating: boolean = false;
	let bottom: HTMLDivElement;

	const handle_message_send = async (event: any) => {
		data.messages = [
			...data.messages,
			await new_message(data.chatid, 'USER', event.detail.message_content)
		];
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
			model_response = await complete(data.chatid);
		} catch (err) {
			generating = false;
			return;
		}

		data.messages = [...data.messages, model_response];
		generating = false;
	};
</script>

<svelte:head>
	<title>{data.chatTitle || 'Aurelian — Chat'}</title>
</svelte:head>

<div class="flex flex-col h-full">
	{#if data.messages.length !== 0}
		<div class="w-full p-5 space-y-3">
			{#each data.messages as message}
				<MessageBubble msg={message} />
			{/each}
			{#if generating}
				<GeneratingSpinner />
			{/if}
			<div class="w-full h-16" />
			<div bind:this={bottom} />
		</div>
	{:else}
		<div class="flex flex-col justify-center h-full">
			<ChatSuggestions />
		</div>
	{/if}
</div>

<div class="fixed right-0 flex justify-center flex-grow left-64 bottom-4">
	<ChatInput on:send_message={handle_message_send} disabled={generating} />
</div>
