<script lang="ts">
	import ItemImage from './ItemImage.svelte';

	let { item }: { item: number } = $props();
</script>

<div class="flex flex-row bg-gray-100 p-2">
	<ItemImage itemId={item} size={12} />

	<p class="my-auto pl-2">
		Item {item}
	</p>
</div>
