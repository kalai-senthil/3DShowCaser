<script lang="ts">
	import { fade, slide } from 'svelte/transition';

	export let title: string;
</script>

<h3 in:fade class="font-bold text-3xl tracking-widest capitalize mb-10">{title}</h3>
