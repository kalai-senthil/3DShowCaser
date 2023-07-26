<script lang="ts">
	import Hero3DText from './Hero3DText.svelte';
	import HeroCaptionsSlider from './HeroCaptionsSlider.svelte';
</script>

<section class="max-w-[80%] grid grid-cols-2 mt-14 mx-auto">
	<div class="col-span-1 p-5 text-6xl font-bold">
		<h1 class="my-2">Show case your</h1>
		<div class="flex items-end gap-x-2">
			<Hero3DText />
			<HeroCaptionsSlider />
		</div>
		<button
			class="group transition-all overflow-hidden relative text-lg hover:scale-110 font-semibold bg-secondary px-5 py-2 rounded-md shadow-md"
		>
			<div
				class="group-hover:bg-gradient-to-tr group-hover:from-gradientFrom opacity-0 group-hover:opacity-100 group-hover:to-gradientTo transition-all group-hover:scale-x-100 origin-left -z-10 absolute w-full h-full left-0 top-0 scale-x-0"
			/>
			Start Showcasing</button
		>
	</div>
</section>
