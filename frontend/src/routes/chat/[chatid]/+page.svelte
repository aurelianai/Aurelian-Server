<script lang="ts">
	import MessageBubble from '$lib/components/chat/MessageBubble.svelte';
	import ChatInput from '$lib/components/chat/ChatInput.svelte';
	import type { Message } from '$lib/types';
	import type { PageData } from './$types';
	import { new_message, complete } from '$lib/ts/chat/util';

	export let data: PageData;
	let generating: boolean = false;
	let signal: AbortSignal

	const handle_message_send = async (event: any) => {
		data.messages = [
			...data.messages,
			await new_message(data.chatid, 'USER', event.detail.message_content)
		];
		generating = true;

		let model_response: Message = {"Role": "MODEL", Content: ""};
		data.messages = [...data.messages, model_response];

		const controller = new AbortController()
		signal = controller.signal

		for await (const newText of complete(data.chatid, signal)) {
			console.log(`RECV: ${newText}`)
			model_response.Content += newText.delta
			data.messages[-1] = model_response
		}

		generating = false;
	};
</script>

<svelte:head>
	<title>{`"`+data.chatTitle+`"` || 'Aurelian — Chat'}</title>
</svelte:head>

<div class="flex flex-col h-full">
		<div class="w-full p-5 space-y-3">
			{#each data.messages as message}
				<MessageBubble msg={message} />
			{/each}
			<div class="w-full h-16" />
		</div>
</div>

<div class="fixed right-0 flex justify-center flex-grow left-64 bottom-4">
	<ChatInput on:send_message={handle_message_send} disabled={generating} />
</div>
