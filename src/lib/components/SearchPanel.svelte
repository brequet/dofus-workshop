<script lang="ts">
	import type { DofusItem } from '$lib/types/dofus-items';
	import SearchResultItem from './SearchResultItem.svelte';
	import Input from './ui/input/input.svelte';
	import ScrollArea from './ui/scroll-area/scroll-area.svelte';
	import Separator from './ui/separator/separator.svelte';

	let inputValue = $state('');
	let foundItemsPromise = $state([] as DofusItem[]);

	$effect(() => {
		search(inputValue);
	});

	async function search(input: string): Promise<void> {
		let res: DofusItem[] = [];

		if (input.length === 0) return;

		try {
			const response = await fetch(`/api/item/search/${input}`);
			if (response.ok) {
				res = await response.json();
			}
		} catch (e) {
			console.log(e);
		}

		foundItemsPromise = res;
	}
</script>

<div class="p-2">
	<Input type="text" placeholder="Search items" bind:value={inputValue} />
</div>

{#await foundItemsPromise}
	Loading..
{:then foundItems}
	<ScrollArea class="flex-1">
		{#each foundItems as item}
			<SearchResultItem {item} />
			<Separator />
		{/each}
	</ScrollArea>
{:catch error}
	<p>Error: {error}</p>
{/await}
