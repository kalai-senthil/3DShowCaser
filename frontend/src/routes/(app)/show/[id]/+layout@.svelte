<script lang="ts">
	import type { PageData } from './$types';

	export let data: PageData;
</script>

{#if data.success}
	<nav class="absolute w-full top-0 bg-secondary p-2 shadow-md px-10 py-5">
		<h1>{data.artData?.name}</h1>
	</nav>
{/if}
<slot />
