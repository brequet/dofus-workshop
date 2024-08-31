<script lang="ts">
	import itemsStore from '$lib/stores/itemStore.svelte';
	import type { DofusItem } from '$lib/types/dofus-items';
	import IngredientItem from './IngredientItem.svelte';
	import ScrollArea from './ui/scroll-area/scroll-area.svelte';
	import Separator from './ui/separator/separator.svelte';

	const mergedRecipes = $derived(mergeRecipes(itemsStore.getSelectedItems()));

	function mergeRecipes(items: DofusItem[]): Map<number, number> {
		const recipeMap = new Map<number, number>();

		for (const item of items) {
			for (const recipeItem of item.recipe) {
				const currentQuantity = recipeMap.get(recipeItem.itemId) || 0;
				recipeMap.set(recipeItem.itemId, currentQuantity + recipeItem.quantity);
			}
		}

		return recipeMap;
	}
</script>

<h1 class="pb-4 text-2xl">Ingredient List</h1>

<ScrollArea class="flex-1">
	{#each mergedRecipes.entries() as [itemId, quantity]}
		<div class="flex flex-row">
			<IngredientItem {itemId} />

			<div class="my-auto ml-auto mr-2">
				<p>x{quantity}</p>
			</div>
		</div>
		<Separator />
	{/each}
</ScrollArea>
