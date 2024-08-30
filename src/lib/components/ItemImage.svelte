<script lang="ts">
	let { itemIconId }: { itemIconId: number; size: number } = $props();
</script>

<img
	class="h-12 w-12"
	src={`https://api.dofusdb.fr/img/items/250/${itemIconId}.png`}
	alt="Item image"
/>
