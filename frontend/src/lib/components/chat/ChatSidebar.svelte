<script lang="ts">
	import { Avatar, LightSwitch, popup } from '@skeletonlabs/skeleton';
	import ChatSidebarItem from './ChatSidebarItem.svelte';
	import type { PopupSettings } from '@skeletonlabs/skeleton';
	import type { Chat } from '$lib/types';
	import { delete_chat, ChatStore } from '$lib/ts/chat/util';
	import { goto } from '$app/navigation';

	export let user = { id: 0, email: 'Hello' };
	export let chats: Chat[];
	ChatStore.subscribe((c) => (chats = c));

	const delete_event_handler = async (event: any) => {
		// TODO better error handling here
		await delete_chat(event.detail.id);
		chats = chats.filter((item) => {
			return item.ID !== event.detail.id;
		});
		goto('/chat');
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
		class="w-full h-10 font-bold rounded-md btn text-md variant-filled-primary"
		on:click={async () => {
			goto(`/chat`);
		}}
	>
		New Chat
	</button>
</div>

<div class="px-3 pt-2">
	<nav>
		<ul class="space-y-1">
			<div class="w-full h-16" />
			{#each chats as chat}
				<!-- svelte-ignore a11y-click-events-have-key-events -->
				<li>
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
			<p class="font-semibold">
				{user.email.split('@')[0].slice(0, 12)}
			</p>
		</button>

		<LightSwitch rounded="rounded-md" />
	</div>
</div>

<!--Popups and Modals-->
<div class="w-40 space-y-2 rounded-md card" data-popup="userPopupBox">
	<div
		class="flex items-center w-full p-2 space-x-2 rounded-md justify-left hover:variant-soft-surface"
	>
		<img src="/logout.svg" class="h-4" alt="logout" />
		<div class="font-medium">Log Out</div>
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
