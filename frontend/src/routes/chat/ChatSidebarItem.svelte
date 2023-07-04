<script lang="ts">
	import type { Chat } from '$lib/types';

	import { createEventDispatcher } from 'svelte';
	import { focusTrap } from '@skeletonlabs/skeleton';
	import { selected_session_id } from '$lib/stores';
	import { update_chat } from './crud';

	export let chat: Chat;
	let isActive: boolean = $selected_session_id === chat.id;
	$: isActive = $selected_session_id === chat.id;
	let hover: boolean = false;
	let confirm_delete: boolean = false;
	let confirm_edit: boolean = false;
	let new_name = '';

	const dispatch = createEventDispatcher();
</script>

<!-- svelte-ignore a11y-mouse-events-have-key-events -->
<div
	class:variant-ringed-primary={isActive}
	class:variant-soft-surface={hover && !isActive}
	class="flex items-center w-full p-2 font-medium rounded-md text-md"
	on:mouseenter={() => (hover = true)}
	on:mouseleave={() => (hover = false)}
>
	<span class="h-1 mr-1 badge bg-primary-500" />
	{#if !confirm_edit}
		<span class="flex-auto w-48 h-6 pr-2 overflow-hidden">
			{chat.name}
		</span>
	{:else}
		<div>
			<input
				use:focusTrap={confirm_edit}
				class="input variant-form-material"
				bind:value={new_name}
			/>
		</div>
	{/if}

	{#if isActive && !confirm_delete && !confirm_edit}
		<button
			class="w-6 h-6 space-x-1 bg-transparent btn"
			on:click={() => {
				confirm_edit = true;
				new_name = chat.name;
			}}
		>
			✏️
		</button>
		<button class="w-6 h-6 bg-transparent btn" on:click={() => (confirm_delete = true)}>
			🗑️
		</button>
	{/if}
	{#if confirm_delete && isActive}
		<div class="flex flex-row">
			<button
				class="w-6 h-6 bg-transparent btn"
				on:click={() => {
					confirm_delete = false;
					dispatch('delete', { id: chat.id });
				}}
			>
				✅
			</button>
			<button class="w-6 h-6 bg-transparent btn" on:click={() => (confirm_delete = false)}>
				❌
			</button>
		</div>
	{/if}
	{#if confirm_edit}
		<div class="flex flex-row">
			<button
				class="w-6 h-6 bg-transparent btn"
				on:click={async () => {
					if (chat.name === new_name) {
						confirm_edit = false;
						return;
					}
					await update_chat(chat.id, new_name);
					chat.name = new_name;
					confirm_edit = false;
				}}
			>
				✅
			</button>
			<button class="w-6 h-6 bg-transparent btn" on:click={() => (confirm_edit = false)}>
				❌
			</button>
		</div>
	{/if}
</div>
