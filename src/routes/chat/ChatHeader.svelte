<script lang="ts">
	import type { PopupSettings } from '@skeletonlabs/skeleton';
	import type { TextGenModel } from '$lib/types';

	import { popup } from '@skeletonlabs/skeleton';
	import { onMount } from 'svelte';
	import { list_models } from './crud';
	import { selected_model_id } from '$lib/stores';

	export let disabled: boolean;
	let models: TextGenModel[];
	let selected_model: TextGenModel;

	onMount(async () => {
		models = await list_models();
		selected_model = models[0];
		$selected_model_id = selected_model.id;
	});

	const modelListPopup: PopupSettings = {
		event: 'click',
		target: 'modelListPopup',
		placement: 'bottom',
		closeQuery: '.close-class'
	};
</script>

<div
	class="flex w-64 p-4 mx-auto rounded-lg shadow-xl bg-surface-100-800-token"
	use:popup={modelListPopup}
>
	<button class="justify-between w-56 py-2 rounded-lg btn variant-filled-tertiary" {disabled}>
		<span>{selected_model?.name}</span>
		<span>↓</span>
	</button>
</div>

<!--Model List Popup-->
<div class="w-64 p-4 space-y-2 rounded-lg shadow-xl card" data-popup="modelListPopup">
	{#if models !== undefined}
		{#each models as model}
			<button
				class="w-56 py-2 rounded-lg btn variant-filled-tertiary close-class"
				{disabled}
				on:click={() => {
					selected_model = model;
					$selected_model_id = selected_model.id;
				}}
			>
				<span>{model.name}</span>
			</button>
		{/each}
	{/if}
</div>

<style>
	.close-class {
	}
</style>
