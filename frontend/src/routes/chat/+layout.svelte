<script lang="ts">
	import ChatSidebar from '$lib/components/chat/ChatSidebar.svelte';
</script>

<svelte:head>
	<title>Aurelian — Home</title>
</svelte:head>

<div class="flex h-full">
	<div class="fixed w-64 h-full overflow-y-hidden shadow-md hover:overflow-y-auto">
		<ChatSidebar />
	</div>

	<div
		class="fixed top-0 bottom-0 right-0 overflow-y-auto left-64 bg-[url('/login-background.svg')] bg-tertiary-50-200-token"
	>
		<slot />
	</div>
</div>
