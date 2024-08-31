<script lang="ts">
	import CraftList from '$lib/components/CraftList.svelte';
	import IngredientList from '$lib/components/IngredientList.svelte';
	import SearchPanel from '$lib/components/SearchPanel.svelte';
</script>

<div class="flex h-screen">
	<div class="flex w-1/4 flex-col bg-gray-100 p-2">
		<SearchPanel />
	</div>
	<div class="flex w-1/2 flex-col bg-gray-200 p-2">
		<IngredientList />
	</div>
	<div class="flex w-1/4 flex-col bg-gray-100 p-2">
		<CraftList />
	</div>
</div>
