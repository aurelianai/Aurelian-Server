<script lang="ts">
	export let ordered: boolean;
	export let start: number;
</script>

<div class="pl-10 pr-4 py-4">
	{#if ordered}
		<ol {start} class="list-decimal space-y-1"><slot /></ol>
	{:else}
		<ul class="list-disc space-y-1"><slot /></ul>
	{/if}
</div>
