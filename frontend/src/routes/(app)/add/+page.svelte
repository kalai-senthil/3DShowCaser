<script lang="ts">
	import PageHeader from '$lib/components/PageHeader.svelte';
	import PrimaryBtn from '$lib/components/PrimaryBtn.svelte';
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<section class="flex flex-col page">
	<PageHeader title="Add Work" />
	<form class="grid grid-cols-2 gap-5 mx-10 place-content-end">
		<section class="flex flex-col">
			<input name="name" type="text" placeholder="Art Work Name" />
			<textarea name="description" placeholder="Description" />
			<input name="tags" type="text" placeholder="Tags,separated by commas" />
		</section>
		<section class="flex flex-col">
			<label>Images</label>
			<input type="file" />
			<label>Art</label>
			<input type="file" />
		</section>
		<div class="flex justify-end">
			<PrimaryBtn type="submit">Submit</PrimaryBtn>
		</div>
	</form>
</section>
