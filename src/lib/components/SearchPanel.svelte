<script lang="ts">
	import SearchResultItem from './SearchResultItem.svelte';
	import Input from './ui/input/input.svelte';
	import ScrollArea from './ui/scroll-area/scroll-area.svelte';
	import Separator from './ui/separator/separator.svelte';

	let inputValue = $state('');

	let itemsMock = [...Array(5).keys()].map((i) => i + 1001);
</script>

<div class="p-2">
	<Input type="text" placeholder="Search items" bind:value={inputValue} />
</div>

<ScrollArea class="flex-1">
	{#each itemsMock as item}
		<SearchResultItem {item} />
		<Separator />
	{/each}
</ScrollArea>
