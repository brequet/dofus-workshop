<script lang="ts">
    let imageId = $state(1001)
</script>

<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>

<input type="number" bind:value={imageId} />
<br/>
<img src="/api/image/{imageId}" alt="Resource image" />
