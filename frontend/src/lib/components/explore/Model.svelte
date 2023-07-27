<script lang="ts">
	import ViewAll from '$lib/images/view_all.svelte';
	export let name: string;
	export let route: string;
</script>

<div class="px-10 py-4">
	<div class="flex bg-secondary rounded-lg p-2 gap-2">
		<div class="flex justify-center items-center flex-grow-[1]">{name}</div>
		<div class="flex justify-center items-center flex-grow-[2] h-80 bg-primary">Image 1</div>
		<div class="flex justify-center items-center flex-grow-[2] h-80 bg-primary">Image 2</div>
		<div class="flex justify-center items-center flex-grow-[2] h-80 bg-primary">Image 3</div>
		<ViewAll {route} />
	</div>
</div>
