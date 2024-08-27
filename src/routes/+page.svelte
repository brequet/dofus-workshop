<script lang="ts">
	import MainContent from '$lib/components/MainContent.svelte';
	import LeftPanel from '$lib/components/SearchPanel.svelte';
</script>

<div class="flex h-screen">
	<div class="flex w-1/4 flex-col bg-gray-100">
		<LeftPanel />
	</div>
	<div class="flex-1 bg-white p-4">
		<MainContent />
	</div>
</div>
