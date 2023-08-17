<script lang="ts">
	import Hero3DText from './Hero3DText.svelte';
	import HeroCaptionsSlider from './HeroCaptionsSlider.svelte';
	import PrimaryBtn from './PrimaryBtn.svelte';
</script>

<section class="max-w-[80%] grid grid-cols-2 mt-14 mx-auto">
	<div class="col-span-1 p-5 text-6xl font-bold">
		<h1 class="my-2">Show case your</h1>
		<div class="flex items-end gap-x-2">
			<Hero3DText />
			<HeroCaptionsSlider />
		</div>
		<a href="/login">
			<PrimaryBtn>Start Showcasing</PrimaryBtn>
		</a>
	</div>
</section>
