<script lang="ts">
	import MessageBubble from '$lib/components/chat/MessageBubble.svelte';
	import ChatInput from '$lib/components/chat/ChatInput.svelte';
	import ChatSuggestions from '$lib/components/chat/ChatSuggestions.svelte';
	import GeneratingSpinner from '$lib/components/chat/GeneratingSpinner.svelte';
	import type { Chat, Message } from '$lib/types';
	import { ChatStore, new_message, complete, new_chat } from '$lib/ts/chat/util';
	import { goto } from '$app/navigation';

	export let messages: Message[] = [];
	let chat: Chat | null = null;
	let generating: boolean = false;
	let bottom: HTMLDivElement;

	const handle_message_send = async (event: any) => {
		$ChatStore = [await new_chat(), ...$ChatStore];
		chat = $ChatStore[0];
		// TODO automatically name chat

		messages = [...messages, await new_message(chat.ID, 'USER', event.detail.message_content)];
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
			model_response = await complete(chat.ID);
		} catch (err) {
			// TODO display error when completion fails
			generating = false;
			return;
		}

		messages = [...messages, model_response];
		generating = false;
		goto(`/chat/${chat.ID}`);
	};
</script>

<svelte:head>
	<title>Aurelian — New Chat</title>
</svelte:head>

<div class="flex flex-col h-full">
	{#if messages.length !== 0}
		<div class="w-full p-5 space-y-3">
			{#each messages as message}
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
