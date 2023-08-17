<script lang="ts">
	export let name: string;
	export let route: string;
</script>

<a href={route}>
	<li class="bg-secondary hover:underline text-white p-2 rounded-md shadow-md font-bold">
		{name}
	</li>
</a>
