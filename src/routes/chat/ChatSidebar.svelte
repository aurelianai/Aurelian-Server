<script lang="ts">
	import ChatSidebarItem from './ChatSidebarItem.svelte';

	import type { Chat } from '$lib/types';

	import { onMount } from 'svelte';
	import { selected_session_id } from '$lib/stores';
	import { list_chats, new_chat, delete_chat } from '$lib/chat/crud';

	let chats: Chat[] = [];
	onMount(async () => {
		chats = await list_chats();
		if (chats.length == 0) {
			chats = [await new_chat()];
		}
		selected_session_id.set(chats[0].id);
	});

	const delete_event_handler = async (event: any) => {
		chats = chats.filter((item) => {
			return item.id !== event.detail.id;
		});
		await delete_chat(event.detail.id);
		if (chats.length == 0) {
			chats = [await new_chat()];
		}
		selected_session_id.set(chats[0].id);
	};
</script>

<div class="px-3 pt-2 space-y-5">
	<button
		class="w-full mx-auto mt-2 font-bold rounded-md btn text-md variant-filled-primary"
		on:click={async () => {
			let chat_to_add = await new_chat();
			chats = [chat_to_add, ...chats];
			$selected_session_id = chat_to_add.id;
		}}
	>
		Create New Chat
	</button>
	<nav class="space-y-4 list-nav">
		<ul>
			{#each chats as chat}
				<!-- svelte-ignore a11y-click-events-have-key-events -->
				<li class="py-1" on:click={() => selected_session_id.set(chat.id)}>
					<ChatSidebarItem {chat} on:delete={delete_event_handler} />
				</li>
			{/each}
		</ul>
	</nav>
</div>