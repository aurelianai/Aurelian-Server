<script lang="ts">
	import { Icon, ChevronDown } from 'svelte-hero-icons';
	import ModelLogo from './ModelLogo.svelte';

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
	class="flex w-40 p-2 mx-auto rounded-lg shadow-xl bg-surface-100-800-token"
	use:popup={modelListPopup}
>
	<button class="w-full p-2 rounded-lg btn" {disabled}>
		<span><ModelLogo /></span>
		<span class="font-medium">{selected_model?.name || 'Select a Model'}</span>
		{#if models?.length > 1}
			<span><Icon src={ChevronDown} size="15" /></span>
		{/if}
	</button>
</div>
{#if models?.length > 1}
	<!--Model List Popup-->
	<div class="flex flex-col w-40 p-2 rounded-md shadow-xl card" data-popup="modelListPopup">
		{#if models !== undefined}
			{#each models as model}
				<button
					class="justify-start w-full p-2 rounded-md btn close-class hover:variant-soft-surface"
					{disabled}
					on:click={() => {
						selected_model = model;
						$selected_model_id = selected_model.id;
					}}
				>
					<span><ModelLogo /></span>
					<span>{model.name}</span>
				</button>
			{/each}
		{/if}
	</div>
{/if}

<style>
	.close-class {
		text-align: left;
	}
</style>
