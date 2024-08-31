<script lang="ts">
	import itemsStore from '$lib/stores/itemStore.svelte';
	import ItemImage from './ItemImage.svelte';
	import ScrollArea from './ui/scroll-area/scroll-area.svelte';
</script>

<h1 class="pb-4 text-2xl">Craft List</h1>

<ScrollArea class="flex-1">
	{#each itemsStore.getSelectedItems() as item}
		<div class="flex flex-row">
			<ItemImage itemIconId={item.iconId} size={12} />

			<h2 class="my-auto px-2">{item.name}</h2>
			<div class="my-auto ml-auto">
				<p>x{item.selectedCount}</p>
			</div>
		</div>
	{/each}
</ScrollArea>
