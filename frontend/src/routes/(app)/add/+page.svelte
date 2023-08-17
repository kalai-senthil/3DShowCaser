<script lang="ts">
	import { enhance } from '$app/forms';
	import PageHeader from '$lib/components/PageHeader.svelte';
	import PrimaryBtn from '$lib/components/PrimaryBtn.svelte';
	import { fly } from 'svelte/transition';
	import type { ActionData } from '../login/$types';
	import Toast from '$lib/components/Toast.svelte';
	import { goto } from '$app/navigation';
	export let form: ActionData;
	let toastMsg: undefined | string = undefined;
	$: if (form) {
		if (form.success) {
			toastMsg = 'Successfully Uploaded';
			form = null;
			setTimeout(() => {
				goto('/', { state: { isLoginSuccess: true } });
			}, 1000);
		} else {
			toastMsg = form?.message;
		}
	}
</script>

<section class="flex flex-col page">
	<PageHeader title="Add Work" />
	{#if toastMsg}
		<Toast
			on:hideToast={() => {
				toastMsg = undefined;
			}}
			message={toastMsg}
		/>
	{/if}
	<form
		in:fly={{ y: 10, opacity: 0 }}
		method="post"
		class="grid grid-cols-2 gap-5 mx-10 place-content-end"
		use:enhance
		action="?/addArtWork"
	>
		<section class="flex flex-col">
			<input name="name" type="text" placeholder="Art Work Name" />
			<textarea name="description" placeholder="Description" />
			<input name="tags" type="text" placeholder="Tags,separated by commas" />
		</section>
		<section class="flex flex-col">
			<label for="images">Images</label>
			<input name="images" type="file" accept="image/jpg" />
			<label for="art">Art</label>
			<input name="art" type="file" />
		</section>
		<div class="flex justify-end">
			<PrimaryBtn type="submit">Submit</PrimaryBtn>
		</div>
	</form>
</section>
