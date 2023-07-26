<script lang="ts">
	import { slide, fade, fly } from 'svelte/transition';
	import { onMount } from 'svelte';
	const featuring = ['arts', 'creations', 'models'];
	let index = 0;
	onMount(() => {
		setInterval(() => {
			index = (index + 1) % (featuring.length - 1);
		}, 5000);
	});
</script>

{#if index % 2}
	<h1
		in:fly={{ y: 30, duration: 1000, delay: 800 }}
		out:fly={{ y: -10, duration: 500, delay: 0 }}
		class={`inline underline mb-5`}
	>
		{featuring[index]}
	</h1>
{:else}
	<h1
		in:fly={{ y: 30, duration: 1000, delay: 800 }}
		out:fly={{ y: -10, duration: 500, delay: 0 }}
		class={`inline underline mb-5`}
	>
		{featuring[index]}
	</h1>
{/if}
