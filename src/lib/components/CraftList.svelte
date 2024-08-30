<script lang="ts">
	import itemsStore from '$lib/stores/itemStore.svelte.ts';
</script>

<h1>Craft list</h1>

<div class="    ">
	{#each itemsStore.getSelectedItems() as item}
		<div>
			Item {item.name}
		</div>
	{/each}
</div>
