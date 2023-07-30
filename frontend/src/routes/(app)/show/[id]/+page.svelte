<script lang="ts">
	import Viewer from '$lib/components/Viewer.svelte';
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<section class="overflow-hidden">
	<Viewer data={data.artData} />
</section>
