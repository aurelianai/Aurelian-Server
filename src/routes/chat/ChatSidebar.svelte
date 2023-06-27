<script lang="ts">
	import { Avatar, LightSwitch, ListBox, ListBoxItem, popup } from '@skeletonlabs/skeleton';
	import ChatSidebarItem from './ChatSidebarItem.svelte';

	import type { PopupSettings } from '@skeletonlabs/skeleton';
	import type { Chat, User } from '$lib/types';

	import { onMount } from 'svelte';
	import { selected_session_id } from '$lib/stores';
	import { list_chats, new_chat, delete_chat } from './crud';

	export let user: User;

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

	const userPopupBox: PopupSettings = {
		event: 'click',
		target: 'userPopupBox',
		placement: 'top',
		closeQuery: ''
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

<div class="px-3 pt-2 space-y-5">
	<nav class="space-y-4 list-nav">
		<ul>
			<div class="w-full h-16" />
			{#each chats as chat}
				<!-- svelte-ignore a11y-click-events-have-key-events -->
				<li on:click={() => selected_session_id.set(chat.id)}>
					<ChatSidebarItem {chat} on:delete={delete_event_handler} />
				</li>
			{/each}
			<div class="w-full h-16" />
		</ul>
	</nav>
</div>

<div class="fixed bottom-0 left-0 w-64 px-3 pb-2 space-y-3 menu-button-bg">
	<hr class="opacity-100" />
	<div class="flex items-center w-full space-x-3">
		<!-- TODO Popup Here-->
		<button
			class="flex items-center w-40 p-2 space-x-2 rounded-md hover:variant-soft-surface"
			use:popup={userPopupBox}
		>
			<Avatar
				initials={user.email[0]}
				background="variant-filled-primary"
				width="w-10"
				rounded="rounded-md"
				fill="fill-white"
			/>
			<p class="font-bold">
				{user.email.split('@')[0].slice(0, 12)}
			</p>
		</button>

		<LightSwitch />
	</div>
</div>

<!--Popups and Modals-->
<div class="w-40 space-y-2 rounded-md card" data-popup="userPopupBox">
	<a
		href="/logout"
		class="flex items-center justify-center w-full p-2 space-x-2 rounded-md hover:variant-soft-surface"
	>
		<img src="/logout.svg" class="h-4" alt="logout" />
		<div class="font-bold">Log Out</div>
	</a>
</div>

<style>
	.menu-button-bg {
		background-color: #dfe0e2;
	}
	:is(.dark .menu-button-bg) {
		background-color: #16171e;
	}
</style>
