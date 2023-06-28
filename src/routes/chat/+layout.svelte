<script lang="ts">
	import type { LayoutData } from './$types';
	import ChatSidebar from './ChatSidebar.svelte';

	export let data: LayoutData;
</script>

<svelte:head>
	<title>Aurelian</title>
</svelte:head>

<div class="flex h-full">
	<div class="fixed w-64 h-full overflow-y-hidden shadow-md hover:overflow-y-auto">
		<ChatSidebar user={data} />
	</div>

	<div
		class="fixed top-0 bottom-0 right-0 overflow-y-auto left-64 bg-[url('/login-background.svg')] bg-tertiary-50-200-token"
	>
		<slot />
	</div>
</div>
