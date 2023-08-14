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
	let signal: AbortSignal

	const handle_message_send = async (event: any) => {
		$ChatStore = [await new_chat(event.detail.message_content.slice(0, 22)), ...$ChatStore];
		chat = $ChatStore[0];

		messages = [...messages, await new_message(chat.ID, 'USER', event.detail.message_content)];
		generating = true;

		let model_response: Message = {"Role": "MODEL", Content: ""};
		messages = [...messages, model_response];

		const controller = new AbortController()
		signal = controller.signal

		for await (const newText of complete(chat.ID, signal)) {
			console.log(`RECV: ${newText}`)
			model_response.Content += newText
			messages[-1] = model_response
		}

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

			<div class="w-full h-16" />

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
