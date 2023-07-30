<script lang="ts">
	import type { Work } from '$lib/types';
	import calendar from '$lib/assets/calendar.svg';
	export let works: Work[];
</script>

<section>
	<h3>Contributions</h3>
	<div class="grid grid-cols-12 gap-5 mt-5">
		{#each works as work (work.id)}
			<div class="col-span-12 md:col-span-6 lg:col-span-4 bg-secondary shadow-sm p-2 rounded-md">
				<a href={`/show/${work.id}`}>
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
