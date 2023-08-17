<script lang="ts">
	import { createEventDispatcher, onMount } from 'svelte';
	import { fly } from 'svelte/transition';
	export let message: string;
	const dispatch = createEventDispatcher();
	onMount(() => {
		setTimeout(() => {
			dispatch('hideToast');
		}, 2000);
	});
</script>

<div
	in:fly={{ y: 10 }}
	out:fly={{ y: -10 }}
	class="fixed bottom-10 text-sm font-bold bg-secondary p-2 rounded-md shadow-md px-4"
>
	<p>{message}</p>
</div>
