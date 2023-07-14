<script lang="ts">
	import { Avatar, LightSwitch, popup } from '@skeletonlabs/skeleton';
	import { Icon, ArrowLeftOnRectangle } from 'svelte-hero-icons';
	import ChatSidebarItem from './ChatSidebarItem.svelte';
	import type { PopupSettings } from '@skeletonlabs/skeleton';
	import type { User } from '$lib/types';
	import { delete_chat, ChatStore } from '$lib/ts/chat/util';
	import { goto } from '$app/navigation';

	export let user: User;

	const delete_event_handler = async (event: any) => {
		await delete_chat(event.detail.id);
		$ChatStore = $ChatStore.filter((item) => {
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

	const logout = async () => {
		const res = await fetch('/api/logout', {
			headers: new Headers({ 'Content-Type': 'application/json' }),
			method: 'POST'
		});
		if (res.status == 200) {
			goto('/login');
		}
	};
</script>

<div class="fixed top-0 left-0 w-64 px-3 pt-4 menu-button-bg">
	<button
		class="w-full h-10 font-bold rounded-md btn text-md variant-filled-primary"
		on:click={async () => {
			goto('/chat');
		}}
	>
		New Chat
	</button>
</div>

<div class="px-3 pt-2">
	<nav>
		<ul class="space-y-1">
			<div class="w-full h-16" />
			{#each $ChatStore as chat}
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
			class="flex items-center justify-center p-2 space-x-2 rounded-md hover:variant-soft-surface"
			use:popup={userPopupBox}
		>
			<Avatar
				initials={user.FirstName.charAt(0) + user.LastName.charAt(0)}
				background="variant-filled-primary"
				class="w-7"
				rounded="rounded-md"
				fill="fill-white"
			/>
			<p class="w-24 h-6 font-semibold truncate">
				{`${user.FirstName} ${user.LastName}`}
			</p>
		</button>

		<LightSwitch rounded="rounded-md" />
	</div>
</div>

<!--Popups and Modals-->
<button class="w-40 space-y-2 rounded-md card" data-popup="userPopupBox" on:click={logout}>
	<div
		class="flex items-center w-full p-2 space-x-2 rounded-md justify-left hover:variant-soft-surface"
	>
		<Icon src={ArrowLeftOnRectangle} class="w-5" />
		<div class="font-medium">Log Out</div>
	</div>
</button>

<style>
	.menu-button-bg {
		background-color: #dfe0e2;
	}
	:is(.dark .menu-button-bg) {
		background-color: #16171e;
	}
</style>
