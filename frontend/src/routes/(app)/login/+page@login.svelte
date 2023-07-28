<script lang="ts">
	import { enhance } from '$app/forms';
	import PageHeader from '$lib/components/PageHeader.svelte';
	import type { ActionData } from './$types';
	import PrimaryBtn from '$lib/components/PrimaryBtn.svelte';
	import { fly } from 'svelte/transition';
	import Loading from '$lib/components/Loading.svelte';
	import Toast from '$lib/components/Toast.svelte';
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';

	export let form: ActionData;
	let toastMsg = undefined as string | undefined;
	$: isRegister = $page.url.searchParams.get('isRegister') === 'true' ?? false;
	$: mode = isRegister ? 'register' : 'login';
	$: if (form) {
		console.log(form);
		if (form.success) {
			toastMsg = isRegister ? 'Successfully Registered' : 'Logging you in';
			form = null;
			setTimeout(() => {
				if (isRegister) {
					goto('/login', { state: { isRegisterSuccess: true } });
				} else {
					goto('/', { state: { isLoginSuccess: true } });
				}
			}, 1000);
		} else {
			toastMsg = form?.message;
		}
	}
</script>

<section class="max-w-xl h-screen mx-auto grid place-items-center">
	{#if toastMsg}
		<Toast
			on:hideToast={() => {
				toastMsg = undefined;
			}}
			message={toastMsg}
		/>
	{/if}
	<div class="w-full">
		<PageHeader title={mode} />
		<form
			use:enhance
			method="post"
			action={`${isRegister ? '?/register' : '?/login'}`}
			in:fly={{ y: 10 }}
			class="shadow bg-secondary px-10 pt-10 pb-5 flex flex-col rounded-lg"
		>
			<input disabled={form?.success} name="email" placeholder="Email" type="email" />
			<input disabled={form?.success} name="password" type="password" placeholder="Password" />
			<div class="self-end my-2">
				{#if form?.success}
					<Loading />
				{:else}
					<PrimaryBtn on:click={() => {}} type="submit">
						<span class="underline capitalize">{mode}</span>
					</PrimaryBtn>
				{/if}
			</div>
			{#if !isRegister}
				<p class="text-end">
					Not Registered. <a class="text-orange-500 underline" href="/login?isRegister=true"
						>Register</a
					>
				</p>
			{:else}
				<p class="text-end">
					Already Registered.<a class="text-orange-500 underline" href="/login">Login</a>
				</p>
			{/if}
		</form>
	</div>
</section>
