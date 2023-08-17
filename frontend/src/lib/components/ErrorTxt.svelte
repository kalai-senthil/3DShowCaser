<script lang="ts">
	import { createEventDispatcher } from 'svelte';

	const dispatch = createEventDispatcher();
	export let txt: string;
</script>

<span class="font-normal text-red-600">{txt}</span>
