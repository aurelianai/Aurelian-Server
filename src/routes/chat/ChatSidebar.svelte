<script lang="ts">
	import { Avatar, LightSwitch } from '@skeletonlabs/skeleton';
	import ChatSidebarItem from './ChatSidebarItem.svelte';

	import type { Chat } from '$lib/types';

	import { onMount } from 'svelte';
	import { selected_session_id } from '$lib/stores';
	import { list_chats, new_chat, delete_chat } from './crud';

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

<div class="fixed top-0 left-0 w-64 px-3 pt-4 menu-button-bg">
	<button
		class="w-full font-bold rounded-md btn text-md variant-filled-primary"
		on:click={async () => {
			const chat_to_add = await new_chat();
			chats = [chat_to_add, ...chats];
			$selected_session_id = chat_to_add.id;
		}}
	>
		New Chat
	</button>
</div>

<hr class="opacity-100" />

<div class="px-3 pt-2 space-y-5">
	<nav class="space-y-4 list-nav">
		<ul>
			<div class="w-full h-16" />
			{#each chats as chat}
				<!-- svelte-ignore a11y-click-events-have-key-events -->
				<li class="py-1" on:click={() => selected_session_id.set(chat.id)}>
					<ChatSidebarItem {chat} on:delete={delete_event_handler} />
				</li>
			{/each}
			<div class="w-full h-16" />
		</ul>
	</nav>
</div>

<div class="fixed bottom-0 left-0 w-64 px-3 pb-2 space-y-3 menu-button-bg">
	<hr class="opacity-100" />
	<div class="flex items-center w-full space-x-5">
		<Avatar
			initials="ES"
			background="variant-filled-primary"
			width="w-10"
			rounded="rounded-md"
			fill="fill-white"
		/>
		<LightSwitch />
	</div>
</div>

<style>
	.menu-button-bg {
		background-color: #dfe0e2;
	}
	:is(.dark .menu-button-bg) {
		background-color: #16171e;
	}
</style>
