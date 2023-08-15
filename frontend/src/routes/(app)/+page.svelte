<script lang="ts">
	import Contributions from '$lib/components/Works.svelte';
	import HomeHero from '$lib/components/HomeHero.svelte';
	import Profile from '$lib/components/Profile.svelte';
	export let data: import('./$types').PageServerData;
</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="3D Show Caser" />
</svelte:head>
{#if data.profile == null}
	<section>
		<HomeHero />
	</section>
{:else}
	<section class="page">
		<Profile user={data.profile} />
		<Contributions works={data.profile.works} />
	</section>
{/if}
