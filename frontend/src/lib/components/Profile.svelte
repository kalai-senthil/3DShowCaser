<script lang="ts">
	import type { User } from '$lib/types';
	export let user: User;
</script>

<div class="mb-5">
	<h1 class="font-bold text-2xl">Welcome</h1>
	<h1 class="font-bold text-2xl">{user.email}</h1>
</div>
