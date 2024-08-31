<script lang="ts">
	import type { DofusItem } from '$lib/types/dofus-items';
	import { onMount } from 'svelte';
	import ItemImage from './ItemImage.svelte';

	let { itemId }: { itemId: number } = $props();

	let dofusItem: DofusItem | undefined = $state(undefined);

	onMount(async () => {
		dofusItem = await getItemFromId(itemId.toString());
	});

	async function getItemFromId(input: string): Promise<DofusItem | undefined> {
		let res: DofusItem | undefined = undefined;

		if (input.length === 0) return undefined;

		try {
			const response = await fetch(`/api/item/${input}`);
			if (response.ok) {
				res = (await response.json()) as DofusItem;
			}
		} catch (e) {
			console.log(e);
		}

		return res;
	}
</script>

{#if !dofusItem}
	<p>Loading...</p>
{:else}
	<ItemImage itemIconId={dofusItem.iconId} size={12} />

	<h2 class="my-auto px-2">{dofusItem.name}</h2>
{/if}
