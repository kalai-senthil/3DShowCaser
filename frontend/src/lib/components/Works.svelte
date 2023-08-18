<script lang="ts">
	import type { Work } from '$lib/types';
	import calendar from '$lib/assets/calendar.svg';
	import AddWork from './AddWork.svelte';
	import { PUBLIC_BACKEND_URL } from '$env/static/public';

	export let works: Work[];
</script>

<section>
	<h3>Works</h3>
	<div class="grid grid-cols-12 place-content-center gap-5 mt-5">
		<AddWork />
		{#each works as work (work.id)}
			<div class="col-span-12 md:col-span-3 bg-secondary shadow-sm p-2 rounded-md">
				<a class="flex flex-col rounded-md overflow-hidden" href={`/show/${work.id}`}>
					<img
						class="w-full rounded-md mb-3 aspect-video object-cover hover:scale-110 overflow-hidden transition-transform"
						src={`${PUBLIC_BACKEND_URL}/${work.image}`}
						alt="artwork-image"
					/>
					<h4 class="font-semibold">
						{work.name}
					</h4>
					<p class="flex items-center gap-x-1">
						<img src={calendar} class="svg-img" alt="calendar" />
						<span class="text-xs">
							{new Date(work.uploadedAt).toDateString()}
						</span>
					</p>
				</a>
			</div>
		{/each}
	</div>
</section>
