<script lang="ts">
	import itemsStore from '$lib/stores/itemStore.svelte';
	import type { DofusItem } from '$lib/types/dofus-items';
	import ItemImage from './ItemImage.svelte';

	let { item }: { item: DofusItem } = $props();
</script>

<div
	on:click={() => itemsStore.selectItem(item)}
	class="flex cursor-pointer flex-row bg-gray-100 p-2 hover:bg-gray-200"
>
	<div class="flex flex-row">
		<ItemImage itemIconId={item.iconId} size={12} />

		<h2 class="my-auto px-2">{item.name}</h2>
	</div>
</div>
