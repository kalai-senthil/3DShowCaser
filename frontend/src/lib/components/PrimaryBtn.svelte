<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	export let type: 'button' | 'submit' | 'reset' | null | undefined = 'button';
	const dispatch = createEventDispatcher();
	function onClick() {
		dispatch('click');
	}
</script>

<button
	{type}
	on:click={onClick}
	class="group transition-all overflow-hidden relative text-lg hover:scale-110 font-semibold bg-secondary px-5 py-2 rounded-md shadow-md"
>
	<div
		class="group-hover:bg-gradient-to-tr group-hover:from-gradientFrom opacity-0 group-hover:opacity-100 group-hover:to-gradientTo transition-all group-hover:scale-x-100 origin-left -z-10 absolute w-full h-full left-0 top-0 scale-x-0"
	/>
	<slot />
</button>
