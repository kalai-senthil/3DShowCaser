<script lang="ts">
	import HomeHero from '$lib/components/HomeHero.svelte';

	export let data: import('./$types').PageServerData;
</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="3D Show Caser" />
</svelte:head>

<section>
	{#if data.profile == null}
		<HomeHero />
	{:else}
		<h1>LOggedf IN</h1>
	{/if}
</section>
