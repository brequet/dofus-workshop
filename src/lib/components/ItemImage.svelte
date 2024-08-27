<script lang="ts">
	let { itemId }: { itemId: number; size: number } = $props();
</script>

<img
	class="h-12 w-12"
	src={`https://api.dofusdb.fr/img/items/250/${itemId}.png`}
	alt="Item image"
/>
