<script lang="ts">
	import {
		Icon,
		ChatBubbleBottomCenterText,
		PencilSquare,
		Trash,
		XCircle,
		CheckCircle
	} from 'svelte-hero-icons';
	import type { Chat } from '$lib/types';
	import { createEventDispatcher } from 'svelte';
	import { focusTrap } from '@skeletonlabs/skeleton';
	import { update_chat } from '$lib/ts/chat/util';
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';

	export let chat: Chat;
	let isActive = false;
	let hover: boolean = false;
	let confirm_delete: boolean = false;
	let confirm_edit: boolean = false;
	let new_name = '';

	page.subscribe((p) => {
		isActive = p.url.pathname === `/chat/${chat?.ID}`;
	});

	const dispatch = createEventDispatcher();
</script>

<!-- svelte-ignore a11y-mouse-events-have-key-events -->
<div
	class:variant-ringed-primary={isActive}
	class:variant-soft-surface={hover && !isActive}
	class="flex items-center w-full h-10 p-2 font-medium rounded-md text-md"
	on:mouseenter={() => (hover = true)}
	on:mouseleave={() => (hover = false)}
>
	{#if !confirm_edit}
		<button
			class="flex items-center justify-start w-48 space-x-2"
			on:click={() => {
				goto(`/chat/${chat.ID}/`);
			}}
		>
			<Icon src={ChatBubbleBottomCenterText} class="w-5 h-5 text-secondary-500-400-token" solid />
			<span class="flex truncate {isActive ? 'w-32' : 'w-40'}">{chat.Title}</span>
		</button>
	{:else}
		<div class="pr-2">
			<input
				use:focusTrap={confirm_edit}
				class="input variant-form-material"
				bind:value={new_name}
			/>
		</div>
	{/if}

	{#if isActive && !confirm_delete && !confirm_edit}
		<div class="flex items-center justify-center space-x-1">
			<button
				on:click={() => {
					confirm_edit = true;
					new_name = chat.Title;
				}}
			>
				<Icon src={PencilSquare} class="w-5 h-5 text-surface-400 hover:brightness-150" />
			</button>
			<button on:click={() => (confirm_delete = true)}>
				<Icon src={Trash} class="w-5 h-5 text-surface-400 hover:brightness-150" />
			</button>
		</div>
	{/if}

	{#if confirm_delete && isActive}
		<div class="flex space-x-1">
			<button
				on:click={() => {
					confirm_delete = false;
					dispatch('delete', { id: chat.ID });
				}}
			>
				<Icon src={CheckCircle} class="w-5 h-5 text-surface-400 hover:brightness-150" />
			</button>
			<button on:click={() => (confirm_delete = false)}>
				<Icon src={XCircle} class="w-5 h-5 text-surface-400 hover:brightness-150" />
			</button>
		</div>
	{/if}
	{#if confirm_edit}
		<div class="flex space-x-1">
			<button
				on:click={async () => {
					if (chat.Title === new_name) {
						confirm_edit = false;
						return;
					}
					// TODO Error handling
					await update_chat(chat.ID, new_name);
					chat.Title = new_name;
					confirm_edit = false;
				}}
			>
				<Icon src={CheckCircle} class="w-5 h-5 text-surface-400 hover:brightness-150" />
			</button>
			<button on:click={() => (confirm_edit = false)}>
				<Icon src={XCircle} class="w-5 h-5 text-surface-400 hover:brightness-150" />
			</button>
		</div>
	{/if}
</div>
