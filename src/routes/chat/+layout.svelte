<script lang="ts">
	import ChatSidebar from './ChatSidebar.svelte';
</script>

<div class="flex h-full">
	<div
		class="fixed w-64 h-full overflow-y-auto border-r bg-secondary-500/5 border-surface-900-50-token"
	>
		<ChatSidebar />
	</div>

	<div class="fixed bottom-0 right-0 overflow-y-auto top-20 left-64">
		<slot />
	</div>
</div>